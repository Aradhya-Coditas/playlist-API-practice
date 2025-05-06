package handlers

import (
	"admin-app/watchlist/business"
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"encoding/json"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"

	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type GetWatchlistScripsController struct {
	service *business.GetWatchlistScripsService
}

// GetWatchListController initializes a new controller with the given service.
func NewGetWatchlistScripsController(service *business.GetWatchlistScripsService) *GetWatchlistScripsController {
	return &GetWatchlistScripsController{
		service: service,
	}
}

// HandleGetWatchListScrips handles the request to fetch watchlist scrips.
func (controller *GetWatchlistScripsController) HandleGetWatchlistScrips(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleGetWatchListScrips")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	log := logger.GetLogger(ctx)
	var bffGetWatchlistScripsRequest models.BFFGetWatchlistScripsRequest

	if err := ctx.ShouldBindJSON(&bffGetWatchlistScripsRequest); err != nil {
		var errorMsg genericModel.ErrorMessage

		// Check if error is a JSON unmarshalling type error
		if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
			errorMsg = genericModel.ErrorMessage{
				Key:          unmarshalErr.Field,
				ErrorMessage: genericConstants.JsonBindingFailedError,
			}
		} else {
			// Generic JSON binding error (e.g., missing field or syntax error)
			errorMsg = genericModel.ErrorMessage{
				Key:          constants.WatchlistId,
				ErrorMessage: genericConstants.JsonBindingFailedError,
			}
		}

		log.With(zap.Error(err)).Error(err.Error())
		responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsg})
		return
	}

	ctx.Set(genericConstants.RequestBody, bffGetWatchlistScripsRequest)

	if err := validations.GetBFFValidator(spanCtx).Struct(bffGetWatchlistScripsRequest); err != nil {
		validationErrors, validationErrorsStr := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		log.With(zap.Error(err)).Error(validationErrorsStr)
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}

	scrips, err := controller.service.GetWatchlistScrips(ctx, spanCtx, bffGetWatchlistScripsRequest)
	if err != nil {
		log.With(zap.Error(err)).Error(err.Error())

		invalidWatchlistErrMsg := fmt.Sprintf(genericConstants.InvalidWatchlistIdError, bffGetWatchlistScripsRequest.WatchlistId)
		if strings.Contains(err.Error(), invalidWatchlistErrMsg) {
			responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{
				{Key: constants.WatchlistId, ErrorMessage: fmt.Sprintf(genericConstants.InvalidWatchlistIdError, bffGetWatchlistScripsRequest.WatchlistId)},
			})
			return
		}

		if strings.Contains(err.Error(), genericConstants.NoWatchlistScripsFoundError) {
			responseUtils.SendNoContentFoundError(ctx, err)
			return
		}

		if strings.Contains(err.Error(), genericConstants.NoDataFoundError) {
			responseUtils.SendNotFoundJSON(ctx, err)
			return
		}

		responseUtils.SendInternalServerError(ctx, err)
		return
	}

	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, scrips)
}
