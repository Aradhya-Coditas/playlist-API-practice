package cryptoRSA

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/tracer"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

// DecryptBlock decrypts cipherText using the provided RSA private key.
// It decodes the cipherText from base64, then uses the RSA PKCS#1 v1.5 padding scheme to decrypt the ciphertext.
func (c *CryptoRSA) DecryptBlock(ctx context.Context, privateKey *rsa.PrivateKey, cipherText string) (string, error) {
	_, span := tracer.AddToSpan(ctx, "DecryptBlock")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)

	if err != nil {
		return "", errors.New(genericConstants.DecryptionError)
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherTextBytes)
	if err != nil {
		return "", errors.New(genericConstants.DecryptionError)
	}

	return string(plainText), nil
}

// Decrypt decrypts the multi-block ciphertext into plaintext using the provided RSA private key.
// The method first decodes the encryptedText from base64, then splits the multi-block ciphertext into individual blocks.
// For each block, it trims any leading or trailing whitespace and checks if the block is not empty.
// If the block is not empty, it calls the DecryptBlock method to decrypt the block using the privateKey.
// The decrypted block is then concatenated to the plainText.
// Finally, the method returns the concatenated plainText as a string and any error encountered during the decryption process.
func (c *CryptoRSA) Decrypt(ctx context.Context, encryptedText string, privateKey *rsa.PrivateKey) (string, error) {
	ctx, span := tracer.AddToSpan(ctx, "Decrypt")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	decodedCipherText, err := base64.StdEncoding.DecodeString(encryptedText)

	if err != nil {
		return "", fmt.Errorf(genericConstants.DecodeCipherTextError, err)
	}

	blockCipherTexts := strings.Split(string(decodedCipherText), "\n")
	var plainText string

	for _, blockCipherText := range blockCipherTexts {
		blockCipherText = strings.TrimSpace(blockCipherText)

		if blockCipherText != "" {
			plainBlock, err := c.DecryptBlock(ctx, privateKey, blockCipherText)

			if err != nil {
				return "", fmt.Errorf(genericConstants.DecryptEncryptedBlockError, err)
			}

			plainText += plainBlock
		}
	}

	return plainText, nil
}

// Decryption decrypts a string, slice, or struct with tagged fields.
func Decryption(ctx context.Context, input interface{}) (interface{}, error) {
	ctx, span := tracer.AddToSpan(ctx, "Decryption")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	val := reflect.ValueOf(input)

	// Retrieve Postgres decryption key
	PostgresConfig, err := configs.Get(genericConstants.PostgresConfig)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.GetPostgresConfigError, err)
	}
	key := []byte(PostgresConfig.GetString(genericConstants.PostgresSecretKey))
	key = utils.PadKey(key)

	switch val.Kind() {
	case reflect.String:
		return decryptString(ctx, val.String(), key)

	case reflect.Slice, reflect.Array:
		return decryptSlice(ctx, val)

	case reflect.Struct:
		return decryptStruct(ctx, val, key)

	default:
		return nil, fmt.Errorf(genericConstants.EncryptDecryptUnsupportedTypeError, val.Kind())
	}
}

// decryptSlice handles decryption of slices or arrays
func decryptSlice(ctx context.Context, val reflect.Value) (interface{}, error) {
	_, span := tracer.AddToSpan(ctx, "decryptSlice")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	decryptedSlice := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		decryptedElem, err := Decryption(ctx, elem.Interface())
		if err != nil {
			return nil, fmt.Errorf(genericConstants.DecryptionProcessError, err)
		}
		decryptedSlice.Index(i).Set(reflect.ValueOf(decryptedElem))
	}
	return decryptedSlice.Interface(), nil
}

// decryptStruct handles decryption of structs with tagged fields
func decryptStruct(ctx context.Context, val reflect.Value, key []byte) (interface{}, error) {
	_, span := tracer.AddToSpan(ctx, "decryptStruct")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	output := reflect.New(val.Type()).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		dbTag := fieldType.Tag.Get(genericConstants.EncryptDecryptTagName)
		if dbTag == genericConstants.EncryptDecryptTagValue {
			switch field.Kind() {
			case reflect.String:
				decryptedValue, err := decryptString(ctx, field.String(), key)
				if err != nil {
					return nil, fmt.Errorf(genericConstants.DecryptionProcessError, err)
				}
				output.Field(i).SetString(decryptedValue)
			case reflect.Slice, reflect.Array:
				decryptedSlice, err := decryptSlice(ctx, field)
				if err != nil {
					return nil, err
				}
				output.Field(i).Set(reflect.ValueOf(decryptedSlice))
			default:
				output.Field(i).Set(field)
			}
		} else {
			output.Field(i).Set(field)
		}
	}
	return output.Interface(), nil
}

// decryptString decrypts a plain string using AES-GCM.
func decryptString(ctx context.Context, cipherText string, keyBytes []byte) (string, error) {
	_, span := tracer.AddToSpan(ctx, "decryptString")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	cipherTextBytes, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf(genericConstants.FailedToDecodeBase64Error, err)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf(genericConstants.FailedToCreateCipherError, err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf(genericConstants.FailedToCreateGCMError, err)
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherTextBytes := cipherTextBytes[:nonceSize], cipherTextBytes[nonceSize:]

	plainTextBytes, err := gcm.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", fmt.Errorf(genericConstants.DecryptionError, err)
	}

	return string(plainTextBytes), nil
}
