package handlers

import (
	"admin-app/watchlist/business"
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"encoding/json"
	"strings"

	genericConstants "omnenest-backend/src/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type GetWatchlistController struct {
	service *business.GetWatchlistService
}

func NewGetWatchlistController(service *business.GetWatchlistService) *GetWatchlistController {
	return &GetWatchlistController{service: service}
}

// HandleGetWatchlist returns watchlists for the user and broker
// @Summary Get user and broker watchlists
// @Description Returns user's custom and broker's predefined watchlists
// @Tags Watchlist
// @Accept json
// @Produce json
// @Param request body models.BFFGetWatchlistRequest true "Request payload"
// @Success 200 {object} models.BFFWatchlistResponse
// @Failure 400 {object} models.ErrorAPIResponse
// @Failure 404 {object} models.ErrorAPIResponse
// @Failure 204 "No content found"
// @Failure 500 {object} models.ErrorAPIResponse
// @Router /api/watchlist/get [post]
func (controller *GetWatchlistController) HandleGetWatchlist(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleGetWatchlist")
	if span != nil {
		defer span.End()
	}

	log := logger.GetLogger(ctx)
	var req models.BFFGetWatchlistRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		var errorMsg genericModel.ErrorMessage
		if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
			errorMsg = genericModel.ErrorMessage{
				Key:          unmarshalErr.Field,
				ErrorMessage: genericConstants.JsonBindingFieldError,
			}
		} else {
			errorMsg = genericModel.ErrorMessage{
				Key:          constants.Request,
				ErrorMessage: genericConstants.JsonBindingFailedError,
			}
		}
		log.Error(constants.JSONBindingFailedError, zap.Error(err))
		responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsg})
		return
	}

	if err := validations.GetBFFValidator(spanCtx).Struct(req); err != nil {
		validationErrors, validationStr := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		log.Error(constants.ValidationFailedError, zap.String("errors", validationStr))
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}

	userWatchlists, brokerWatchlists, err := controller.service.GetWatchlist(ctx, spanCtx, req)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, constants.NoUserIdFoundError) || strings.Contains(errMsg, constants.NoBrokerIdFoundError) {
			responseUtils.SendNotFoundJSON(ctx, err)
			return
		}
		if strings.Contains(errMsg, constants.BothWatchlistError) {
			responseUtils.SendNoContentFoundError(ctx, nil)
			return
		}
		responseUtils.SendInternalServerError(ctx, err)
		return
	}

	response := models.BFFWatchlistResponse{
		Userdefine: userWatchlists,
		Predefine:  brokerWatchlists,
	}

	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, response)
}
