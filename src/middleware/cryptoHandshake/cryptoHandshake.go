package cryptoHandshake

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	genericConstants "omnenest-backend/src/constants"
	repositories "omnenest-backend/src/database/repository"
	"omnenest-backend/src/models"
	"omnenest-backend/src/utils"
	"omnenest-backend/src/utils/cryptoRSA"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"
	"strings"

	"github.com/bytedance/sonic"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CustomResponseWriter is a custom response writer that captures the response body.
type CustomResponseWriter struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

// NewCustomResponseWriter creates a new CustomResponseWriter.
func NewCustomResponseWriter(rw gin.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{rw, &bytes.Buffer{}, 0}
}

// Write captures the response body.
func (rw *CustomResponseWriter) Write(b []byte) (int, error) {
	return rw.body.Write(b)
}

// WriteHeader captures the response status code.
func (rw *CustomResponseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// handleEncryptionError handles encryption errors and sends an error response.
func (rw *CustomResponseWriter) handleEncryptionError(err error) {
	rw.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	responseBytes, _ := sonic.Marshal(models.ErrorAPIResponse{
		Message: []models.ErrorMessage{{
			Key:          genericConstants.GenericErrorKey,
			ErrorMessage: err.Error(),
		}},
		Error: http.StatusText(http.StatusInternalServerError),
	})
	rw.ResponseWriter.Write(responseBytes)
}

func EncryptMiddleware(deviceRepo repositories.DeviceRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appEnv := utils.GetEnv(genericConstants.APPEnvironment, genericConstants.ProdEnvironment)
		bypass := ctx.Value(genericConstants.Bypass).(string)

		if (bypass == genericConstants.AutomationFlag && appEnv != genericConstants.ProdEnvironment) || (bypass == genericConstants.ChartSource && ctx.Request.URL.Path == genericConstants.StocksIntradayAggrDataEndpoint) {
			ctx.Next()
			return
		}

		spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "EncryptMiddleware")
		defer func() {
			if span != nil {
				span.End()
			}
		}()
		deviceID := ctx.Request.Header.Get(genericConstants.DeviceIdHeader)

		// Create a capture writer to intercept the response and get the DeviceId
		captureWriter := NewCustomResponseWriter(ctx.Writer)
		ctx.Writer = captureWriter
		ctx.Next()

		// Get the captured response body and status code
		responseBody := captureWriter.body.String()
		responseStatus := captureWriter.status

		if responseStatus == http.StatusOK || responseStatus == http.StatusCreated {
			encryptedResponse, err := encryptResponse(ctx, spanCtx, deviceRepo, []byte(responseBody), deviceID)
			if err != nil {
				captureWriter.handleEncryptionError(err)
				return
			}
			response := models.EncryptResponse{EncResponse: string(encryptedResponse)}
			responseByte, err := sonic.Marshal(response)
			if err != nil {
				captureWriter.handleEncryptionError(fmt.Errorf(genericConstants.MarshalResponseError, err))
				return
			}
			captureWriter.body = &bytes.Buffer{}
			captureWriter.Write(responseByte)
			captureWriter.ResponseWriter.WriteHeader(responseStatus)
			captureWriter.ResponseWriter.Write(captureWriter.body.Bytes())
			return
		} else {
			captureWriter.ResponseWriter.WriteHeader(responseStatus)
			captureWriter.ResponseWriter.Write([]byte(responseBody))
			return
		}
	}
}

// encryptResponse encrypts the response data using the device's public key.
//
// Parameters:
// - ctx: the Gin context.
// - deviceRepo: the device repository.
// - data: the data to be encrypted.
// - deviceID: the ID of the device.
//
// Returns:
// - string: the encrypted data.
// - error: an error if encryption fails.
func encryptResponse(ctx *gin.Context, spanCtx context.Context, deviceRepo repositories.DeviceRepository, data []byte, deviceID string) (string, error) {
	cryptoRsa := cryptoRSA.NewCryptoRSA()
	var stringDevicePublicKey string

	jwtToken := ctx.Request.Header.Get(genericConstants.Authorization)
	if jwtToken != "" {
		stringDevicePublicKey = ctx.Value(genericConstants.DeviceCtxPublicKey).(string)
	} else {
		client := postgres.GetPostGresClient()

		// Retrieve the device's public from the database
		device, exists, err := deviceRepo.GetDeviceByDeviceID(client.GormDb, deviceID, genericConstants.DevicePublicKey)
		if err != nil {
			return "", fmt.Errorf(genericConstants.DeviceRetrieveError, err)
		}

		if !exists {
			return "", fmt.Errorf(genericConstants.DeviceNotFoundError)
		}

		stringDevicePublicKey = device.DevicePublicKey
	}

	devicePubKey, err := cryptoRSA.GetPublicKeyFromPEMData(spanCtx, stringDevicePublicKey)
	if err != nil {
		return "", fmt.Errorf(genericConstants.ParseDevicePublicKeyError, err)
	}
	encrypted, err := cryptoRsa.Encrypt(spanCtx, string(data), devicePubKey, genericConstants.KeySize)
	if err != nil {
		return "", fmt.Errorf(genericConstants.EncryptResponseBodyError, err)
	}
	return encrypted, nil
}

// DecryptionMiddleware is a middleware function that decrypts the request body using RSA and AES encryption.
//
// It takes a deviceRepo parameter of type repositories.DeviceRepository, which is used to retrieve the device information.
// The function returns a gin.HandlerFunc.
func DecryptionMiddleware(deviceRepo repositories.DeviceRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appEnv := utils.GetEnv(genericConstants.APPEnvironment, genericConstants.ProdEnvironment)
		bypass := ctx.Value(genericConstants.Bypass).(string)
		if (bypass == genericConstants.AutomationFlag && appEnv != genericConstants.ProdEnvironment) || (bypass == genericConstants.ChartSource && ctx.Request.URL.Path == genericConstants.StocksIntradayAggrDataEndpoint) {
			ctx.Next()
			return
		}

		spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "DecryptionMiddleware")
		defer func() {
			if span != nil {
				span.End()
			}
		}()

		cryptoRsa := cryptoRSA.NewCryptoRSA()
		var stringBFFPrivateKey string
		jwtToken := ctx.Request.Header.Get(genericConstants.Authorization)
		if jwtToken != "" {
			stringBFFPrivateKey = ctx.Value(genericConstants.BFFCtxPrivateKey).(string)
		} else {
			client := postgres.GetPostGresClient()
			deviceId := ctx.Request.Header.Get(genericConstants.DeviceIdHeader)

			device, exists, err := deviceRepo.GetDeviceByDeviceID(client.GormDb, deviceId, genericConstants.BFFPrivateKey)
			if err != nil {
				responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.DeviceRetrieveError, err))
				return
			}

			if !exists {
				responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.DeviceNotFoundError))
				return
			}
			stringBFFPrivateKey = device.BFFPrivateKey
		}
		var encRequest models.EncryptRequest
		if err := ctx.ShouldBindJSON(&encRequest); err != nil {
			if err == io.EOF {
				ctx.Next()
				return
			}
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.RequestBodyReadError, err))
			return
		}

		if err := validations.GetBFFValidator(spanCtx).Struct(encRequest); err != nil {
			_, validationErrorsStr := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(validationErrorsStr))
			return
		}

		bffPrivateKey, err := cryptoRSA.GetPrivateKeyFromPKCS8PEMData(spanCtx, stringBFFPrivateKey)
		if err != nil {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.ParseBFFPrivateKeyError, err))
			return
		}

		// Decrypt the request body using AES encryption
		decryptedBody, err := cryptoRsa.Decrypt(spanCtx, encRequest.EncRequest, bffPrivateKey)
		if err != nil {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.DecryptRequestBodyError, err))
			return
		}

		// Create an io.ReadCloser from the decrypted content
		decryptedBodyReader := io.NopCloser(strings.NewReader(decryptedBody))

		// Replace the request body with the decrypted content
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, decryptedBodyReader, int64(len(decryptedBody)))
		ctx.Next()
	}
}
