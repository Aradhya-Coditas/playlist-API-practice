# mapStruct Package

The MapStruct package provides functionality for mapping fields from one struct to another in Go. It includes support for nested structs and handling slice mapping.

## Functions

### MapStruct

The function `MapStruct` that recursively maps the fields from a source struct to a destination struct. It handles mapping of nested structs and slices as well. It iterates over the fields of the source struct, checks compatibility, and sets the values accordingly in the destination struct.

- **Parameters**: 
  - `from interface{}` : This parameter is the source struct from which the fields are mapped.
  - `to interface{}` : This parameter is the address of the destination struct to which the fields are mapped. 

### MapStructArray

The function `MapStructArray` is implemented as a generic function in Go. This means that it can work with slices of any type, as the type parameters `T1` and `T2` are defined with the `any` constraint. The function iterates over the elements in the inputList and calls the `MapStruct` function to map each item from the `inputList` to the corresponding item in the `outputList`. This flexibility in type handling allows the function to be reused with different types of slices without the need to rewrite the function for each specific type.

- **Parameters**: 
  - `inputList` : This parameter is a slice of type T1, which is the type of elements in the input list.
  - `outputList` : This parameter is a slice of type T2, which is the type of elements in the output list.


### mapSlice

 The function `mapSlice` that maps values from one slice to another based on their types.
 
- **Parameters**: 
  - `fromField` : This is a `reflect.Value` representing the source slice from which values will be mapped.
  - `toField` : This is a `reflect.Value` representing the target slice where the mapped values will be stored.

### mapField

 The function `mapField` that maps the value of a source data field to a destination data field based on their types.It handles conversions between different data types such as string, int, uint, float32, and float64 by using reflection in Go.

- **Parameters**: 
  - `fromField` : This is a `reflect.Value` representing the source data field that needs to be mapped to the destination field.
  - `toField` : This is a `reflect.Value` representing destination data field where the value from the source field will be mapped to.


## Usage

To integrate the `MapStruct` library into your application, import the `mapstruct` package.

Example:

```go
package main

import (
	"fmt"
	"your/package/mapStruct"
)

type BFFUser struct {
    Name string
    Age  int
}

type NestUser struct {
    Name string
    Age  String
}

func main() {
    var bffUser BFFUser
    nestUser := NestUser{Name: "Alice", Age: "30"}

    mapStruct.MapStruct(bffUser, &nestUser)
    
    fmt.println(bffUser)
}
