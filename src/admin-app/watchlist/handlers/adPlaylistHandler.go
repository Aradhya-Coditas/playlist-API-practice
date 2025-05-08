package handlers

import (
	"admin-app/watchlist/business"
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"encoding/json"
	"errors"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type adSongToPlaylistController struct {
	service *business.AdSongToPlaylistService
}

func NewAdSongToPlaylistController(service *business.AdSongToPlaylistService) *adSongToPlaylistController {
	return &adSongToPlaylistController{
		service: service,
	}
}

// HandleAdToSongPlaylist.
// @Summary Modify Playlist
// @Description Modify Playlist based on the provided request
// @Tags Playlist
// @Accept json
// @Produce json
// @Param payload body models.BFFAdPlaylistRequest true "Modify Playlist Request"
// @Success 200 {object} models.BFFAdPlaylistResponse "Playlist songs data"
// @Failure 400 {object} models.ErrorMessage "Bad request"
// @Failure 404 {object} models.ErrorMessage "Resource not found"
// @Failure 409 {object} models.ErrorMessage "Conflict"
// @Failure 500 {object} models.ErrorMessage "Internal server error"
// @Router /api/playlist/modify [post]
func (controller *adSongToPlaylistController) HandleAdToSongPlaylist(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleAdToSongPlaylist")
	defer span.End()

	var request models.BFFAdPlaylistRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		var jsonErr *json.UnmarshalTypeError
		if errors.As(err, &jsonErr) {
			responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{
				{Key: jsonErr.Field, ErrorMessage: genericConstants.JsonBindingFieldError},
			})
			return
		}
		responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{
			{Key: "request", ErrorMessage: genericConstants.InvalidRequestFormatError},
		})
		return
	}

	if err := validations.GetBFFValidator(spanCtx).Struct(&request); err != nil {
		validationErrors, _ := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}

	response, err := controller.service.AdSongToPlaylist(ctx, spanCtx, request)
	if err != nil {
		if strings.Contains(err.Error(), constants.InvalidPlaylistIdExistError) {
			responseUtils.SendNotFoundJSON(ctx, fmt.Errorf(constants.InvalidPlaylistIdExistError))
			return
		}
		if strings.Contains(err.Error(), constants.DuplicatePlaylistError) {
			responseUtils.SendConflict(ctx, constants.DuplicatePlaylistError)
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseUtils.SendNotFoundJSON(ctx, fmt.Errorf(constants.InvalidSongIdsError))
			return
		}

		responseUtils.SendInternalServerError(ctx, err)
		return
	}

	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, response)
}
