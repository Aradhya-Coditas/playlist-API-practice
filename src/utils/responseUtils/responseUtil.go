package responseUtils

import (
	"errors"
	"net/http"
	"omnenest-backend/src/constants"
	"omnenest-backend/src/models"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func SendBadRequest(ctx *gin.Context, errorMsgs []models.ErrorMessage) {
	if errorMsgs[0].Key == constants.GenericErrorKey {
		errorStr := errorMsgs[0].ErrorMessage //err.Error()
		startIndex := strings.Index(errorStr, constants.NestInternalServerErrorStartTag)
		if startIndex != -1 {
			endIndex := strings.Index(errorStr, constants.NestInternalServerErrorEndTag)
			errorStr = errorStr[startIndex+len(constants.NestInternalServerErrorStartTag) : endIndex]
		}
		ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(errorStr), Type: gin.ErrorTypePrivate})
		ctx.JSON(http.StatusBadRequest, genericModel.ErrorAPIResponse{
			Message: []models.ErrorMessage{{
				Key:          constants.GenericErrorKey,
				ErrorMessage: errorStr,
			}},
			Error: http.StatusText(http.StatusBadRequest),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, genericModel.ErrorAPIResponse{
			Message: errorMsgs,
			Error:   http.StatusText(http.StatusBadRequest),
		})
	}
}

func SendUnauthorized(ctx *gin.Context, err error) {
	errorStr := err.Error()
	startIndex := strings.Index(errorStr, constants.NestInternalServerErrorStartTag)
	if startIndex != -1 {
		endIndex := strings.Index(errorStr, constants.NestInternalServerErrorEndTag)
		errorStr = errorStr[startIndex+len(constants.NestInternalServerErrorStartTag) : endIndex]
	}
	ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(errorStr), Type: gin.ErrorTypePrivate})
	ctx.JSON(http.StatusUnauthorized, genericModel.ErrorAPIResponse{
		Message: []models.ErrorMessage{{
			Key:          constants.GenericErrorKey,
			ErrorMessage: errorStr,
		}},
		Error: http.StatusText(http.StatusUnauthorized),
	})
}

func SendInternalServerError(ctx *gin.Context, err error) {
	errorStr := err.Error()
	startIndex := strings.Index(errorStr, constants.NestInternalServerErrorStartTag)
	if startIndex != -1 {
		endIndex := strings.Index(errorStr, constants.NestInternalServerErrorEndTag)
		errorStr = errorStr[startIndex+len(constants.NestInternalServerErrorStartTag) : endIndex]
	} else {
		errorStr = constants.InternalServerError
	}
	ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(errorStr), Type: gin.ErrorTypePrivate})
	ctx.JSON(http.StatusInternalServerError, genericModel.ErrorAPIResponse{
		Message: []models.ErrorMessage{{
			Key:          constants.GenericErrorKey,
			ErrorMessage: errorStr,
		}},
		Error: http.StatusText(http.StatusInternalServerError),
	})
}

func SendNoContentFoundError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNoContent, genericModel.ErrorAPIResponse{})
}

func SendConflict(ctx *gin.Context, errorMsg string) {
	ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(errorMsg), Type: gin.ErrorTypePrivate})
	ctx.JSON(http.StatusConflict, genericModel.ErrorAPIResponse{
		Message: []models.ErrorMessage{{
			Key:          constants.GenericErrorKey,
			ErrorMessage: errorMsg,
		}},
		Error: http.StatusText(http.StatusConflict),
	})
}

func SendStatusOK(ctx *gin.Context, msg string, data interface{}) {
	ctx.Set(constants.ResponseBody, data)
	ctx.JSON(http.StatusOK, data)
}

func SendCreated(ctx *gin.Context, msg string, data interface{}) {
	ctx.Set(constants.ResponseBody, data)
	ctx.JSON(http.StatusCreated, data)
}

func SendStatusUnprocessableEntity(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusUnprocessableEntity, data)
}

// This needs to be used only in middleware
func SendAbortWithStatusJSON(ctx *gin.Context, status int, errorMsg error) {
	ctx.Errors = append(ctx.Errors, &gin.Error{Err: errorMsg, Type: gin.ErrorTypePrivate})
	ctx.AbortWithStatusJSON(status, genericModel.ErrorAPIResponse{
		Message: []models.ErrorMessage{{
			Key:          constants.GenericErrorKey,
			ErrorMessage: errorMsg.Error(),
		}},
		Error: http.StatusText(status),
	})
}

func SendNotFoundJSON(ctx *gin.Context, errorMsg error) {
	// Append the error to the context's error slice
	ctx.Errors = append(ctx.Errors, &gin.Error{Err: errorMsg, Type: gin.ErrorTypePrivate})

	// Create the JSON response structure
	ctx.AbortWithStatusJSON(http.StatusNotFound, genericModel.ErrorAPIResponse{
		Message: []models.ErrorMessage{{
			Key:          constants.GenericErrorKey,
			ErrorMessage: errorMsg.Error(),
		}},
		Error: http.StatusText(http.StatusNotFound),
	})
}

func SendForbidden(ctx *gin.Context, err error) {
	errorStr := err.Error()
	startIndex := strings.Index(errorStr, constants.NestInternalServerErrorStartTag)
	if startIndex != -1 {
		endIndex := strings.Index(errorStr, constants.NestInternalServerErrorEndTag)
		errorStr = errorStr[startIndex+len(constants.NestInternalServerErrorStartTag) : endIndex]
	}
	ctx.Errors = append(ctx.Errors, &gin.Error{Err: errors.New(errorStr), Type: gin.ErrorTypePrivate})
	ctx.JSON(http.StatusForbidden, genericModel.ErrorAPIResponse{
		Message: []models.ErrorMessage{{
			Key:          constants.GenericErrorKey,
			ErrorMessage: errorStr,
		}},
		Error: http.StatusText(http.StatusForbidden),
	})
}

// StructHide recursively hides fields in a struct based on the `hide` tag
func StructHide(input interface{}, source string) interface{} {
	val := reflect.ValueOf(input)

	// If input is a pointer, get the underlying value
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Check the kind of value and apply hiding rules accordingly
	switch val.Kind() {
	case reflect.Slice:
		// If value is a slice, iterate over elements and hide recursively
		resultSlice := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i).Interface()
			resultSlice[i] = StructHide(elem, source)
		}
		return resultSlice
	case reflect.Map:
		// If value is a map, iterate over keys and hide values recursively
		resultMap := make(map[string]interface{})
		for _, key := range val.MapKeys() {
			resultMap[key.String()] = StructHide(val.MapIndex(key).Interface(), source)
		}
		return resultMap
	case reflect.Struct:
		// If value is a struct, check for hide tag and apply hiding
		if !hasHideTag(val, source) {
			return input
		}
		result := make(map[string]interface{})
		handleFields(val, result, source)
		return result
	default:
		return input
	}
}

func handleFields(val reflect.Value, result map[string]interface{}, source string) {
	typ := val.Type()

	// Iterate over each field in the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Check if the field should be hidden for the specified source
		hideTag := fieldType.Tag.Get(constants.HideFieldTag)
		hideTags := strings.Split(hideTag, ",")
		if utils.Contains(hideTags, source) {
			continue
		}

		jsonTag := fieldType.Tag.Get("json")
		jsonTags := strings.Split(jsonTag, ",")
		if len(jsonTags) > 0 && jsonTags[0] == "-" {
			continue
		}

		jsonFieldName := fieldType.Name
		if len(jsonTags) > 0 && jsonTags[0] != "" {
			jsonFieldName = jsonTags[0]
		}

		// Handle different kinds of fields
		switch field.Kind() {
		case reflect.Struct:
			// Recursively handle nested struct fields
			nestedResult := make(map[string]interface{})
			handleFields(field, nestedResult, source)
			if len(nestedResult) > 0 || !utils.Contains(jsonTags, "omitempty") {
				result[jsonFieldName] = nestedResult
			}
		case reflect.Slice:
			if field.Len() > 0 && field.Index(0).Kind() == reflect.Struct {
				// Handle slice of structs
				sliceResult := make([]map[string]interface{}, 0)
				for j := 0; j < field.Len(); j++ {
					elemResult := make(map[string]interface{})
					handleFields(field.Index(j), elemResult, source)
					if len(elemResult) > 0 || !utils.Contains(jsonTags, "omitempty") {
						sliceResult = append(sliceResult, elemResult)
					}
				}
				result[jsonFieldName] = sliceResult
			} else {
				// Handle other types of slices
				if field.Len() > 0 || !utils.Contains(jsonTags, "omitempty") {
					result[jsonFieldName] = field.Interface()
				}
			}
		default:
			if !field.IsZero() || !utils.Contains(jsonTags, "omitempty") {
				result[jsonFieldName] = field.Interface()
			}
		}
	}
}

func hasHideTag(val reflect.Value, source string) bool {
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		hideTag := fieldType.Tag.Get(constants.HideFieldTag)
		hideTags := strings.Split(hideTag, ",")

		// Check if the field should be hidden for the specified source
		if utils.Contains(hideTags, source) {
			return true
		}

		// Check if the field is a struct and recursively check for hide tag
		if field.Kind() == reflect.Struct {
			if hasHideTag(field, source) {
				return true
			}
		}
	}
	return false
}
