# Authorization Package

This package provides utilities for handling JWT tokens and related cryptographic operations.

## Overview

The `authorization` package contains functions for parsing, creating, encrypting, and decrypting JWT tokens, as well as compressing and decompressing token data.

## Dependencies

Before running the backend services, make sure you have the following installed:

- [Sonic](https://github.com/bytedance/sonic) : Used for JSON encoding and decoding.
- [Jwt-go](https://github.com/dgrijalva/jwt-go) : Used for JWT token parsing and creation.
- [zstd](https://pkg.go.dev/github.com/klauspost/compress@v1.17.4/zstd) : Used for Zstandard compression.

## JwtTokenUtils Interface

The `JwtTokenUtils` interface defines methods for parsing, creating, encrypting, decrypting, compressing, and decompressing JWT tokens and token data.

```go
type JwtTokenUtils interface {
	// Methods
	ParseJwtToken(tokenString string, secretKey string) (jwt.MapClaims, error) // Parses a JWT token string using the provided secret key and returns the token claims.
	CreateJwtToken(tokenData string, expiryDays int, tokenType string, secretKey string) (string, error) // Creates a new JWT token with the provided token data, expiry duation, token type, and secret key.
	DecryptTokenData(encryptedTokenData string) (*models.TokenData, error) // Decrypts the encrypted token data and returns the token data model.
	EncryptTokenData(tokenDataString string) (string, error) // Encrypts the token data string and returns the encrypted data.
	CompressTokenData(tokenData string) (string, error) // Compresses the token data string using zstd compression.
	DecompressTokenData(compressedTokenData []byte) (string, error) // Decompresses the compressed token data using zstd decompression.
}
```

## jwtTokenUtils Struct

The `jwtTokenUtils` struct implements the `JwtTokenUtils` interface. It provides the concrete implementations for the methods defined in the interface.

## Constructor

- `NewJwtTokenUtils() *jwtTokenUtils`: Creates a new instance of the `jwtTokenUtils` struct.


## Functions

1. `ParseJwtToken(tokenString string, secretKey string) (jwt.MapClaims, error)`

- **Parameters**: 
  - `tokenString` (string): The JWT token string to be parsed.
  - `secretKey` (string): The secret key used for token validation.
- **Returns**: 
  - `jwt.MapClaims`: A map of claims extracted from the JWT token.
  - `error`: An error if the token parsing or validation fails.

2. `CreateJwtToken(tokenData string, expiryDays int, tokenType string, secretKey string) (string, error)`

- **Parameters**: 
  - `tokenData` (string): The data to be included in the token payload.
  - `expiryDays` (int): The number of days until the token expires.
  - `tokenType` (string): The type of the token.
  - `secretKey` (string): The secret key used for token signing.
- **Returns**: 
  - `string`: The signed JWT token string.
  - `error`: An error if token creation fails.

3. `DecryptTokenData(encryptedTokenData string) (*models.TokenData, error)`

- **Parameters**: 
  - `encryptedTokenData` (string): The encrypted token data to be decrypted.
- **Returns**: 
  - `*models.TokenData`: A pointer to the decrypted token data model.
  - `error`: An error if token decryption fails.

4. `EncryptTokenData(tokenDataString string) (string, error)`

- **Parameters**: 
  - `tokenDataString` (string): The token data string to be encrypted.
- **Returns**: 
  - `string`: The encrypted token data.
  - `error`: An error if token encryption fails.
  
5. `CompressTokenData(tokenDataString string) (string, error)`

- **Parameters**: 
  - `tokenDataString` (string): The token data string to be compressed.
- **Returns**: 
  - `string`: The compressed token data.
  - `error`: An error if token compression fails.

6. `DecompressTokenData(compressedTokenData []byte) (string, error)`

- **Parameters**: 
  - `compressedTokenData` ([]byte): The compressed token data to be decompressed.
- **Returns**: 
  - `string`: The decompressed token data.
  - `error`: An error if token decompression fails.
