package logger

import (
	"omnenest-backend/src/constants"
	"reflect"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type SensitiveFieldEncoder struct {
	zapcore.Encoder
	cfg zapcore.EncoderConfig
}

// SensitiveFieldEncoder wraps zapcore.Encoder to provide additional functionality
// for sanitizing sensitive fields.
func NewSensitiveFieldsEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	encoder := zapcore.NewJSONEncoder(config)
	return &SensitiveFieldEncoder{encoder, config}
}

func (encode *SensitiveFieldEncoder) Clone() zapcore.Encoder {
	return &SensitiveFieldEncoder{
		Encoder: encode.Encoder.Clone(),
	}
}

func (encode *SensitiveFieldEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	filtered := make([]zapcore.Field, 0, len(fields))
	for _, field := range fields {
		if field.Key == constants.RequestBody || field.Key == constants.ResponseBody || field.Key == constants.ResponseConfig || field.Key == constants.RequestPayloadConfig || field.Key == constants.DBResponseConfig || field.Key == constants.NestCPPCallBackFunctionParametersConfig {
			processedField := encode.processField(field)
			filtered = append(filtered, processedField)
		} else {
			filtered = append(filtered, field)
		}
	}
	return encode.Encoder.EncodeEntry(entry, filtered)
}

// processField processes a zapcore.Field by sanitizing sensitive fields.
//
// It takes a zapcore.Field as input and returns a modified zapcore.Field.
// The function first checks if the field's value is a pointer and not nil.
// If it is, it dereferences the pointer. Then, it checks if the value is a struct.
// If it is, the function creates a copy of the original value and iterates over its fields.
// For each field, it checks if the field has a logger configuration tag with the value "sensitive".
// If it does, it calls the sanitizeField method to sanitize the field.
// Finally, it updates the field's interface with the modified value and returns the modified field.
func (encode *SensitiveFieldEncoder) processField(field zapcore.Field) zapcore.Field {
	originalValue := reflect.ValueOf(field.Interface)
	if originalValue.Kind() == reflect.Ptr && !originalValue.IsNil() {
		originalValue = originalValue.Elem()
	}
	switch originalValue.Kind() {
	case reflect.Struct:
		field.Interface = encode.sanitizeStruct(originalValue).Interface()
	case reflect.Slice:
		if originalValue.Type().Elem().Kind() == reflect.Struct {
			field.Interface = encode.sanitizeSliceOfStructs(originalValue).Interface()
		} else if originalValue.Type().Elem().Kind() == reflect.String {
			field.Interface = encode.sanitizeSliceOfStrings(originalValue).Interface()
		}
	}
	return field
}

func (encode *SensitiveFieldEncoder) sanitizeSliceOfStructs(slice reflect.Value) reflect.Value {
	newSlice := reflect.MakeSlice(slice.Type(), slice.Len(), slice.Cap())
	for fieldIndex := 0; fieldIndex < slice.Len(); fieldIndex++ {
		elem := slice.Index(fieldIndex)
		copiedElem := reflect.New(elem.Type()).Elem()
		copiedElem.Set(elem)
		newSlice.Index(fieldIndex).Set(encode.sanitizeStruct(copiedElem))
	}
	return newSlice
}

func (encode *SensitiveFieldEncoder) sanitizeStruct(originalValue reflect.Value) reflect.Value {
	copiedValue := reflect.New(originalValue.Type()).Elem()
	copiedValue.Set(originalValue)
	structType := copiedValue.Type()
	for fieldIndex := 0; fieldIndex < copiedValue.NumField(); fieldIndex++ {
		fieldValue := copiedValue.Field(fieldIndex)
		fieldInfo := structType.Field(fieldIndex)
		if fieldInfo.Tag.Get(constants.LoggerConfig) == constants.LoggerSensitiveTag {
			encode.sanitizeField(fieldValue)
		} else if fieldValue.Kind() == reflect.Struct {
			fieldValue.Set(encode.sanitizeStruct(fieldValue))
		} else if fieldValue.Kind() == reflect.Slice && fieldValue.Type().Elem().Kind() == reflect.Struct {
			if fieldValue.Type().Elem().Kind() == reflect.Struct {
				fieldValue.Set(encode.sanitizeSliceOfStructs(fieldValue))
			} else if fieldValue.Type().Elem().Kind() == reflect.String {
				fieldValue.Set(encode.sanitizeSliceOfStrings(fieldValue))
			}
		}
	}
	return copiedValue
}

func (encode *SensitiveFieldEncoder) sanitizeSliceOfStrings(slice reflect.Value) reflect.Value {
	newSlice := reflect.MakeSlice(slice.Type(), slice.Len(), slice.Cap())
	for fieldIndex := 0; fieldIndex < slice.Len(); fieldIndex++ {
		newSlice.Index(fieldIndex).SetString("********")
	}
	return newSlice
}

func (encode *SensitiveFieldEncoder) sanitizeField(fieldValue reflect.Value) {
	switch fieldValue.Kind() {
	case reflect.String:
		fieldValue.SetString("********")
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fieldValue.SetUint(0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldValue.SetInt(0)
	case reflect.Float32, reflect.Float64:
		fieldValue.SetFloat(0)
	case reflect.Struct:
		fieldValue.Set(encode.sanitizeStruct(fieldValue))
	case reflect.Slice:
		switch fieldValue.Type().Elem().Kind() {
		case reflect.String:
			fieldValue.Set(encode.sanitizeSliceOfStrings(fieldValue))
		case reflect.Struct:
			fieldValue.Set(encode.sanitizeSliceOfStructs(fieldValue))
		}
	case reflect.Ptr:
		encode.sanitizePointer(fieldValue)
	}
}

func (encode *SensitiveFieldEncoder) sanitizePointer(fieldValue reflect.Value) {
	if !fieldValue.IsNil() {
		switch fieldValue.Type().Elem().Kind() {
		case reflect.Uint:
			newValue := uint(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Uint16:
			newValue := uint16(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Uint32:
			newValue := uint32(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Uint64:
			newValue := uint64(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Float32:
			newValue := float32(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Float64:
			newValue := float64(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.String:
			fieldValue.Set(reflect.ValueOf("********"))
		case reflect.Int:
			newValue := int(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Int8:
			newValue := int8(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Int16:
			newValue := int16(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Int32:
			newValue := int32(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Int64:
			newValue := int64(0)
			fieldValue.Set(reflect.ValueOf(&newValue))
		case reflect.Struct:
			fieldValue.Set(encode.sanitizeStruct(fieldValue))
		}
	}
}
