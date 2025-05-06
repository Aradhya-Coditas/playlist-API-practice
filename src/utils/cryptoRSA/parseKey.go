package cryptoRSA

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/tracer"
	"os"

	"github.com/pkg/errors"
)

// ParsePrivateKey parses a PEM-encoded RSA private key from a byte slice and returns it as a parsed RSA private key.
func ParsePrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New(genericConstants.PrivateKeyError)
	}

	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.ParsePrivateKeyError, err)
	}

	rsaPrivateKey, ok := private.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New(genericConstants.NotRSAPrivateKeyError)
	}

	return rsaPrivateKey, nil
}

// ParseRSAPrivateKeyPKCS8FromFile reads a PEM-encoded RSA private key from a file specified by filePath and returns it as a parsed RSA private key.
func ParseRSAPrivateKeyPKCS8FromFile(filePath string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.ReadPrivateKeyFileError, err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New(genericConstants.DecodePEMError)
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.ParsePrivateKeyError, err)
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New(genericConstants.ParseRSAPrivateKeyError)
	}
	return rsaKey, nil
}

// LoadPublicKeyFromFile loads a public key from a PEM file and optionally a hash from a separate file, and returns the public key and the hashed public key.
func LoadPublicKeyFromFile(ctx context.Context, publicKeyPath string, publicKeyHashPath string) (*genericModels.NestKeyPair, error) {
	ctx, span := tracer.AddToSpan(ctx, "LoadPublicKeyFromFile")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	var nestServerKey genericModels.NestKeyPair

	publicKey, err := ReadPublicKeyFromPEMFile(ctx, publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.ReadPublicKeyError, err)
	}

	nestServerKey.PublicKey = publicKey
	if publicKeyHashPath != "" {
		publicKeyHash, err := ReadHashFromFile(ctx, publicKeyHashPath)
		if err != nil {
			return nil, fmt.Errorf(genericConstants.ReadPublicKeyHashError, err)
		}
		nestServerKey.PublicHashedKey = publicKeyHash
	}
	return &nestServerKey, nil
}

// ParsePublicKey parses a public key string and returns the public key and hashed public key.
func ParsePublicKey(ctx context.Context, publicKey string) (*genericModels.NestKeyPair, error) {
	ctx, span := tracer.AddToSpan(ctx, "ParsePublicKey")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	publicKeyBytes, decodeStringErr := base64.StdEncoding.DecodeString(publicKey)
	if decodeStringErr != nil {
		return nil, fmt.Errorf(genericConstants.DecodePublicKeyError, decodeStringErr)
	}

	hashedPublicKey, hashDataErr := HashData(ctx, string(publicKeyBytes))
	if hashDataErr != nil {
		return nil, fmt.Errorf(genericConstants.HashPublicKeyError, hashDataErr)
	}

	publicKeyRSA, errFromPEMData := GetPublicKeyFromPEMData(ctx, string(publicKeyBytes))
	if errFromPEMData != nil {
		return nil, fmt.Errorf(genericConstants.ParsePublicKeyRsaPemKeyError, errFromPEMData)
	}

	key := &models.NestKeyPair{
		PublicKey:       publicKeyRSA,
		PublicHashedKey: hashedPublicKey,
	}

	return key, nil
}
