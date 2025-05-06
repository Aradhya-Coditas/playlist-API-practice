## cryptoRSA Package

### CryptoRSA

- **Description:** This struct provides RSA encryption and decryption functions.

#### NewCryptoRSA()

- **Description:** Creates a new instance of the CryptoRSA struct.

#### MarshalPEMPublicKeyToString(publicKey *rsa.PublicKey) (string, error)

- **Description:** Converts a PEM-encoded RSA public key to a string.
- **Parameters:**
  - `publicKey`: The RSA public key to be encoded.
- **Return Value:** A string representing the PEM-encoded public key.
- **Example Usage:**
  ```go
  publicKey := &rsa.PublicKey{...}
  pemString, err := cryptoRSAInstance.MarshalPEMPublicKeyToString(publicKey)
  if err != nil {
      // Handle error
  }
  // Use pemString
  ```

#### MarshalPEMPrivateKeyToString(privateKey *rsa.PrivateKey) (string, error)

- **Description:** Converts an RSA private key to a PEM-encoded string.
- **Parameters:**
  - `privateKey`: The RSA private key to be encoded.
- **Return Value:** A string representing the PEM-encoded private key.
- **Example Usage:**
  ```go
  privateKey := &rsa.PrivateKey{...}
  pemString, err := cryptoRSAInstance.MarshalPEMPrivateKeyToString(privateKey)
  if err != nil {
      // Handle error
  }
  // Use pemString
  ```

#### GetPublicKeyFromPEMData(pemPublicKey string) (*rsa.PublicKey, error)

- **Description:** Extracts an RSA public key from a PEM-encoded string.
- **Parameters:**
  - `pemPublicKey`: The PEM-encoded RSA public key string.
- **Return Value:** The RSA public key.
- **Example Usage:**
  ```go
  pemPublicKey := "-----BEGIN PUBLIC KEY-----..."
  publicKey, err := cryptoRSAInstance.GetPublicKeyFromPEMData(pemPublicKey)
  if err != nil {
      // Handle error
  }
  // Use publicKey
  ```

#### CreateRSAKeys(publicKeyWriter, privateKeyWriter io.Writer, keyLength int) error

- **Description:** Generates RSA public and private key pairs and writes them to the provided writers.
- **Parameters:**
  - `publicKeyWriter`: The writer to which the public key will be written.
  - `privateKeyWriter`: The writer to which the private key will be written.
  - `keyLength`: The length of the RSA key in bits.
- **Return Value:** An error if key generation or writing fails.
- **Example Usage:**
  ```go
  publicKeyWriter := bytes.NewBuffer([]byte{})
  privateKeyWriter := bytes.NewBuffer([]byte{})
  err := cryptoRSAInstance.CreateRSAKeys(publicKeyWriter, privateKeyWriter, 2048)
  if err != nil {
      // Handle error
  }
  ```

#### ReadPublicKeyFromPEMFile(filePath string) (*rsa.PublicKey, error)

- **Description:** Reads an RSA public key from a PEM-encoded file.
- **Parameters:**
  - `filePath`: The path to the PEM-encoded public key file.
- **Return Value:** The RSA public key.
- **Example Usage:**
  ```go
  publicKey, err := cryptoRSAInstance.ReadPublicKeyFromPEMFile("/path/to/public_key.pem")
  if err != nil {
      // Handle error
  }
  // Use publicKey
  ```

#### SavePublicKeyToFile(publicKey *rsa.PublicKey, filePath string) error

- **Description:** Saves the RSA public key to a PEM-encoded file.
- **Parameters:**
  - `publicKey`: The RSA public key to be saved.
  - `filePath`: The path to the output file.
- **Return Value:** An error if saving fails.
- **Example Usage:**
  ```go
  publicKey := &rsa.PublicKey{...}
  err := cryptoRSAInstance.SavePublicKeyToFile(publicKey, "/path/to/public_key.pem")
  if err != nil {
      // Handle error
  }
  ```

#### SaveRSAPrivateKeyPKCS8ToFile(privateKey *rsa.PrivateKey, filePath string) error

- **Description:** Saves an RSA private key in PKCS#8 format to a PEM-encoded file.
- **Parameters:**
  - `privateKey`: The RSA private key to be saved.
  - `filePath`: The path to the output file.
- **Return Value:** An error if saving fails.
- **Example Usage:**
  ```go
  privateKey := &rsa.PrivateKey{...}
  err := cryptoRSAInstance.SaveRSAPrivateKeyPKCS8ToFile(privateKey, "/path/to/private_key.pem")
  if err != nil {
      // Handle error
  }
  ```

#### GetPrivateKeyFromPKCS8PEMData(pemPrivateKey string) (*rsa.PrivateKey, error)

- **Description:** Extracts an RSA private key from a PEM-encoded string.
- **Parameters:**
  - `pemPrivateKey`: The PEM-encoded RSA private key string.
- **Return Value:** The RSA private key.
- **Example Usage:**
  ```go
  pemPrivateKey := "-----BEGIN PRIVATE KEY-----..."
  privateKey, err := cryptoRSAInstance.GetPrivateKeyFromPKCS8PEMData(pemPrivateKey)
  if err != nil {
      // Handle error
  }
  // Use privateKey
  ```

#### CreateBFFKeyPair(log logger.Logger) (*rsa.PrivateKey, *rsa.PublicKey, error)

- **Description:** Generates a pair of RSA keys - a private key and a public key.
- **Parameters:**
  - `log`: The logger instance for logging errors.
- **Return Value:** The RSA private key and public key pair.
- **Example Usage:**
  ```go
  privateKey, publicKey, err := cryptoRSAInstance.CreateBFFKeyPair(loggerInstance)
  if err != nil {
      // Handle error
  }
  // Use privateKey and publicKey
  ```

### Encryption Functions
    These functions are used to encrpyt/convert the plain text into no human readable format.

#### EncryptBlock

- **Description:** Encrypts the provided plaintext using the RSA public key.
- **Parameters:**
  - `publicKey`: The RSA public key used for encryption.
  - `plainText`: The plaintext string to be encrypted.
- **Return Value:** The base64-encoded ciphertext string.

#### Encrypt

- **Description:** Encrypts a given plaintext string using RSA encryption with a custom block size.
- **Parameters:**
  - `plainText` (string): The plaintext string to be encrypted.
  - `publicKey` (*rsa.PublicKey): The RSA public key used for encryption.
  - `keySize` (int): The size of the RSA key in bits.
- **Return Value:** (string, error) - The base64-encoded ciphertext string and an error if encryption fails.
- **Example Usage:**
  ```go
  publicKey := &rsa.PublicKey{...}
  cipherText, err := cryptoRSAInstance.Encrypt(plainText, publicKey, keySize)
  if err != nil {
      // Handle error
  }
  ```

#### General Encryption Flow

- **Encryption function** first retrieves the Postgres encryption key from the configuration and pads it to the correct length for AES encryption.
- It then checks the type of the input (reflect.String, reflect.Slice, reflect.Struct, etc.).
- Based on the input type, it calls the corresponding helper function:
  - `encryptString` for strings.
  - `encryptSlice` for slices or arrays.
  - `encryptStruct` for structs with tagged fields.
-The helper functions perform AES-GCM encryption and return the encrypted result or an error.

#### Encryption

- **Description:** This function encrypts a string, slice, or struct with tagged fields using AES-GCM encryption. It dynamically detects the type of input and calls the appropriate helper function (encryptString, encryptSlice, or encryptStruct) to perform encryption.
- **Parameters:**
  - `ctx`(context.Context): The context for the operation, including trace spans.
  - `input`(interface{}): The input value to be encrypted. Can be a string, slice, array, or struct.
- **Return Value:** (interface{}, error) - The encrypted data, or an error if encryption fails.
- **Example Usage:**
```go
input := "sensitive data"
encryptedData, err := Encryption(ctx, input)
if err != nil {
    // Handle error
}
  ```

#### encryptString

- **Description:** This function encrypts a string, slice, or struct with tagged fields using AES-GCM encryption. It dynamically detects the type of input and calls the appropriate helper function (encryptString, encryptSlice, or encryptStruct) to perform encryption.
- **Parameters:**
  - `ctx`(context.Context): The context for the operation.
  - `plainText`(string): The string to be encrypted.
  - `keyBytes`([]byte): The AES encryption key, padded to the correct length.
- **Return Value:** (string, error) - The base64-encoded ciphertext string, or an error if encryption fails.
- **Example Usage:**
```go
encryptedStr, err := encryptString(ctx, "secret", keyBytes)
if err != nil {
    // Handle error
}
  ```

#### encryptSlice

- **Description:** This function encrypts a string, slice, or struct with tagged fields using AES-GCM encryption. It dynamically detects the type of input and calls the appropriate helper function (encryptString, encryptSlice, or encryptStruct) to perform encryption.
- **Parameters:**
  - `ctx`(context.Context): The context for the operation.
  - `val`(reflect.Value): The slice or array to be encrypted.
- **Return Value:** (interface{}, error) - The encrypted slice or array, or an error if encryption fails.
- **Example Usage:**
```go
inputSlice := []string{"data1", "data2"}
encryptedSlice, err := encryptSlice(ctx, reflect.ValueOf(inputSlice))
if err != nil {
    // Handle error
}
  ```

#### encryptStruct

- **Description:** Encrypts a struct by processing fields tagged with the encryption tag (e.g., db:"crypt"). Only fields marked with the tag are encrypted.
- **Parameters:**
  - `ctx`(context.Context): The context for the operation.
  - `val`(reflect.Value): The struct to be encrypted.
  - `key`([]byte): The AES encryption key.
- **Return Value:** (interface{}, error) - The struct with encrypted fields, or an error if encryption fails.
- **Example Usage:**
```go
type User struct {
    Name     string `db:"crypt"`
    Password string `db:"crypt"`
}
user := User{Name: "John", Password: "password123"}
encryptedUser, err := encryptStruct(ctx, reflect.ValueOf(user), keyBytes)
if err != nil {
    // Handle error
}
  ```

### Decrypt Functions

The `Decrypt` function in the `cryptoRSA` package is responsible for decrypting multi-block ciphertext into plaintext using the provided RSA private key. It decodes the base64-encoded ciphertext, splits it into individual blocks, decrypts each block using the private key, and concatenates the decrypted blocks to generate the plaintext.

#### DecryptBlock Function

- **Description:** Decrypts the provided ciphertext using the provided RSA private key. It decodes the base64-encoded ciphertext, then uses the RSA PKCS#1 v1.5 padding scheme to decrypt the ciphertext.
- **Parameters:**
  - `privateKey` (*rsa.PrivateKey): The RSA private key used for decryption.
  - `cipherText` (string): The base64-encoded ciphertext string to be decrypted.
- **Return Value:** (string, error) - The plaintext string obtained after decryption and an error if decryption fails.

#### Decrypt Function

- **Description:** Decrypts multi-block ciphertext into plaintext using the provided RSA private key. It decodes the base64-encoded ciphertext, splits it into individual blocks, and decrypts each block using the DecryptBlock method.
- **Parameters:**
  - `encryptedText` (string): The base64-encoded ciphertext string to be decrypted.
  - `privateKey` (*rsa.PrivateKey): The RSA private key used for decryption.
- **Return Value:** (string, error) - The plaintext string obtained after decryption and an error if decryption fails.
- **Example Usage:**
  ```go
  encryptedText := "..." // base64-encoded ciphertext
  privateKey := &rsa.PrivateKey{...} // RSA private key
  plainText, err := cryptoRSAInstance.Decrypt(encryptedText, privateKey)
  if err != nil {
      // Handle decryption error
  }
  // Use plainText
  ```
  
#### General Decryption Flow

- **Decryption function** first retrieves the Postgres decryption key from the configuration and pads it to the correct length for AES decryption.
- It then checks the type of the input (reflect.String, reflect.Slice, reflect.Struct, etc.).
- Based on the input type, it calls the corresponding helper function:
  - `decryptString` for strings.
  - `decryptSlice` for slices or arrays.
  - `decryptStruct` for structs with tagged fields.
-The helper functions perform AES-GCM decryption and return the decrypted result or an error.

#### Decryption

- **Description:** This function decrypts a string, slice, or struct with tagged fields using AES-GCM decryption. It dynamically detects the type of input and calls the appropriate helper function (decryptString, decryptSlice, or decryptStruct) to perform decryption.
- **Parameters:**
  - `ctx`(context.Context): The context for the operation, including trace spans.
  - `input`(interface{}): The encrypted data to be decrypted. Can be a string, slice, array, or struct.
- **Return Value:** (interface{}, error) - The decrypted data, or an error if decryption fails.
- **Example Usage:**
```go
encryptedData := "encryptedText"
decryptedData, err := Decryption(ctx, encryptedData)
if err != nil {
    // Handle error
}
  ```

#### decryptString

- **Description:** Decrypts a base64-encoded AES-GCM ciphertext string into plaintext.
- **Parameters:**
  - `ctx`(context.Context): The context for the operation.
  - `cipherText`(string): The base64-encoded ciphertext to be decrypted.
  - `keyBytes`([]byte): The AES decryption key, padded to the correct length.
- **Return Value:** (string, error) - The decrypted plaintext string, or an error if decryption fails.
- **Example Usage:**
```go
decryptedStr, err := decryptString(ctx, "cipherText", keyBytes)
if err != nil {
    // Handle error
}
  ```

#### decryptSlice

- **Description:** Decrypts each element of a slice or array. Supports recursive decryption of slice elements if they are strings, slices, or structs with tagged fields.
- **Parameters:**
  - `ctx`(context.Context): The context for the operation.
  - `val`(reflect.Value): The slice or array to be decrypted.
- **Return Value:** (interface{}, error) - The decrypted slice or array, or an error if decryption fails.
- **Example Usage:**
```go
encryptedSlice := []string{"encrypted1", "encrypted2"}
decryptedSlice, err := decryptSlice(ctx, reflect.ValueOf(encryptedSlice))
if err != nil {
    // Handle error
}
  ```

#### decryptStruct

- **Description:** Decrypts a struct by processing fields tagged with the decryption tag (e.g., db:"crypt"). Only fields marked with the tag are decrypted.
- **Parameters:**
  - `ctx`(context.Context): The context for the operation.
  - `val`(reflect.Value): The struct to be decrypted.
  - `key`([]byte): The AES decryption key.
- **Return Value:** (interface{}, error) - The struct with decrypted fields, or an error if decryption fails.
- **Example Usage:**
```go
type User struct {
    Name     string `db:"crypt"`
    Password string `db:"crypt"`
}
encryptedUser := User{Name: "encryptedName", Password: "encryptedPassword"}
decryptedUser, err := decryptStruct(ctx, reflect.ValueOf(encryptedUser), keyBytes)
if err != nil {
    // Handle error
}
  ```
  
### Hashing functions 

#### HashData function
 - **Description:** Takes a string of data as input and returns its SHA-256 hash as a hexadecimal string.
    It uses the SHA-256 hashing algorithm provided by the crypto package in Go.
  
  - **Parameters:**
    - `data` (string): The input data to be hashed.
  - **Return Value:**  (string, error) - 
    - `string`: The SHA-256 hash of the input data represented as a hexadecimal string.
    - `error`: An error if hashing fails.
  
 - **Example Usage:**
    ```go
    data := "example"
    hashedData, err := HashData(data)
    if err != nil {
        // Handle error
    }
    ```
#### IterativeStringHashing function

- **Description:** Performs iterative string hashing using the SHA-256 algorithm.
    It returns the hashed password as a hexadecimal string.
    It uses the SHA-256 hashing algorithm provided by the crypto package in Go.

  - **Parameters:**
    - `data` (string): The input data (password) to be hashed iteratively.

  - **Return Value:** 
    - `string`: The SHA-256 hash of the input data (password) represented as a hexadecimal string.
    - `error`: An error if hashing fails.

  - **Example Usage:**
    ```go
    password := "password"
    hashedPassword, err := IterativeStringHashing(password)
    if err != nil {
        // Handle error
    }
    ```
#### ReadHashFromFile Function

- **Description:** Reads the hash data from a file located at the specified file path.

  - **Parameters:**
    - `filePath` (string): The path to the file from which hash data will be read.

  - **Return Value:** 
    - `string`: The hash data read from the file.
    - `error`: An error if reading fails.

  - **Example Usage:**
    ```go
    filePath := "/path/to/file.txt"
    hashData, err := ReadHashFromFile(filePath)
    if err != nil {
        // Handle error
    }
    ```

#### SaveHashToFile function
- **Description:** Saves the provided hash data to a file at the specified file path.

  - **Parameters:**
    - `filePath` (string): The path to the file where the hash data will be saved.
    - `data` (string): The hash data to be saved to the file.

  - **Return Value:** 
    - `error`: An error if saving fails.

  - **Example Usage:**
    ```go
    filePath := "/path/to/file.txt"
    hashData := "hashed_data"
    err := SaveHashToFile(filePath, hashData)
    if err != nil {
        // Handle error
    }
    ```

### RSA Key Parsing Functions
The RSA key parsing functions in the `cryptoRSA` package facilitate the parsing of RSA keys from various sources such as byte slices, files, and PEM-encoded strings.

#### ParsePrivateKey Function

- **Description:** Parses a PEM-encoded RSA private key from a byte slice and returns it as a parsed RSA private key.
- **Parameters:**
  - `privateKey` ([]byte): The PEM-encoded RSA private key byte slice.
- **Return Value:** (*rsa.PrivateKey, error) - The parsed RSA private key and an error if parsing fails.

#### ParseRSAPrivateKeyPKCS8FromFile Function

- **Description:** Reads a PEM-encoded RSA private key from a file specified by `filePath` and returns it as a parsed RSA private key.
- **Parameters:**
  - `filePath` (string): The path to the PEM-encoded RSA private key file.
- **Return Value:** (*rsa.PrivateKey, error) - The parsed RSA private key and an error if parsing fails.

#### LoadPublicKeyFromFile Function

- **Description:** Loads a public key from a PEM file and optionally a hash from a separate file, and returns the public key and the hashed public key.
- **Parameters:**
  - `publicKeyPath` (string): The path to the PEM-encoded public key file.
  - `publicKeyHashPath` (string): The path to the file containing the hash of the public key (optional).
- **Return Value:** (*genericModels.NestKeyPair, error) - The public key pair (`PublicKey` and `PublicHashedKey`) and an error if loading fails.

#### ParsePublicKey Function

- **Description:** Parses a public key string and returns the public key and hashed public key.
- **Parameters:**
  - `publicKey` (string): The base64-encoded PEM public key string.
- **Return Value:** (*genericModels.NestKeyPair, error) - The parsed public key pair (`PublicKey` and `PublicHashedKey`) and an error if parsing fails.

#### ParseKey Example Usage:

```go
// Example for parsing a PEM-encoded RSA private key from a byte slice
privateKeyBytes := []byte("...") // PEM-encoded RSA private key byte slice
parsedPrivateKey, err := cryptoRSA.ParsePrivateKey(privateKeyBytes)
if err != nil {
    // Handle error
}
// Use parsedPrivateKey

// Example for reading a PEM-encoded RSA private key from a file
privateKeyFilePath := "/path/to/private_key.pem"
parsedPrivateKeyFromFile, err := cryptoRSA.ParseRSAPrivateKeyPKCS8FromFile(privateKeyFilePath)
if err != nil {
    // Handle error
}
// Use parsedPrivateKeyFromFile

// Example for loading a public key from a PEM file along with its hash
publicKeyFilePath := "/path/to/public_key.pem"
publicKeyHashFilePath := "/path/to/public_key_hash.txt"
keyPair, err := cryptoRSA.LoadPublicKeyFromFile(publicKeyFilePath, publicKeyHashFilePath)
if err != nil {
    // Handle error
}
// Use keyPair.PublicKey and keyPair.PublicHashedKey

// Example for parsing a base64-encoded PEM public key string
base64EncodedPublicKey := "..." // Base64-encoded PEM public key string
parsedKeyPair, err := cryptoRSA.ParsePublicKey(base64EncodedPublicKey)
if err != nil {
    // Handle error
}
// Use parsedKeyPair.PublicKey and parsedKeyPair.PublicHashedKey
```