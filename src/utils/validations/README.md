# Validation Package

This package provides generic validation functions for the request body and custom validation functions for various fields.

## Define Variable

Define `customErrorMap` which maps validator package's key used in service's model to user defined error message.

## Generic Functions

1. `FormatValidationErrors(validationErrors validator.ValidationErrors) ([]models.ErrorMessage, string)`

- **Description**: Formats validation errors into a user-friendly format.
- **Parameters**:
  - `validationErrors`: An instance of `validator.ValidationErrors`.
- **Returns**:
  - `[]models.ErrorMessage, string`: a list of `models.ErrorMessage` and a concatenated string of error messages.
  
#### Pseudo Code

```go
func FormatValidationErrors(validationErrors validator.ValidationErrors) ([]models.ErrorMessage, string) {
  // Define variables :errorMessages,errorMessagesString
  // Get the flag `useFrontendErrorFormat` from config
  // Loop over validationErrors
  for _, err := range validationErrors {
    // Get the err.Tag() and err.Param() in variables errorMessage and errorParam respectively
    // Check if errorMessage presents in customErrorMap
    if ok {
      // Get the corresponding errMsg in errorMessage
    }
    // Prepare errorMessagesString by appending all the errorMessage in a list
    // Check the 'useFrontendErrorFormat' flag
    if false {
      // Prepare errorMessages by appending errorMessages with key as err.Field()
    }
  }
  // Prepare errorMessagesJoined by concatenating all the errorMessagesString elements
  // Check the 'useFrontendErrorFormat' flag
  if true {
    // Prepare errorMessages by appending errorMessages with key:`Generic`
  }
  // Return formatted error messages and a joined string of error messages.
  return errorMessages, errorMessagesJoined
}
```

2. `PrepareNestValidationErrors(key, errMsg string) ([]models.ErrorMessage, string)`

- **Description**: Prepares Nest validation errors with key and error message.
- **Parameters**:
  - `key`: Key associated with the error.
  - `errMsg`: Error message.
- **Returns**:
  - `[]models.ErrorMessage, string`: a list of `models.ErrorMessage` and a concatenated string of error messages.

#### Pseudo Code

```go
func PrepareNestValidationErrors(key, errMsg string) ([]models.ErrorMessage, string) {
  // Get the flag `useFrontendErrorFormat` from config
  // Define variable :errorMessages
  if key != "" && key != "Generic"{
  // Prepare errMsg String with key 
  } else {
    // Assign key as `Generic`
  }
  // Check the 'useFrontendErrorFormat' flag
  if false {
    errorMessages = append(errorMessages, models.ErrorMessage{Key: key, ErrorMessage: errMsg})
    // Prepare errorMessages by appending errorMessages with received key
  } else {
    errorMessages = append(errorMessages, models.ErrorMessage{Key: genericConstants.GenericErrorKey, ErrorMessage: errMsg})
    // Prepare errorMessages by appending errorMessages with key:`Generic`
  }
    
  // Return the error messages and the final error message string.
  return errorMessages, errMsg
}
```

## Custom Validations

1. `ValidateEnum[E Enum](fl validator.FieldLevel) bool`

- **Description**:Checks if the value of a field implements the Enum interface and is valid.
- **Parameters**:
  - `fl`: An instance of `validator.FieldLevel`.
- **Returns**:
  -  a boolean indicating whether the value adheres to the Enum interface definition.

2. `RetentionDateValidation(fl validator.FieldLevel) bool`

- **Description**: Checks if the value of a field is a valid retention date format.
- **Parameters**:
  - `fl`: An instance of `validator.FieldLevel`.
- **Returns**:
  -  a boolean indicating whether the retention date format is valid.

3. `ScannerTypeValueValidation(fl validator.FieldLevel) bool`

- **Description**: Checks if the value of a field is a valid scanner type value.
- **Parameters**:
  - `fl`: An instance of `validator.FieldLevel`.
- **Returns**:
  -  a boolean indicating whether the scanner type value is valid.

4. `PANValidation(fl validator.FieldLevel) bool`

- **Description**: Validates PAN (Permanent Account Number) format.
- **Parameters**:
  - `fl`: An instance of `validator.FieldLevel`.
- **Returns**:
  - a boolean indicating whether the PAN format is valid.

5. `DateOfBirthValidation(fl validator.FieldLevel) bool`

- **Description**: Validates date of birth format and checks if it's not in the future.
- **Parameters**:
  - `fl`: An instance of `validator.FieldLevel`.
- **Returns**:
  -  a boolean indicating whether the date of birth format is valid and it's not in the future.

6. `BidLengthValidation(fl validator.FieldLevel) bool`

- **Description**: Validates if a slice field has a length greater than zero.
- **Parameters**:
  - `fl`: An instance of `validator.FieldLevel`.
- **Returns**:
  -  a boolean indicating whether the slice field has a length greater than zero.
