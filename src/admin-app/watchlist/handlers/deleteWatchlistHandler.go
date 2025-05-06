package handlers

import (
	"encoding/json"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/responseUtils"
	"strings"

	"admin-app/watchlist/business"
	"admin-app/watchlist/models"
	genericConstants "omnenest-backend/src/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// DeleteWatchlistController handles deleting watchlists
type DeleteWatchlistController struct {
	service *business.DeleteWatchlistService
}

// NewDeleteWatchlistController creates a new instance of the controller
func NewDeleteWatchlistController(service *business.DeleteWatchlistService) *DeleteWatchlistController {
	return &DeleteWatchlistController{
		service: service,
	}
}

// HandleDeleteWatchlist deletes created watchlist
// @Summary Delete Watchlist API
// @Description Delete Watchlist API for deleting the watchlist created by the user.
// @Tags Delete Watchlist
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param xRequestId header string true "Unique request identifier" default(123456789)
// @Param deviceId header string true "Unique device identifier" default(PKQ1.180904.001)
// @Param appVersion header string true "Current app version" default(1.0.0)
// @Param source header string true "Source (MOB or WEB)" default(MOB)
// @Param bypass header string false "Bypass (AUTOMATION or CHART)"
// @Param appInstallId header string true "Unique appInstall identifier" default(ba6eb330-4f7f-11eb-a2fb-67c34e9ac07c)
// @Param userAgent header  string true "userAgent" default(BrokerAppName/3.3.6 (OnePlus ONEPLUS A6010; Android 11 SDK30))
// @Param timestamp header string true "device current day epoch milliseconds timestamp" default(1700839140000)
// @Param request body models.BFFDeleteWatchlistRequest true "DeleteWatchlistRequest JSON"
// @Success 200 {object} models.BFFDeleteWatchlistResponse
// @Failure 404 {object} models.ErrorAPIResponse "Not Found: User not found"
// @Failure 500 {object} models.ErrorAPIResponse "Internal Server Error"
// @Router /api/watchlist/delete [delete]
func (controller *DeleteWatchlistController) HandleDeleteWatchlist(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleDeleteWatchlist")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	log := logger.GetLogger(ctx)
	var bffDeleteWatchlistRequest models.BFFDeleteWatchlistRequest

	if err := ctx.ShouldBindJSON(&bffDeleteWatchlistRequest); err != nil {
		errorMsgs := genericModel.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: genericConstants.JsonBindingFieldError}
		log.With(zap.Error(err)).Error(err.Error())
		responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsgs})
		return
	}

	ctx.Set(genericConstants.RequestBody, bffDeleteWatchlistRequest)

	if err := validations.GetBFFValidator(spanCtx).Struct(bffDeleteWatchlistRequest); err != nil {
		validationErrors, validationErrorsStr := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		log.With(zap.Error(err)).Error(validationErrorsStr)
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}
	err := controller.service.DeleteWatchlist(ctx, spanCtx, bffDeleteWatchlistRequest)
	if err != nil {
		log.With(zap.Error(err)).Error(err.Error())
		if strings.Contains(strings.ToLower(err.Error()), genericConstants.NoDataFoundError) {
			responseUtils.SendNotFoundJSON(ctx, err)
			return
		}
		responseUtils.SendInternalServerError(ctx, err)
		return
	}

	bffDeleteWatchlistsResponse := models.BFFDeleteWatchlistResponse{
		Message: genericConstants.BFFResponseSuccessMessage,
	}
	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, bffDeleteWatchlistsResponse)
}
