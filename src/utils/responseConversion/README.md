# Response Conversion Package
This package provides functionalities for converting field values in structs or slices of structs based on specified conversion rules. It includes methods for converting date formats, normalizing string values, and converting individual fields.
# Functions

## ConvertFieldValues
`func ConvertFieldValues(input interface{}, isRequest bool)`
- **Description**:
Converts the field values of the input struct or slice of structs based on the specified conversion rules.
- **Parameters**:
	- `input`: Input struct or slice of structs.
	- `isRequest`: Boolean indicating whether the conversion is for request or response.
## ConvertDateFormat
`func ConvertDateFormat(inputDate string, inputLayouts []string, outputLayout string) (string, error)`
- **Parameters**:
Converts the input date string to the specified output layout using a list of input layouts.
- **Parameters**:
	- `inputDate`: Input date string.
	- `inputLayouts`: List of input date layouts.
	- `outputLayout`: Output date layout.
- **Returns**:
	- string: Formatted date string.
	- error: Error, if any.
## NormalizeValue
`func NormalizeValue(value string, mapping map[string]string) string`
- **Description**:
Normalizes the given value using the provided mapping.
- **Parameters**:
	- `value`: Value to normalize.
	- `mapping`: Mapping of values for normalization.
	- `string`: Normalized value.

## NormalizeValueArray
`func NormalizeValueArray(values *[]string, mapping map[string]string)`
- **Description**:
Normalizes an array of values using the provided mapping.
- **Parameters**:
	- `values`: Pointer to the array of values to normalize.
	- `mapping`: Mapping of values for normalization.

## SingleFieldConvert

`func SingleFieldConvert(input interface{}, fieldName string, mapStruct map[string]string)`
- **Description**:
Converts the value of the specified field in the input struct to uppercase, normalizes it using the NormalizeValue function, and sets the normalized value back to the field.
- **Parameters**:
	- `input`: Input struct.
	- `fieldName`: Name of the field to convert.
	- `mapStruct`: Mapping of values for normalization.

## Example Usage

```go
package main

import (
	"fmt"
	"your/package/responseConversion"
)

type ExampleStruct struct {
	Name      string
	BirthDate string
	Items     []string
}

func main() {
	exampleSlice := []ExampleStruct{
		{Name: "John", BirthDate: "1990-01-01", Items: []string{"item1", "item2"}},
		{Name: "Alice", BirthDate: "1985-05-20", Items: []string{"item3", "item4"}},
	}

	fmt.Println("Before conversion:")
	for _, item := range exampleSlice {
		fmt.Printf("Name: %s, BirthDate: %s, Items: %v\n", item.Name, item.BirthDate, item.Items)
	}

	responseConversion.ConvertFieldValues(&exampleSlice, false)

	fmt.Println("\nAfter conversion:")
	for _, item := range exampleSlice {
		fmt.Printf("Name: %s, BirthDate: %s, Items: %v\n", item.Name, item.BirthDate, item.Items)
	}
}
```
