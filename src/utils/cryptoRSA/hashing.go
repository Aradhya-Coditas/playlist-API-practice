package cryptoRSA

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/tracer"
	"os"
)

// HashData takes a string of data as input and returns its SHA-256 hash as a hexadecimal string.
// It uses the SHA-256 hashing algorithm provided by the crypto package in Go.
func HashData(ctx context.Context, data string) (string, error) {
	_, span := tracer.AddToSpan(ctx, "HashData")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	hash := sha256.New()
	_, err := hash.Write([]byte(data))
	if err != nil {
		return "", err
	}

	digest := hash.Sum(nil)
	hashedData := hex.EncodeToString(digest)
	return hashedData, nil
}

// IterativeStringHashing is a function that takes a string as input and performs iterative string hashing using SHA-256 algorithm.
// It returns the hashed password as a hexadecimal string.
// It uses the SHA-256 hashing algorithm provided by the crypto package in Go.
func IterativeStringHashing(ctx context.Context, data string) (string, error) {
	_, span := tracer.AddToSpan(ctx, "IterativeStringHashing")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	log := logger.GetLoggerWithoutContext()

	hash := sha256.New()
	_, err := hash.Write([]byte(data))
	if err != nil {
		return "", fmt.Errorf(genericConstants.HashPasswordError, err)
	}

	digest := hash.Sum(nil)

	for i := 1; i <= 999; i++ {
		hash.Reset()
		hash.Write(digest)
		digest = hash.Sum(nil)
	}

	hashedPassword := hex.EncodeToString(digest)
	log.Info(genericConstants.ObtainedHashedPassword)
	return hashedPassword, nil
}

// SaveHashToFile saves the provided data to a file at the specified file path.
func SaveHashToFile(ctx context.Context, filePath string, data string) error {
	_, span := tracer.AddToSpan(ctx, "SaveHashToFile")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

// ReadHashFromFile reads the hash data from a file located at the specified file path.
func ReadHashFromFile(ctx context.Context, filePath string) (string, error) {
	_, span := tracer.AddToSpan(ctx, "ReadHashFromFile")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
