package cryptoRSA

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/tracer"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

// CryptoRSA provides RSA encryption and decryption functions.
type CryptoRSA struct{}

// NewCryptoRSA creates a new instance of the CryptoRSA struct.
func NewCryptoRSA() *CryptoRSA {
	return &CryptoRSA{}
}

// MarshalPEMPublicKeyToString converts a PEM-encoded RSA public key to a string
func MarshalPEMPublicKeyToString(ctx context.Context, publicKey *rsa.PublicKey) (string, error) {
	_, span := tracer.AddToSpan(ctx, "MarshalPEMPublicKeyToString")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	// Encode the RSA public key to DER format
	derBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	// Create a PEM block for the public key
	pemBlock := &pem.Block{
		Type:  genericConstants.KeyTypePublic,
		Bytes: derBytes,
	}

	// Encode the PEM block to a string
	pemString := string(pem.EncodeToMemory(pemBlock))

	// Remove trailing newline character, if any
	pemString = strings.TrimRight(pemString, "\n")

	return pemString, nil
}

// MarshalPEMPrivateKeyToString converts an RSA private key to a PEM-encoded string.
func MarshalPEMPrivateKeyToString(ctx context.Context, privateKey *rsa.PrivateKey) (string, error) {
	_, span := tracer.AddToSpan(ctx, "MarshalPEMPrivateKeyToString")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	// Marshal the RSA private key to PKCS#8 DER format
	derBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", err
	}

	// Create a PEM block for the private key
	pemBlock := &pem.Block{
		Type:  genericConstants.KeyTypePrivate,
		Bytes: derBytes,
	}

	// Encode the PEM block to a string
	pemString := string(pem.EncodeToMemory(pemBlock))

	// Remove trailing newline character, if any
	pemString = strings.TrimRight(pemString, "\n")

	return pemString, nil
}

// GetPublicKeyFromPEMData extracts an RSA public key from a PEM-encoded string.
func GetPublicKeyFromPEMData(ctx context.Context, pemPublicKey string) (*rsa.PublicKey, error) {
	_, span := tracer.AddToSpan(ctx, "GetPublicKeyFromPEMData")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	block, _ := pem.Decode([]byte(pemPublicKey))
	if block == nil {
		return nil, errors.New(genericConstants.DecodePEMError)
	}

	if block.Type != genericConstants.KeyTypePublic {
		return nil, errors.New(genericConstants.UnexpectedPEMBlockTypeError)
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New(genericConstants.ParseRSAPublicKeyError)
	}

	publicKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New(genericConstants.ParseRSAPublicKeyError)
	}

	return publicKey, nil
}

// CreateRSAKeys generates RSA public and private key pairs and writes them to the provided writers.
func CreateRSAKeys(ctx context.Context, publicKeyWriter, privateKeyWriter io.Writer, keyLength int) error {
	_, span := tracer.AddToSpan(ctx, "CreateRSAKeys")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return fmt.Errorf(genericConstants.GeneratePrivateKeyError, err)
	}

	derStream, err := x509.MarshalPKCS8PrivateKey(privateKey)
	block := &pem.Block{
		Type:  genericConstants.KeyTypePrivate,
		Bytes: derStream,
	}

	if err != nil {
		return fmt.Errorf(genericConstants.MarshalPrivateKeyError, err)
	}

	err = pem.Encode(privateKeyWriter, block)
	if err != nil {
		return fmt.Errorf(genericConstants.EncodePrivateKeyError, err)
	}

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf(genericConstants.MarshalPublicKeyError, err)
	}

	block = &pem.Block{
		Type:  genericConstants.KeyTypePublic,
		Bytes: derPkix,
	}

	err = pem.Encode(publicKeyWriter, block)
	if err != nil {
		return fmt.Errorf(genericConstants.EncodePrivateKeyError, err)
	}

	return nil
}

// ReadPublicKeyFromPEMFile reads an RSA public key from a PEM-encoded file.
func ReadPublicKeyFromPEMFile(ctx context.Context, filePath string) (*rsa.PublicKey, error) {
	_, span := tracer.AddToSpan(ctx, "ReadPublicKeyFromPEMFile")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	pubKeyFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", genericConstants.PublicKeyFileOpeningError, err)
	}
	defer pubKeyFile.Close()

	pubKeyBytes, err := io.ReadAll(pubKeyFile)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", genericConstants.PublicKeyFileReadingError, err)
	}

	pubBlock, _ := pem.Decode(pubKeyBytes)
	if pubBlock == nil {
		return nil, fmt.Errorf(genericConstants.DecodePublicKeyError, err)
	}

	pubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", genericConstants.PublicKeyParsingError, err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New(genericConstants.RSAPublicKeyCastingError)
	}

	return rsaPubKey, nil
}

// SavePublicKeyToFile saves the RSA public key to a PEM-encoded file.
func SavePublicKeyToFile(ctx context.Context, publicKey *rsa.PublicKey, filePath string) error {
	_, span := tracer.AddToSpan(ctx, "SavePublicKeyToFile")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	log := logger.GetLoggerWithoutContext()
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return fmt.Errorf(genericConstants.MarshalPublicKeyError, err)
	}

	pubKeyPEM := &pem.Block{
		Type:  genericConstants.KeyTypePublic,
		Bytes: pubKeyBytes,
	}

	file, err := os.Create(filePath)

	if err != nil {
		return fmt.Errorf(genericConstants.CreatePublicKeyFileError, err)
	}

	defer func(file *os.File) {
		err := file.Close()

		if err != nil {
			log.With(zap.Error(err), zap.String(genericConstants.ErrorLogParam, err.Error()))
			return
		}

	}(file)

	if err := pem.Encode(file, pubKeyPEM); err != nil {
		return fmt.Errorf(genericConstants.EncodePublicKeyError, err)
	}

	return nil
}

// SaveRSAPrivateKeyPKCS8ToFile saves an RSA private key in PKCS#8 format to a PEM-encoded file.
func SaveRSAPrivateKeyPKCS8ToFile(ctx context.Context, privateKey *rsa.PrivateKey, filePath string) error {
	_, span := tracer.AddToSpan(ctx, "SavePublicKeyToFile")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	log := logger.GetLoggerWithoutContext()
	keyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)

	if err != nil {
		return fmt.Errorf(genericConstants.MarshalPrivateKeyFileError, err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf(genericConstants.MarshalPublicKeyFileError, err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.With(zap.Error(err), zap.String(genericConstants.ErrorLogParam, err.Error()))
			return
		}

	}(file)

	if err := pem.Encode(file, &pem.Block{Type: genericConstants.KeyTypePrivate, Bytes: keyBytes}); err != nil {
		return fmt.Errorf(genericConstants.EncodePrivateKeyFileError, err)
	}

	return nil
}

// GetPrivateKeyFromPKCS8PEMData extracts an RSA private key from a PEM-encoded string.
func GetPrivateKeyFromPKCS8PEMData(ctx context.Context, pemPrivateKey string) (*rsa.PrivateKey, error) {
	_, span := tracer.AddToSpan(ctx, "GetPrivateKeyFromPKCS8PEMData")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	block, _ := pem.Decode([]byte(pemPrivateKey))
	if block == nil {
		return nil, errors.New(genericConstants.DecodePEMError)
	}

	if block.Type != genericConstants.KeyTypePrivate {
		return nil, errors.New(genericConstants.UnexpectedPEMBlockTypeError)
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New(genericConstants.PrivateKeyParseError)
	}

	return rsaPrivateKey, nil
}

// CreateBFFKeyPair generates a pair of RSA keys - a private key and a public key.
func CreateBFFKeyPair(ctx context.Context, log logger.Logger) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	ctx, span := tracer.AddToSpan(ctx, "CreateBFFKeyPair")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	bffPublicKey := bytes.NewBuffer([]byte{})
	bffPrivateKey := bytes.NewBuffer([]byte{})
	err := CreateRSAKeys(ctx, bffPublicKey, bffPrivateKey, genericConstants.KeySize)
	if err != nil {
		log.With(zap.Error(err)).Error(err.Error())
		return nil, nil, err
	}

	bffPublicKeyRSA, errFromPEMData := GetPublicKeyFromPEMData(ctx, bffPublicKey.String())
	if errFromPEMData != nil {
		log.With(zap.Error(errFromPEMData)).Error(errFromPEMData.Error())
		return nil, nil, errFromPEMData
	}

	bffPrivateKeyRSAParsed, parsePrivateKeyErr := ParsePrivateKey(bffPrivateKey.Bytes())
	if parsePrivateKeyErr != nil {
		log.With(zap.Error(parsePrivateKeyErr)).Error(parsePrivateKeyErr.Error())
		return nil, nil, parsePrivateKeyErr
	}

	return bffPrivateKeyRSAParsed, bffPublicKeyRSA, nil
}
