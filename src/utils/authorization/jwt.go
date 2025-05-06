package authorization

import (
	"context"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/models"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/cryptoRSA"
	"time"

	"github.com/bytedance/sonic"
	"github.com/golang-jwt/jwt/v5"
	"github.com/klauspost/compress/zstd"
)

type JwtTokenUtils interface {
	ParseJwtToken(tokenString string, secretKey string) (jwt.MapClaims, error)
	CreateJwtToken(tokenData string, expiryDays int, tokenType string, secretKey string) (string, error)
	DecryptTokenData(encryptedTokenData string, privateKeyPath string) (*models.TokenData, error)
	EncryptTokenData(tokenDataString string, publicKeyPath string) (string, error)
	CompressTokenData(tokenData string) (string, error)
	DecompressTokenData(tokenData []byte) (string, error)
}

type jwtTokenUtils struct{}

func NewJwtTokenUtils() *jwtTokenUtils {
	return &jwtTokenUtils{}
}

func (jwtUtil *jwtTokenUtils) ParseJwtToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if tokenErr != nil || !token.Valid {
		return nil, fmt.Errorf(genericConstants.JWTInvalidTokenError, tokenErr)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf(genericConstants.ExtractClaimsError)
	}

	return claims, nil
}

func (jwtUtil *jwtTokenUtils) CreateJwtToken(tokenData string, expiryDays int, tokenType string, secretKey string) (string, error) {
	claims := jwt.MapClaims{}
	claims[genericConstants.TokenPayload] = tokenData
	claims[genericConstants.TokenExpiration] = time.Now().Add(time.Duration(expiryDays) * 24 * time.Hour).Unix()
	claims[genericConstants.TokenType] = tokenType

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (jwtUtil *jwtTokenUtils) DecryptTokenData(encryptedTokenData string, privateKeyPath string) (*models.TokenData, error) {
	applicationConfig := configs.GetApplicationConfig()

	clientPrivateKey, err := cryptoRSA.ParseRSAPrivateKeyPKCS8FromFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.ParseClientPrivateKeyError, err)
	}

	decryptedTokenData, err := cryptoRSA.NewCryptoRSA().Decrypt(context.Background(), encryptedTokenData, clientPrivateKey)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.DecryptTokenDataError, err)
	}

	if applicationConfig.Token.EnableTokenCompression {
		decryptedTokenData, err = jwtUtil.DecompressTokenData([]byte(decryptedTokenData))
		if err != nil {
			return nil, err
		}
	}

	var tokenData models.TokenData
	if err := sonic.Unmarshal([]byte(decryptedTokenData), &tokenData); err != nil {
		return nil, fmt.Errorf(genericConstants.UnmarshalTokenDataError, err)
	}

	return &tokenData, nil
}

func (jwtUtil *jwtTokenUtils) EncryptTokenData(tokenDataString string, publicKeyPath string) (string, error) {
	applicationConfig := configs.GetApplicationConfig()

	clientPublicKey, err := cryptoRSA.ReadPublicKeyFromPEMFile(context.Background(), publicKeyPath)
	if err != nil {
		return "", fmt.Errorf(genericConstants.PublicKeyFileReadingError, err)
	}

	if applicationConfig.Token.EnableTokenCompression {
		tokenDataString, err = jwtUtil.CompressTokenData(tokenDataString)
		if err != nil {
			return "", err
		}
	}

	encryptedTokenData, err := cryptoRSA.NewCryptoRSA().Encrypt(context.Background(), tokenDataString, clientPublicKey, genericConstants.KeySize)
	if err != nil {
		return "", fmt.Errorf(genericConstants.JWTTokenDataEncryptionError, err)
	}

	return encryptedTokenData, nil
}

func (jwtUtil *jwtTokenUtils) CompressTokenData(tokenDataString string) (string, error) {
	writer, err := zstd.NewWriter(nil)
	if err != nil {
		return "", fmt.Errorf(genericConstants.JWTTokenCompressionError, err)
	}
	defer writer.Close()

	compressedData := writer.EncodeAll([]byte(tokenDataString), make([]byte, 0, len([]byte(tokenDataString))))
	if err != nil {
		return "", fmt.Errorf(genericConstants.JWTTokenCompressionError, err)
	}
	return string(compressedData), nil
}

func (jwtUtil *jwtTokenUtils) DecompressTokenData(compressedTokenData []byte) (string, error) {

	reader, err := zstd.NewReader(nil)
	if err != nil {
		return "", fmt.Errorf(genericConstants.JWTTokenDecompressionError, err)
	}
	defer reader.Close()

	decompressed, err := reader.DecodeAll(compressedTokenData, nil)
	if err != nil {
		return "", fmt.Errorf(genericConstants.JWTTokenDecompressionError, err)
	}
	return string(decompressed), nil
}
