package validations

import (
	"context"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/tracer"
	"reflect"
	"strings"

	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var bffValidator *validator.Validate
var customErrorMap = map[string]string{
	"RetentionDateValidation":    genericConstants.RetentionDateFormatValidationError,
	"ScannerTypeValueValidation": genericConstants.ScannerTypeValueValidationError,
	"DateOfBirthValidation":      genericConstants.DateFormatValidationError,
	"PANValidation":              genericConstants.PANFormatValidationError,
	"ValidateEnum":               genericConstants.EnumValidationError,
	"lt":                         genericConstants.LtValidationError,
	"lte":                        genericConstants.LteValidationError,
	"gt":                         genericConstants.GtValidationError,
	"gte":                        genericConstants.GteValidationError,
	"min":                        genericConstants.MinValidationError,
	"ltefield":                   genericConstants.LteFieldValidationError,
	"required_if":                genericConstants.RequiredValidationError,
	"required_with":              genericConstants.RequiredValidationError,
	"required_without":           genericConstants.RequiredWithoutValidationError,
	"required_without_all":       genericConstants.RequiredWithoutAllValidationError,
	"max":                        genericConstants.MaxValidationError,
	"alphanum":                   genericConstants.AlphaNumericValidationError,
	"BidLengthValidation":        genericConstants.BidLengthValidationError,
}

// formatValidationErrors formats validation errors into a user-friendly format
func FormatValidationErrors(ctx context.Context, validationErrors validator.ValidationErrors) ([]models.ErrorMessage, string) {
	_, span := tracer.AddToSpan(ctx, "FormatValidationErrors")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	var errorMessages []models.ErrorMessage
	var errorMessagesString []string

	applicationConfig, _ := configs.Get(genericConstants.ApplicationConfig)
	useFrontendErrorFormat := applicationConfig.GetBool(genericConstants.UseFrontendErrorFormatConfig)

	for _, err := range validationErrors {
		errorMessage := err.Tag()
		errorParam := err.Param()
		if errMsg, ok := customErrorMap[errorMessage]; ok {
			errorMessage = errMsg
			if errorParam != "" {
				errorMessage = fmt.Sprintf(errMsg, errorParam)
			}
		}

		errorMessagesString = append(errorMessagesString, fmt.Sprintf("%s: %s", err.Field(), errorMessage))
		if !useFrontendErrorFormat {
			errorMessages = append(errorMessages, models.ErrorMessage{Key: err.Field(), ErrorMessage: errorMessage})
		}
	}

	errorMessagesJoined := strings.Join(errorMessagesString, ", ")
	if useFrontendErrorFormat {
		errorMessages = append(errorMessages, models.ErrorMessage{Key: genericConstants.GenericErrorKey, ErrorMessage: errorMessagesJoined})
	}

	return errorMessages, errorMessagesJoined
}

type Enum interface {
	IsValid() bool
}

// PrepareNestValidationErrors is a function that prepares validation errors in a user-friendly format.
func PrepareNestValidationErrors(ctx context.Context, key, errMsg string) ([]models.ErrorMessage, string) {
	_, span := tracer.AddToSpan(ctx, "PrepareNestValidationErrors")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	applicationConfig, _ := configs.Get(genericConstants.ApplicationConfig)
	useFrontendErrorFormat := applicationConfig.GetBool(genericConstants.UseFrontendErrorFormatConfig)
	var errorMessages []models.ErrorMessage
	if key != "" && key != genericConstants.GenericErrorKey {
		errMsg = fmt.Sprintf("%s: %s", key, errMsg)
	} else {
		key = genericConstants.GenericErrorKey
	}
	if !useFrontendErrorFormat {
		errorMessages = append(errorMessages, models.ErrorMessage{Key: key, ErrorMessage: errMsg})
	} else {
		errorMessages = append(errorMessages, models.ErrorMessage{Key: genericConstants.GenericErrorKey, ErrorMessage: errMsg})
	}
	return errorMessages, errMsg
}

// ValidateEnum is a validation function that checks if the value of a field implements the Enum interface and is valid.
func ValidateEnum[E Enum](fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(E)
	return value.IsValid()
}

// RetentionDateValidation is a validation function that checks if the value of a field is a valid retention date format.
func RetentionDateValidation(fl validator.FieldLevel) bool {
	// Define a regular expression pattern for a valid retention date format (DD/MM/YYYY OR D/M/YYYY)
	pattern := genericConstants.RetentionDateFormat
	input := fl.Field().String()

	// Check if the retention date matches the pattern
	_, err := regexp.MatchString(pattern, input)
	if err != nil {
		return false
	}

	_, err = time.Parse(genericConstants.RetentionDateFormatMatch, input)
	return err == nil
}

// ScannerTypeValueValidation is a validation function that checks if the value of a field is a valid scanner type value.
func ScannerTypeValueValidation(fl validator.FieldLevel) bool {
	scannerTypeValue := fl.Field().String()
	scannerTypeValueKey := strings.ReplaceAll(strings.ToLower(scannerTypeValue), " ", "")

	scannerType := fl.Parent().FieldByName(genericConstants.ScannerType).Interface()
	scannerTypeString := fmt.Sprintf("%s", scannerType)

	// Extract unique values from the ScannersTypeMap
	uniqueScannerTypes := make(map[string]struct{})
	for _, value := range genericConstants.ScannersTypeMap {
		uniqueScannerTypes[value] = struct{}{}
	}

	// Check if the ScannerType is valid before applying ScannerTypeValue validation
	if _, valid := uniqueScannerTypes[scannerTypeString]; valid {
		return scannerTypeString == genericConstants.ScannersTypeMap[scannerTypeValueKey]
	}

	return true // Skip ScannerTypeValue validation if ScannerType is not valid
}

func PANValidation(fl validator.FieldLevel) bool {
	// Define a regular expression pattern for a valid PAN
	input := fl.Field().String()

	// Check if the PAN matches the pattern
	configs.InitRegexPatterns()
	regex := configs.GetRegexPattern(genericConstants.PanCardValidationKey)
	return regex.MatchString(input)
}

func DateOfBirthValidation(fl validator.FieldLevel) bool {
	// Define a regular expression pattern for a valid date of birth (YYYY-MM-DD)
	pattern := genericConstants.DateOfBirthFormat
	input := fl.Field().String()

	// Check if the DOB matches the pattern
	_, err := regexp.MatchString(pattern, input)
	if err != nil {
		return false
	}

	parsedDOB, err := time.Parse(genericConstants.DateOfBirthFormatMatch, input)
	if err != nil {
		return false
	}

	// Get the current date
	currentDate := time.Now()

	// Compare the parsed DOB with the current date
	return !parsedDOB.After(currentDate)
}

// NewBFFValidator initializes a new instance of the BFFValidator and registers custom validation functions.
func NewBFFValidator(ctx context.Context) {
	_, span := tracer.AddToSpan(ctx, "NewBFFValidator")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	bffValidator = validator.New()
	bffValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(genericConstants.JsonConfig)
	})
	bffValidator.RegisterValidation(genericConstants.ValidateEnumConfig, ValidateEnum[Enum])
	bffValidator.RegisterValidation(genericConstants.RetentionDateConfig, RetentionDateValidation)
	bffValidator.RegisterValidation(genericConstants.ScannerTypeValueConfig, ScannerTypeValueValidation)
	bffValidator.RegisterValidation(genericConstants.PANConfig, PANValidation)
	bffValidator.RegisterValidation(genericConstants.DataOfBirthConfig, DateOfBirthValidation)
	bffValidator.RegisterValidation(genericConstants.BidLengthValidation, BidLengthValidation)
}

func GetBFFValidator(ctx context.Context) *validator.Validate {
	ctx, span := tracer.AddToSpan(ctx, "GetBFFValidator")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	if bffValidator == nil {
		NewBFFValidator(ctx)
	}
	return bffValidator
}

// BidLengthValidation is a validation function that checks if the value of a field is a non-empty slice.
func BidLengthValidation(fl validator.FieldLevel) bool {
	f1Value := reflect.ValueOf(fl.Field().Interface())
	if f1Value.Kind() == reflect.Slice {
		if f1Value.Len() > 0 {
			return true
		}
	}
	return false
}
