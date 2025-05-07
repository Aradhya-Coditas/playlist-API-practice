package handlers

import (
	"encoding/json"

	"admin-app/watchlist/business"
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"

	genericConstants "omnenest-backend/src/constants"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type PlaylistHandler struct {
	service *business.CreateSongPlaylistService
}

func NewPlaylistHandler(service *business.CreateSongPlaylistService) *PlaylistHandler {
	return &PlaylistHandler{
		service: service,
	}
}

// CreatePlaylist
// @Summary Create a new playlist
// @Description Creates a new playlist for a user with specified songs
// @Tags Playlists
// @Accept json
// @Produce json
// @Param request body models.BFFPlaylistRequest true "Playlist Creation Request"
// @Success 200 {object} map[string]interface{} "Successfully created playlist"
// @Failure 400 {object} models.ErrorMessage "Bad request - Invalid input or validation errors"
// @Failure 500 {object} models.ErrorMessage "Internal server error"
// @Router /api/Playlist/create [post]
func (h *PlaylistHandler) HandleCreatePlaylist(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), constants.HandleCreatePlaylistLog)
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	log := logger.GetLogger(ctx)

	var req models.BFFPlaylistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var errorMsg genericModels.ErrorMessage
		if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
			errorMsg = genericModels.ErrorMessage{
				Key:          unmarshalErr.Field,
				ErrorMessage: genericConstants.JsonBindingFailedError,
			}
		} else {
			errorMsg = genericModels.ErrorMessage{
				Key:          constants.Name,
				ErrorMessage: genericConstants.JsonBindingFailedError,
			}
		}
		log.With(zap.Error(err)).Error(constants.JSONBindingFailedError)
		responseUtils.SendBadRequest(ctx, []genericModels.ErrorMessage{errorMsg})
		return
	}

	ctx.Set(genericConstants.RequestBody, req)

	if err := validations.GetBFFValidator(spanCtx).Struct(req); err != nil {
		validationErrors, validationErrMsg := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		log.With(zap.Error(err)).Error(validationErrMsg)
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}

	playlistID, err := h.service.CreatePlaylist(ctx, spanCtx, req)
	if err != nil {
		log.With(zap.Error(err)).Error(constants.ServiceFailedToCreatePlaylist)
		switch err.Error() {
		case constants.DuplicatePlaylistError:
			responseUtils.SendBadRequest(ctx, []genericModels.ErrorMessage{
				{Key: constants.Name, ErrorMessage: constants.DuplicatePlaylistError},
			})
			return
		case constants.InvalidUserID:
			responseUtils.SendBadRequest(ctx, []genericModels.ErrorMessage{
				{Key: constants.User_id, ErrorMessage: constants.InvalidUserID},
			})
			return
		default:
			responseUtils.SendInternalServerError(ctx, err)
			return
		}
	}

	resp := models.BFFPlaylistResponse{
		SuccessMessage: constants.SuccessfullyCreatedPlaylist,
	}
	respData := map[string]interface{}{
		"playlist_id": playlistID,
		"message":     resp.SuccessMessage,
	}
	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, respData)
}
