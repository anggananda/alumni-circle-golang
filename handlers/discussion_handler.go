package handlers

import (
	"alumni-circle-api/services"
	"alumni-circle-api/utils"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DiscussionHandler struct {
	DiscussionService *services.DiscussionService
}

func NewDiscussionHandler(service *services.DiscussionService) *DiscussionHandler {
	return &DiscussionHandler{
		DiscussionService: service,
	}
}

func (h *DiscussionHandler) GetAllDiscussion(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	search := c.DefaultQuery("search", "")
	limit := 10

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offset := (page - 1) * limit

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	discussion, totalItems, err := h.DiscussionService.GetAllDiscussion(ctx, limit, offset, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isNext := (offset + limit) < int(totalItems)

	c.JSON(http.StatusOK, gin.H{
		"datas":   discussion,
		"is_next": isNext,
		"message": "OK",
	})
}

func (h *DiscussionHandler) GetDiscussionByID(c *gin.Context) {
	idDiskusi, err := utils.StringToInt64(c.Param("idDiskusi"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	discussion, err := h.DiscussionService.GetDiscussionByID(ctx, idDiskusi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": discussion, "message": "OK"})
}
