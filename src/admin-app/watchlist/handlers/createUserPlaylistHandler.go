package handlers

import (
	"encoding/json"

	"admin-app/watchlist/business"
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
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

type PlaylistController struct {
	service *business.CreateUserPlaylistService
}

func NewPlaylistController(service *business.CreateUserPlaylistService) *PlaylistController {
	return &PlaylistController{
		service: service,
	}
}

func (controller *PlaylistController) HandleCreatePlaylist(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleCreatePlaylist")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	log := logger.GetLogger(ctx)

	var createRequest models.BFFPlaylistRequest
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		var errorMsg genericModel.ErrorMessage
		if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
			errorMsg = genericModel.ErrorMessage{
				Key:          unmarshalErr.Field,
				ErrorMessage: genericConstants.JsonBindingFailedError,
			}
		} else {
			errorMsg = genericModel.ErrorMessage{
				Key:          constants.Name,
				ErrorMessage: genericConstants.JsonBindingFailedError,
			}
		}
		log.With(zap.Error(err)).Error(err.Error())
		responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsg})
		return
	}

	ctx.Set(genericConstants.RequestBody, createRequest)

	if err := validations.GetBFFValidator(spanCtx).Struct(createRequest); err != nil {
		validationErrors, validationErrorsStr := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		log.With(zap.Error(err)).Error(validationErrorsStr)
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}

	response, err := controller.service.CreateUserPlaylist(ctx, spanCtx, createRequest)
	if err != nil {
		log.With(zap.Error(err)).Error(err.Error())
		if err.Error() == constants.DuplicatePlaylistError {
			responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{
				{Key: constants.Name, ErrorMessage: constants.DuplicatePlaylistError},
			})
			return
		}
		if err.Error() == constants.InvalidSongIDsError {
			responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{
				{Key: constants.Song_ids, ErrorMessage: constants.InvalidSongIDsError},
			})
			return
		}
		if err.Error() == constants.NoUserIdFoundError {
			responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{
				{Key: constants.User_id, ErrorMessage: constants.InvalidUserID},
			})
			return
		}
		responseUtils.SendInternalServerError(ctx, err)
		return
	}

	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, response)
}
