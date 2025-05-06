package authorization

import (
	"errors"
	"fmt"
	"net/http"
	genericConstants "omnenest-backend/src/constants"
	dbDeviceRepository "omnenest-backend/src/database/repository"
	"omnenest-backend/src/models"
	jwtUtils "omnenest-backend/src/utils/authorization"
	// "omnenest-backend/src/utils/cockroachDB"
	"omnenest-backend/src/utils/configs"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"

	"time"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AuthorizeJwtToken is a Go function that handles authorization of a JWT access token.
//
// It takes a gin.Context as a parameter.
// It does not return anything.
func AuthorizeJWtToken(jwtUtils jwtUtils.JwtTokenUtils, appType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, span := tracer.AddToSpan(ctx.Request.Context(), "middleware-AuthorizeJWtToken")
		defer func() {
			if span != nil {
				span.End()
			}
		}()
		log := logger.GetLogger(ctx)
		tokenString := ctx.GetHeader(genericConstants.Authorization)
		if tokenString == "" {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.JWTTokenMissingError))
			return
		}

		if tokenString[:len(genericConstants.Bearer)] != genericConstants.Bearer {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.JWTTokenBearerMissingError))
			return
		}
		tokenString = tokenString[len(genericConstants.Bearer):]

		applicationConfig, appConfigErr := configs.Get(genericConstants.ApplicationConfig)
		if appConfigErr != nil {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.GetApplicationConfigError, appConfigErr))
			return
		}

		secretKey := applicationConfig.GetString(genericConstants.AccessTokenSecretKey)

		claims, err := jwtUtils.ParseJwtToken(tokenString, secretKey)
		if err != nil {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusUnauthorized, err)
			return
		}

		if claims[genericConstants.TokenType] != genericConstants.AccessToken {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusUnauthorized, fmt.Errorf(genericConstants.JWTAccessTokenInvalidError))
			return
		}

		expirationTime := claims[genericConstants.TokenExpiration]
		currentTime := time.Now().Unix()
		if int64(expirationTime.(float64)) < currentTime {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusForbidden, fmt.Errorf(genericConstants.JWTTokenExpiredError))
			return
		}

		tokenPayload, ok := claims[genericConstants.TokenPayloadClaims].(string)
		if !ok {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusUnauthorized, fmt.Errorf(genericConstants.ExtractPayloadError))
			return
		}

		decodeTokenData, err := jwtUtils.DecryptTokenData(tokenPayload, genericConstants.EncryptionKeyMiddlewareMapping[appType])
		if err != nil {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusUnauthorized, err)
			return
		}

		deviceId := ctx.GetHeader(genericConstants.DeviceIdHeader)
		if deviceId == "" {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.DeviceIdKeyMissingError))
			return
		}

		deviceRepository := dbDeviceRepository.NewDeviceRepository()
		var deviceInfo *models.Devices
		var exists bool

		if appType == genericConstants.AppEncryptionKey {
			deviceInfo, exists, err = deviceRepository.GetDeviceByDeviceID(postgres.GetPostGresClient().GormDb, deviceId, genericConstants.Username)
		} else {
			deviceInfo, exists, err = deviceRepository.GetDeviceByDeviceID(postgres.GetPostGresClient().GormDb, deviceId, genericConstants.Username)
		}

		if err != nil {
			log.Error(genericConstants.ErrorLogParam, zap.Error(err))
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusInternalServerError, fmt.Errorf(genericConstants.InternalServerError))
			return
		}

		if !exists || deviceInfo == nil || deviceInfo.Username != decodeTokenData.Username {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusPreconditionFailed, fmt.Errorf(genericConstants.DeviceRetrieveError, errors.New(deviceId)))
			return
		}

		if decodeTokenData != nil {
			ctx.Set(genericConstants.UserID, decodeTokenData.UserId)
			ctx.Set(genericConstants.Username, decodeTokenData.Username)
			ctx.Set(genericConstants.ServerKeyPair, decodeTokenData.ServerPublicKey)
			ctx.Set(genericConstants.UserSessionToken, decodeTokenData.UserSessionId)
			ctx.Set(genericConstants.BFFCtxPublicKey, decodeTokenData.BFFPublicKey)
			ctx.Set(genericConstants.BFFCtxPrivateKey, decodeTokenData.BFFPrivateKey)
			ctx.Set(genericConstants.DeviceCtxPublicKey, decodeTokenData.DevicePublicKey)
			ctx.Set(genericConstants.AccountID, decodeTokenData.AccountId)
			ctx.Set(genericConstants.BrokerName, decodeTokenData.BrokerName)
			ctx.Set(genericConstants.BranchName, decodeTokenData.BranchName)
			ctx.Set(genericConstants.CriteriaAttributeKey, decodeTokenData.CriteriaAttribute)
			ctx.Set(genericConstants.ProductAlias, decodeTokenData.ProductAlias)
			ctx.Set(genericConstants.ClearingOrg, decodeTokenData.ClearingOrg)
			ctx.Set(genericConstants.EnabledExchangesKey, decodeTokenData.EnabledExchanges)
			ctx.Set(genericConstants.GttEnabledKey, decodeTokenData.GttEnabled)
		}
	}
}

// AuthorizeRefreshJWtToken is a Go function that handles the authorization of a refresh JWT token.
//
// This function takes a *gin.Context parameter.
// It does not return any values.
func AuthorizeRefreshJWtToken(jwtUtils jwtUtils.JwtTokenUtils) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, span := tracer.GetTracer().Start(ctx.Request.Context(), "middleware-AuthorizeRefreshJWtToken")
		defer func() {
			if span != nil {
				span.End()
			}
		}()
		tokenString := ctx.GetHeader(genericConstants.Authorization)
		if tokenString == "" {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.JWTTokenMissingError))
			return
		}

		applicationConfig, appConfigErr := configs.Get(genericConstants.ApplicationConfig)
		if appConfigErr != nil {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusBadRequest, fmt.Errorf(genericConstants.GetApplicationConfigError, appConfigErr))
			return
		}

		secretKey := applicationConfig.GetString(genericConstants.RefreshTokenSecretKey)

		claims, err := jwtUtils.ParseJwtToken(tokenString, secretKey)
		if err != nil {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusUnauthorized, err)
			return
		}

		if claims[genericConstants.TokenType] != genericConstants.RefreshToken {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusUnauthorized, fmt.Errorf(genericConstants.JWTRefreshTokenInvalidError))
			return
		}

		expirationTime := claims[genericConstants.TokenExpiration]
		currentTime := time.Now().Unix()
		if int64(expirationTime.(float64)) < currentTime {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusForbidden, fmt.Errorf(genericConstants.JWTRefreshTokenExpiredError))
			return
		}

		tokenPayload, ok := claims[genericConstants.TokenPayloadClaims].(string)
		if !ok {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusUnauthorized, fmt.Errorf(genericConstants.ExtractPayloadError))
			return
		}

		var tokenData models.TokenData
		if err := sonic.Unmarshal([]byte(tokenPayload), &tokenData); err != nil {
			responseUtils.SendAbortWithStatusJSON(ctx, http.StatusUnauthorized, fmt.Errorf(genericConstants.UnmarshalTokenDataError, err))
			return
		}

		if tokenData.Username != "" {
			ctx.Set(genericConstants.Username, tokenData.Username)
		}
	}
}
