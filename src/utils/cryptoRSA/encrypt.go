package cryptoRSA

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io"

	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/tracer"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

// EncryptBlock encrypts the plainText using the provided RSA public key.
func (c *CryptoRSA) EncryptBlock(ctx context.Context, publicKey *rsa.PublicKey, plainText string) (string, error) {
	_, span := tracer.AddToSpan(ctx, "EncryptBlock")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	data := []byte(plainText)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		return "", fmt.Errorf(genericConstants.EncryptionError, err)
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Encrypt is a method that encrypts a given plaintext string using RSA encryption with a custom block size.
func (c *CryptoRSA) Encrypt(ctx context.Context, plainText string, publicKey *rsa.PublicKey, keySize int) (string, error) {
	ctx, span := tracer.AddToSpan(ctx, "Encrypt")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	src := []byte(plainText)
	buffer := bytes.Buffer{}
	numberOfBytes := keySize/8 - 11

	start := 0
	end := numberOfBytes

	for end <= len(src) {
		blockBytes := src[start:end]
		encryptedBlock, err := c.EncryptBlock(ctx, publicKey, string(blockBytes))
		if err != nil {
			return "", fmt.Errorf("%v: %w", genericConstants.EncryptionError, err)
		}
		buffer.WriteString(encryptedBlock + "\n")
		start = end
		end += numberOfBytes
	}

	if start < len(src) {
		blockBytes := src[start:]
		encryptedBlock, err := c.EncryptBlock(ctx, publicKey, string(blockBytes))
		if err != nil {
			return "", fmt.Errorf(genericConstants.EncryptionProcessError, err)
		}
		buffer.WriteString(encryptedBlock + "\n")
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

func (c *CryptoRSA) EncryptPassword(ctx context.Context, spanCtx context.Context, plainText string) (string, error) {
	_, span := tracer.AddToSpan(ctx, "EncryptPassword")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf(genericConstants.UserPasswordHashingError, err)
	}
	return string(hashedPassword), nil
}

func (c *CryptoRSA) ComparePasswords(ctx context.Context, spanCtx context.Context, hashedPassword string, newPassword string) error {
	_, span := tracer.AddToSpan(ctx, "ComparePasswords")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(newPassword)); err != nil {
		return fmt.Errorf(genericConstants.ComparePasswordError, err)
	}
	return nil
}

// Encryption encrypts a string, slice, or struct with tagged fields.
func Encryption(ctx context.Context, input interface{}) (interface{}, error) {
	ctx, span := tracer.AddToSpan(ctx, "Encryption")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	val := reflect.ValueOf(input)

	// Retrieve Postgres encryption key
	PostgresConfig, err := configs.Get(genericConstants.PostgresConfig)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.GetPostgresConfigError, err)
	}
	key := []byte(PostgresConfig.GetString(genericConstants.PostgresSecretKey))
	key = utils.PadKey(key)

	switch val.Kind() {
	case reflect.String:
		return encryptString(ctx, val.String(), key)

	case reflect.Slice, reflect.Array:
		return encryptSlice(ctx, val)

	case reflect.Struct:
		return encryptStruct(ctx, val, key)

	default:
		return nil, fmt.Errorf(genericConstants.EncryptDecryptUnsupportedTypeError, val.Kind())
	}
}

// encryptSlice handles encryption of slices or arrays
func encryptSlice(ctx context.Context, val reflect.Value) (interface{}, error) {
	_, span := tracer.AddToSpan(ctx, "encryptSlice")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	encryptedSlice := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		encryptedElem, err := Encryption(ctx, elem.Interface())
		if err != nil {
			return nil, fmt.Errorf(genericConstants.EncryptionProcessError, err)
		}
		encryptedSlice.Index(i).Set(reflect.ValueOf(encryptedElem))
	}
	return encryptedSlice.Interface(), nil
}

// encryptStruct handles encryption of structs with tagged fields
func encryptStruct(ctx context.Context, val reflect.Value, key []byte) (interface{}, error) {
	_, span := tracer.AddToSpan(ctx, "encryptStruct")
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
				encryptedValue, err := encryptString(ctx, field.String(), key)
				if err != nil {
					return nil, fmt.Errorf(genericConstants.EncryptionProcessError, err)
				}
				output.Field(i).SetString(encryptedValue)
			case reflect.Slice, reflect.Array:
				encryptedSlice, err := encryptSlice(ctx, field)
				if err != nil {
					return nil, err
				}
				output.Field(i).Set(reflect.ValueOf(encryptedSlice))
			default:
				output.Field(i).Set(field)
			}
		} else {
			output.Field(i).Set(field)
		}
	}
	return output.Interface(), nil
}

// encryptString encrypts a plain string using AES-GCM.
func encryptString(ctx context.Context, plainText string, keyBytes []byte) (string, error) {
	_, span := tracer.AddToSpan(ctx, "encryptString")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	plainTextBytes := []byte(plainText)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf(genericConstants.FailedToCreateCipherError, err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf(genericConstants.FailedToCreateGCMError, err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf(genericConstants.FailedToGenerateNonceError, err)
	}

	cipherText := gcm.Seal(nonce, nonce, plainTextBytes, nil)
	return base64.URLEncoding.EncodeToString(cipherText), nil
}
