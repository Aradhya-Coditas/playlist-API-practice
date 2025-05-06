# Response Package

## Introduction
The responseUtils package provides a set of utility functions to manage HTTP responses in a Go web application using the Gin framework. It streamlines the process of sending consistent JSON responses for different HTTP status codes and includes functionality for conditionally omitting fields in response data based on custom tags and request sources.

## Usage
The package offers various functions to send HTTP responses and manipulate response data.

## Functions

### SendBadRequest
This function formats and sends a 400 Bad Request response using the Gin framework. This function is useful when the client has sent invalid data or parameters.
### SendUnauthorized
This function sends a 401 Unauthorized response with a parsed error message in a Gin context.
### SendInternalServerError
This Go function handles internal server errors in a web application using Gin, extracting and sending appropriate error messages in JSON format with a status code of 500.
### SendNoContentFoundError
This function sends a 204 No Content response using the Gin context.
### SendConflict
This Go code defines a function called SendConflict that adds an error message to the context and returns a JSON response with the error message and status code 409.
### SendStatusOK
This Go function SendStatusOK processes data, sets it in the context, and sends a JSON response with status 200.
### SendCreated
The function that sends a response with status code 201 (Created) and includes the provided data.
### SendStatusUnprocessableEntity
The function sends a JSON response with status code 422 (Unprocessable Entity) using provided data in a Gin context.
### SendAbortWithStatusJSON
This function is designed to be used in middleware. It appends an error to the context's errors, then aborts with a JSON response containing the error message and status information.
### SendForbidden
This function sends a 403 Forbidden response with a specific error message extracted from the input error.
### StructHide
`StructHide` is a utility for conditionally omitting fields in a JSON response based on custom tags and the request source.
#### hasHideTag
The `hasHideTag `function checks if a struct field has a specific tag indicating that it should be hidden from the JSON response. It supports dynamic response customization.
#### handleFields
The `handleFields` function processes struct fields, applying logic to include or exclude them based on tags and request context. It assists in generating tailored JSON responses.

Example:

```go
package main

import (
    "fmt"
    "your/package/responseUtils"
)

type UserResponse struct {
    UserId      string  `json:"userId" hide:"MOB"`
    Name        string  `json:"name"`
    Age         int     `json:"age"`
}

func main() {
    userResponse := UserResponse{
        UserId: "123",
        Name:   "John Doe",
        Age:    30,
    }
    
    data := responseUtils.StructHide(userResponse, "MOB")
    fmt.Println(data)
}

```