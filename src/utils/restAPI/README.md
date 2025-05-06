# REST API Package

## Introduction
The `restAPI` package provides utility functions for interacting with RESTful APIs in Go applications. It includes functions for sending HTTP requests, handling responses, and caching data if required.


### Functions 
1. `HTTPGetRequest(ctx context.Context, baseURL string, requestPayload *models.Request, queryParams url.Values) (interface{}, error)`
	- **Description** : The HTTPGetRequest function sends an HTTP GET request to a specified API endpoint and returns the response data.
	- **Parameters** :
		- `ctx`: The context object.
		- `baseURL`: The base URL of the API endpoint.
		- `requestPayload`: Payload containing information about the request.
		- `queryParams` : query parameters for the request URL.
	- **Returns**:
		- `interface`: The response data.
		- `error`:  Any error encountered during the HTTP request process.

2. `HttpPostRequest (ctx context.Context, baseURL string, requestPayload *models.Request) (interface{}, error)`
	- **Description** : The `HttpPostRequest` function makes a POST request to the specified API endpoint and handles the response
	- **Parameters**:
		- `ctx context.Context`: The context for the HTTP request.
		- `baseURL string`: The base URL of the API endpoint.
		- `requestPayload *models.Request`: The payload containing information about the request.
	- **Returns**:
		- `interface{}`: The decoded response data.
		- `error`: Any error encountered during the HTTP request process.

3. `HttpDeleteRequest(ctx context.Context, baseURL string, requestPayload *models.Request) (interface{}, error)`
	- **Description**: The `HttpDeleteRequest` function sends an DELETE request to the specified API endpoint and handles the response.
	- **Parameters**: 
		- `ctx context.Context`: The context for the HTTP request.
		- `baseURL string`: The base URL of the API endpoint.
		- `requestPayload *models.Request`: The payload containing information about the request.
	- **Returns**:
		- `interface{}`: The decoded response data.
		- `error`: Any error encountered during the HTTP request process.

4. `HttpPatchRequest(ctx context.Context, baseURL string, requestPayload *models.Request) (interface{}, error)`
	- **Description** : The `HttpPatchRequest` function sends an PATCH request to the specified API endpoint and handles the response.
	- **Parameters**: 
		- `ctx context.Context`: The context for the HTTP request.
		- `baseURL string`: The base URL of the API endpoint.
		- `requestPayload *models.Request`: The payload containing information about the request.
	- **Returns**:
		- `interface{}`: The decoded response data.
		- `error`: Any error encountered during the HTTP request process.

5. `HttpRequest(baseUrl string, requestPayload *models.Request) (*http.Response, error)`
	- **Description** : The `HttpRequest` function makes an HTTP request to the given base URL using the provided request payload.
	- **Parameters**: 
		- `baseUrl string`: The base URL of the API endpoint.
		- `requestPayload *models.Request`: The payload containing information about the request.
	- **Returns**:
		- `*http.Response`: The HTTP response returned by the API.
		- `error`: Any error encountered during the HTTP request process.

5. `EncodeQueryParams(params url.Values) string`
	- **Description** : The `EncodeQueryParams` function encodes the given URL parameters into a query string.
	- **Parameters**: 
		`params url.Values`: The URL parameters to be encoded.
	- **Returns**:
		- `string`: The encoded query string.

## Usage Examples 
1. HttpRequest

```go 
// Define base URL and request payload
baseUrl := "https://api.example.com"
requestPayload := &models.Request{ /* Populate request payload */ }

// Make an HTTP request
response, err := restAPI.HttpRequest(baseUrl, requestPayload)
if err != nil {
	// Handle error
}

// Process response
// ...

// Encode query parameters
queryParams := url.Values{ /* Add query parameters */ }
queryString := restAPI.EncodeQueryParams(queryParams)

```
2. Usage using PostRequest example: 

```go 
ctx := context.Background()

// Define base URL and request payload
baseURL := "https://api.example.com"
requestPayload := &models.Request{ /* Populate request payload */ }

// Send HTTP POST request
responseData, err := restAPI.HttpPostRequest(ctx, baseURL, requestPayload)
if err != nil {
	// Handle error
}
```
Similarly we can use the other methods as well.
