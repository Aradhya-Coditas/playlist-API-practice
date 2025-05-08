package handlers

import (
	"admin-app/watchlist/business"
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"errors"
	"fmt"
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
	"gorm.io/gorm"
)

type AdSongToPlaylistHandler struct {
	service *business.AdSongToPlaylistService
}

func NewAdSongToPlaylistHandler(service *business.AdSongToPlaylistService) *AdSongToPlaylistHandler {
	return &AdSongToPlaylistHandler{
		service: service,
	}
}

func (h *AdSongToPlaylistHandler) HandleAdToSongPlaylist(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleAdToSongPlaylist")
	defer span.End()

	log := logger.GetLogger(ctx)

	var req models.BFFAdPlaylistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.sendInvalidRequestResponse(ctx, log, err)
		return
	}

	if err := validations.GetBFFValidator(spanCtx).Struct(&req); err != nil {
		h.sendValidationErrorResponse(ctx, log, err)
		return
	}

	response, err := h.service.ModifyPlaylistSongs(ctx, spanCtx, req)
	if err != nil {
		h.handleServiceError(ctx, log, err)
		return
	}

	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, response)
}

func (h *AdSongToPlaylistHandler) sendInvalidRequestResponse(ctx *gin.Context, log logger.Logger, err error) {
	log.Error("JSON Binding Error", zap.Error(err))
	errorMsg := genericModel.ErrorMessage{
		Key:          constants.Name,
		ErrorMessage: genericConstants.JsonBindingFieldError,
	}
	responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsg})
}

func (h *AdSongToPlaylistHandler) sendValidationErrorResponse(ctx *gin.Context, log logger.Logger, err error) {
	validationErrors, _ := validations.FormatValidationErrors(ctx.Request.Context(), err.(validator.ValidationErrors))
	log.Error("Validation Error", zap.Error(err))
	responseUtils.SendBadRequest(ctx, validationErrors)
}

func (h *AdSongToPlaylistHandler) handleServiceError(ctx *gin.Context, log logger.Logger, err error) {
	log.Error("Service Error", zap.Error(err))

	if strings.Contains(err.Error(), constants.ForeignKeyViolationError) {
		if strings.Contains(err.Error(), constants.InvalidPlaylistIdExistError) {
			responseUtils.SendNotFoundJSON(ctx, fmt.Errorf(constants.InvalidPlaylistIdExistError))
			return
		}
		if strings.Contains(err.Error(), constants.InvalidSongIdsError) {
			responseUtils.SendNotFoundJSON(ctx, fmt.Errorf(constants.InvalidSongIdsError))
			return
		}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		responseUtils.SendNotFoundJSON(ctx, fmt.Errorf(constants.InvalidPlaylistIdExistError))
		return
	}

	if strings.Contains(err.Error(), genericConstants.DuplicateKeyError) {
		responseUtils.SendConflict(ctx, constants.DuplicatePlaylistError)
		return
	}

	responseUtils.SendInternalServerError(ctx, err)
}
