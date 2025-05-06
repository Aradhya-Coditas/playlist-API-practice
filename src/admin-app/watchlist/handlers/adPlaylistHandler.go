package handlers

import (
	"admin-app/watchlist/business"
	"admin-app/watchlist/commons/constants"
	"admin-app/watchlist/models"
	"net/http"
	genericConstants "omnenest-backend/src/constants"

	"github.com/gin-gonic/gin"
)

type AdPlaylistHandler struct {
	service *business.AdPlaylistService
}

func NewAdPlaylistHandler(service *business.AdPlaylistService) *AdPlaylistHandler {
	return &AdPlaylistHandler{service: service}
}

func (h *AdPlaylistHandler) HandleModifyPlaylistSongs() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.BFFAdPlaylistRequest

		// Parse JSON body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": genericConstants.InvalidRequestFormatError})
			return
		}

		// Validate required fields
		if req.Action != constants.ActionAdd && req.Action != constants.ActionDelete {
			c.JSON(http.StatusBadRequest, gin.H{"error": genericConstants.InvalidActionChoice})
			return
		}
		if len(req.SongIDs) == 0 || req.PlaylistID == 0 || req.UserID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": genericConstants.MissingRequiredFields})
			return
		}

		response, err := h.service.ModifyPlaylistSongs(c.Request.Context(), c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
