package mapStruct

import (
	"context"
	"database/sql"
	"omnenest-backend/src/constants"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/tracer"
	"reflect"
	"strconv"
	"strings"
)

/*
MapStruct maps the fields from source to destination struct.
It recursively maps the fields of nested structs and handles slice mapping as well.
*/
func MapStruct(ctx context.Context, from interface{}, to interface{}) {
	ctx, span := tracer.AddToSpan(ctx, "MapStruct")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	fromValue := reflect.ValueOf(from)
	toValue := reflect.ValueOf(to).Elem()

	for i := 0; i < fromValue.NumField(); i++ {
		fromField := fromValue.Field(i)
		toField := toValue.FieldByName(fromValue.Type().Field(i).Name)

		if toField.IsValid() && toField.CanSet() {
			toFieldType := toField.Type()

			if fromField.Kind() == reflect.Ptr && fromField.Type().Elem() == toFieldType {
				if !fromField.IsNil() {
					toField.Set(fromField.Elem())
				}
			} else if fromField.Type() == toFieldType {
				toField.Set(fromField)
			} else if fromField.Kind() == reflect.Struct && toFieldType.Kind() == reflect.Struct {
				MapStruct(ctx, fromField.Interface(), toField.Addr().Interface())
			} else if fromField.Kind() == reflect.Slice && toFieldType.Kind() == reflect.Slice {
				mapSlice(ctx, fromField, toField)
			} else {
				mapField(ctx, fromField, toField)
			}
		}
	}
}

// mapSlice is a function that maps the values from one slice to another based on their types.
func mapSlice(ctx context.Context, fromField reflect.Value, toField reflect.Value) {
	ctx, span := tracer.AddToSpan(ctx, "mapSlice")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	fromElemType := fromField.Type().Elem()
	toElemType := toField.Type().Elem()

	// Specific case for converting an array of strings to an array of uint
	if fromElemType.Kind() == reflect.String && toElemType.Kind() == reflect.Uint {
		for i := 0; i < fromField.Len(); i++ {
			strVal := fromField.Index(i).String()
			uintVal, err := strconv.ParseUint(strVal, 10, toElemType.Bits())
			if err == nil {
				toField.Set(reflect.Append(toField, reflect.ValueOf(uintVal).Convert(toElemType)))
			}
		}
		return
	}

	// Specific case for converting an array of *unit16 to an array of string
	if fromElemType.Kind() == reflect.Ptr && fromElemType.Elem().Kind() == reflect.Uint16 && toElemType.Kind() == reflect.String {
		for i := 0; i < fromField.Len(); i++ {
			uint16Val := fromField.Index(i).Elem().Uint()
			toField.Set(reflect.Append(toField, reflect.ValueOf(strconv.FormatUint(uint16Val, 10))))
		}
		return
	}

	// Specific case for converting an array of strings to an array of float32
	if fromElemType.Kind() == reflect.String && toElemType.Kind() == reflect.Float32 {
		for i := 0; i < fromField.Len(); i++ {
			strVal := fromField.Index(i).String()
			floatVal, err := strconv.ParseFloat(strVal, toElemType.Bits())
			if err == nil {
				toField.Set(reflect.Append(toField, reflect.ValueOf(float32(floatVal)).Convert(toElemType)))
			}
		}
		return
	}

	// General case for converting slices
	if fromElemType == toElemType {
		toField.Set(fromField)
		return
	}

	if fromElemType.Kind() == reflect.Struct && toElemType.Kind() == reflect.Struct {
		for i := 0; i < fromField.Len(); i++ {
			fromElem := fromField.Index(i)
			toElem := reflect.New(toElemType).Elem()
			MapStruct(ctx, fromElem.Interface(), toElem.Addr().Interface())
			toField.Set(reflect.Append(toField, toElem))
		}
		return
	}
}

// mapField maps the value of the source data to the destination data based on their types.
func mapField(ctx context.Context, fromField reflect.Value, toField reflect.Value) {
	ctx, span := tracer.AddToSpan(ctx, "mapField")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	// Dereference the pointers if they are pointers
	if fromField.Kind() == reflect.Ptr {
		fromField = fromField.Elem()
	}

	if toField.Kind() == reflect.Ptr {
		toField = toField.Elem()
	}

	if !fromField.IsValid() || !toField.IsValid() {
		return
	}

	toFieldType := toField.Type()
	fromFieldType := fromField.Type()

	switch toFieldType.Kind() {
	case reflect.String:
		if fromFieldType.Kind() == reflect.Int || fromFieldType.Kind() == reflect.Int8 || fromFieldType.Kind() == reflect.Int16 || fromFieldType.Kind() == reflect.Int32 || fromFieldType.Kind() == reflect.Int64 {
			toField.SetString(strconv.FormatInt(fromField.Int(), 10))
		} else if fromFieldType.Kind() == reflect.Uint || fromFieldType.Kind() == reflect.Uint8 || fromFieldType.Kind() == reflect.Uint16 || fromFieldType.Kind() == reflect.Uint32 || fromFieldType.Kind() == reflect.Uint64 {
			toField.SetString(strconv.FormatUint(fromField.Uint(), 10))
		} else if fromFieldType.Kind() == reflect.Float32 {
			toField.SetString(strconv.FormatFloat(fromField.Float(), 'f', -1, 32))
		} else if fromFieldType.Kind() == reflect.Float64 {
			toField.SetString(strconv.FormatFloat(fromField.Float(), 'f', -1, 64))
		} else if fromFieldType == reflect.TypeOf(sql.NullString{}) {
			nullString := fromField.Interface().(sql.NullString)
			if nullString.Valid {
				toField.SetString(nullString.String)
			} else {
				toField.SetString("")
			}
		} else if fromFieldType == reflect.TypeOf(sql.NullInt64{}) {
			nullInt64 := fromField.Interface().(sql.NullInt64)
			if nullInt64.Valid {
				toField.SetString(strconv.FormatInt(nullInt64.Int64, 10))
			} else {
				toField.SetString("")
			}
		} else if fromFieldType == reflect.TypeOf(sql.NullInt32{}) {
			nullInt32 := fromField.Interface().(sql.NullInt32)
			if nullInt32.Valid {
				toField.SetString(strconv.FormatInt(int64(nullInt32.Int32), 10))
			} else {
				toField.SetString("")
			}
		} else if fromFieldType == reflect.TypeOf(sql.NullInt16{}) {
			nullInt16 := fromField.Interface().(sql.NullInt16)
			if nullInt16.Valid {
				toField.SetString(strconv.FormatInt(int64(nullInt16.Int16), 10))
			} else {
				toField.SetString("")
			}
		} else if fromFieldType == reflect.TypeOf(sql.NullFloat64{}) {
			nullFloat64 := fromField.Interface().(sql.NullFloat64)
			if nullFloat64.Valid {
				toField.SetString(strconv.FormatFloat(nullFloat64.Float64, 'f', -1, 64))
			} else {
				toField.SetString("")
			}
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if fromFieldType.Kind() == reflect.String {
			fromVal := fromField.String()
			regex := configs.GetRegexPattern(constants.DecimalZeroOrCommaKey)
			fromVal = regex.ReplaceAllString(fromVal, "")
			intVal, err := strconv.ParseInt(fromVal, 10, toFieldType.Bits())
			if err == nil {
				toField.SetInt(intVal)
			}
		} else if fromFieldType.Kind() == reflect.Uint || fromFieldType.Kind() == reflect.Uint32 || fromFieldType.Kind() == reflect.Uint64 {
			toField.SetInt(int64(fromField.Uint()))
		} else if fromFieldType.AssignableTo(reflect.TypeOf(sql.NullInt64{})) {
			nullInt64 := fromField.Interface().(sql.NullInt64)
			if nullInt64.Valid {
				toField.SetInt(nullInt64.Int64)
			} else {
				toField.SetInt(0)
			}
		} else if fromFieldType.AssignableTo(reflect.TypeOf(sql.NullInt32{})) {
			nullInt32 := fromField.Interface().(sql.NullInt32)
			if nullInt32.Valid {
				toField.SetInt(int64(nullInt32.Int32))
			} else {
				toField.SetInt(0)
			}
		} else if fromFieldType.AssignableTo(reflect.TypeOf(sql.NullInt16{})) {
			nullInt16 := fromField.Interface().(sql.NullInt16)
			if nullInt16.Valid {
				toField.SetInt(int64(nullInt16.Int16))
			} else {
				toField.SetInt(0)
			}
		} else if fromFieldType.Kind() == reflect.Float32 || fromFieldType.Kind() == reflect.Float64 {
			toField.SetInt(int64(fromField.Float()))
		} else if fromFieldType.AssignableTo(reflect.TypeOf(sql.NullFloat64{})) {
			nullFloat64 := fromField.Interface().(sql.NullFloat64)
			if nullFloat64.Valid {
				toField.SetInt(int64(nullFloat64.Float64))
			} else {
				toField.SetInt(0)
			}
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if fromFieldType.Kind() == reflect.String {
			fromVal := fromField.String()
			fromVal = strings.Replace(fromVal, ",", "", -1)
			floatVal, err := strconv.ParseFloat(fromVal, toFieldType.Bits())
			if err == nil {
				uintVal := uint64(floatVal)
				toField.SetUint(uintVal)
			}
		} else if fromFieldType.Kind() == reflect.Interface {
			// Convert the interface value to uint64
			var uintVal uint64
			switch interfaceType := fromField.Interface().(type) {
			case int, int8, int16, int32, int64:
				uintVal = uint64(interfaceType.(int64))
			case uint, uint8, uint16, uint32, uint64:
				uintVal = uint64(interfaceType.(uint64))
			case float32:
				uintVal = uint64(interfaceType)
			case float64:
				uintVal = uint64(interfaceType)
			case string:
				floatVal, err := strconv.ParseFloat(interfaceType, toFieldType.Bits())
				if err == nil {
					uintVal = uint64(floatVal)
				}
			}
			toField.SetUint(uintVal)
		} else if fromFieldType.Kind() == reflect.Int || fromFieldType.Kind() == reflect.Int16 || fromFieldType.Kind() == reflect.Int32 || fromFieldType.Kind() == reflect.Int64 || fromFieldType.Kind() == reflect.Int8 {
			intVal := fromField.Int()
			if intVal >= 0 {
				toField.SetUint(uint64(intVal))
			}
		} else if fromFieldType.Kind() == reflect.Uint || fromFieldType.Kind() == reflect.Uint16 || fromFieldType.Kind() == reflect.Uint32 || fromFieldType.Kind() == reflect.Uint64 {
			uintVal := fromField.Uint()
			toField.SetUint(uintVal)
		} else if fromFieldType.AssignableTo(reflect.TypeOf(sql.NullInt64{})) {
			nullInt64 := fromField.Interface().(sql.NullInt64)
			if nullInt64.Valid {
				toField.SetUint(uint64(nullInt64.Int64))
			} else {
				toField.SetUint(0)
			}
		} else if fromFieldType.AssignableTo(reflect.TypeOf(sql.NullInt32{})) {
			nullInt32 := fromField.Interface().(sql.NullInt32)
			if nullInt32.Valid {
				toField.SetUint(uint64(nullInt32.Int32))
			} else {
				toField.SetUint(0)
			}
		} else if fromFieldType.AssignableTo(reflect.TypeOf(sql.NullInt16{})) {
			nullInt16 := fromField.Interface().(sql.NullInt16)
			if nullInt16.Valid {
				toField.SetUint(uint64(nullInt16.Int16))
			} else {
				toField.SetUint(0)
			}
		} else if fromFieldType.Kind() == reflect.Float32 || fromFieldType.Kind() == reflect.Float64 {
			floatVal := fromField.Float()
			toField.SetUint(uint64(floatVal))
		}

	case reflect.Float32, reflect.Float64:
		if fromFieldType.Kind() == reflect.String {
			fromVal := fromField.String()
			fromVal = strings.Replace(fromVal, ",", "", -1)
			floatVal, err := strconv.ParseFloat(fromVal, toFieldType.Bits())
			if err == nil {
				toField.SetFloat(floatVal)
			}
		} else if fromFieldType.Kind() == reflect.Float64 || fromFieldType.Kind() == reflect.Float32 {
			floatVal := fromField.Float()
			toField.SetFloat(floatVal)
		} else if fromFieldType.Kind() == reflect.Uint32 || fromFieldType.Kind() == reflect.Uint64 {
			toField.SetFloat(float64(fromField.Uint()))
		} else if fromFieldType.Kind() == reflect.Int8 || fromFieldType.Kind() == reflect.Int16 || fromFieldType.Kind() == reflect.Int32 || fromFieldType.Kind() == reflect.Int64 {
			toField.SetFloat(float64(fromField.Int()))
		} else if fromFieldType.Kind() == reflect.Float32 || fromFieldType.Kind() == reflect.Float64 {
			toField.SetFloat(fromField.Float())
		} else if fromFieldType == reflect.TypeOf(sql.NullFloat64{}) {
			nullFloat := fromField.Interface().(sql.NullFloat64)
			if nullFloat.Valid {
				toField.SetFloat(nullFloat.Float64)
			} else {
				toField.SetFloat(0)
			}
		}
	}
}

// MapStructArray maps each item in the inputList to the corresponding item in the outputList using the MapStruct function.
func MapStructArray[T1, T2 any](ctx context.Context, inputList []T1, outputList []T2) {
	ctx, span := tracer.AddToSpan(ctx, "MapStructArray")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	for i, item := range inputList {
		MapStruct(ctx, item, &outputList[i])
	}
}
