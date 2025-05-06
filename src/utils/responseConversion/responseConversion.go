package responseConversion

import (
	"context"
	"fmt"
	"omnenest-backend/src/constants"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/tracer"
	"reflect"
	"strings"
	"time"
)

// ConvertFieldValues converts the field values of the input struct or slice of structs based on the specified conversion rules.
func ConvertFieldValues(ctx context.Context, input interface{}, isRequest bool) {
	ctx, span := tracer.AddToSpan(ctx, "ConvertFieldValues")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	inputValue := reflect.ValueOf(input).Elem()

	// Check if input is a slice
	if inputValue.Kind() == reflect.Slice {
		for i := 0; i < inputValue.Len(); i++ {
			itemValue := inputValue.Index(i)
			convertFields(ctx, itemValue, isRequest)
		}
	} else if inputValue.Kind() == reflect.Struct {
		convertFields(ctx, inputValue, isRequest)
	}
}

// convertFields is a function that performs field conversion on the given input value based on certain conditions.
func convertFields(ctx context.Context, inputValue reflect.Value, isRequest bool) {
	ctx, span := tracer.AddToSpan(ctx, "convertFields")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	// Converting all date formats to a common format.
	// Iterate over each field and perform the date format conversion if the field name contains 'date'.
	if !isRequest {
		dateFormatCheckRegex := configs.GetRegexPattern(constants.DateFormatMatchKey)
		for i := 0; i < inputValue.NumField(); i++ {
			field := inputValue.Field(i)
			fieldName := inputValue.Type().Field(i).Name

			if strings.Contains(strings.ToLower(fieldName), constants.DateInField) && field.Kind() == reflect.String {
				dateValue := field.String()

				if dateFormatCheckRegex.MatchString(dateValue) {
					continue
				}
				inputLayouts := []string{constants.DateOldLayout1, constants.DateOldLayout2, constants.DateOldLayout3, constants.DateOldLayout4, constants.DateOldLayout5, constants.DateOldLayout6, constants.DateOldLayout7, constants.DateOldLayout8, constants.DateOldLayout9, constants.DateOldLayout10, constants.DateOldLayout11, constants.DateOldLayout12}
				formattedDate, _ := ConvertDateFormat(dateValue, inputLayouts, constants.DateNewLayout)
				field.SetString(formattedDate)
			}
		}
	}

	// List of fields to be converted.
	// Iterate over each field and perform the conversion if the value is of type string.
	fieldsToConvert := []string{constants.TransactionType, constants.Leg1TransactionType, constants.Leg2TransactionType, constants.Leg3TransactionType, constants.Leg4TransactionType, constants.PriceType, constants.PositionType, constants.ProductCode, constants.RetentionType, constants.ExchangeName, constants.OrderType, constants.OriginalPriceType, constants.SendAlertsOn, constants.InputOrderStatus, constants.BFFOrderStatus, constants.BFFBidStatus, constants.BidHistory, constants.InstrumentName, constants.Option, constants.SourceField, constants.AlertType, constants.ScannerType, constants.ScannerTypeValue, constants.CurrentProductCode, constants.TargetProductCode, constants.AfterMarketOrderFlag, constants.OFSStatus, constants.IPOStatus, constants.ClientSubCategoryCode, constants.BFFTriggerStatus, constants.PaymentMode, constants.SubCategory, constants.BiddingType}
	var normalizedValue string

	for _, fieldName := range fieldsToConvert {
		if fieldValue := inputValue.FieldByName(fieldName); fieldValue.Kind() == reflect.String {
			if isRequest {
				value := strings.ToUpper(fieldValue.String())
				normalizedValue = NormalizeValue(value, constants.BFFToNestRequestMapping)
			} else {
				value := fieldValue.String()
				normalizedValue = NormalizeValue(value, constants.NestToBFFResponseMapping)
			}
			fieldValue.SetString(normalizedValue)
		} else if fieldValue := inputValue.FieldByName(fieldName); fieldValue.Kind() == reflect.Slice {
			if isRequest {
				if stringValue, ok := fieldValue.Interface().([]string); ok {
					for i := range stringValue {
						stringValue[i] = strings.ToUpper(stringValue[i])
					}
					NormalizeValueArray(&stringValue, constants.BFFToNestRequestMapping)
					fieldValue.Set(reflect.ValueOf(stringValue))
				} else {
					for i := 0; i < fieldValue.Len(); i++ {
						itemValue := fieldValue.Index(i)
						convertFields(ctx, itemValue, isRequest)
					}
				}
			} else {
				if stringValue, ok := fieldValue.Interface().([]string); ok {
					NormalizeValueArray(&stringValue, constants.NestToBFFResponseMapping)
					fieldValue.Set(reflect.ValueOf(stringValue))
				} else {
					for i := 0; i < fieldValue.Len(); i++ {
						itemValue := fieldValue.Index(i)
						convertFields(ctx, itemValue, isRequest)
					}
				}
			}
		}
	}
}

// ConvertDateFormat takes an input date string, a list of input layouts, and an output layout and converts the input date string to the specified output layout.
func ConvertDateFormat(inputDate string, inputLayouts []string, outputLayout string) (string, error) {
	for _, inputLayout := range inputLayouts {
		parsedDate, err := time.Parse(inputLayout, inputDate)
		if err == nil {
			formattedDate := parsedDate.Format(outputLayout)
			return formattedDate, nil
		}
	}
	return "", fmt.Errorf(constants.DateParseError)
}

func NormalizeValue(value string, mapping map[string]string) string {
	normalizedValue, ok := mapping[value]
	if !ok {
		return value
	}
	return normalizedValue
}

func NormalizeValueArray(values *[]string, mapping map[string]string) {
	for i, value := range *values {
		normalizedValue, ok := mapping[value]
		if ok {
			(*values)[i] = normalizedValue
		}
	}
}

// SingleFieldConvert converts the value of the specified field in the input struct to uppercase, normalizes it using the NormalizeValue function, and sets the normalized value back to the field.
func SingleFieldConvert(ctx context.Context, input interface{}, fieldName string, mapStruct map[string]string) {
	_, span := tracer.AddToSpan(ctx, "SingleFieldConvert")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	fieldValue := reflect.ValueOf(input).Elem().FieldByName(fieldName)

	if fieldValue.Kind() == reflect.String {
		value := strings.ToUpper(fieldValue.String())
		normalizedValue := NormalizeValue(value, mapStruct)
		fieldValue.SetString(normalizedValue)
	}
}
