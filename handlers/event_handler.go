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

type EventHandler struct {
	EventService *services.EventService
}

func NewEventHandler(service *services.EventService) *EventHandler {
	return &EventHandler{
		EventService: service,
	}
}

func (h *EventHandler) GetAllEvent(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	search := c.DefaultQuery("search", "")
	limit := 10

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	event, totalItems, err := h.EventService.GetAllEvent(ctx, limit, offset, search)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isNext := (offset + limit) < int(totalItems)
	c.JSON(http.StatusOK, gin.H{
		"datas":   event,
		"is_next": isNext,
		"message": "OK",
	})
}

func (h *EventHandler) GetEventByCategory(c *gin.Context) {
	idkategori, err := utils.StringToInt64(c.Param("idKategori"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	event, err := h.EventService.GetEventByCategory(ctx, idkategori)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": event, "message": "OK"})
}

func (h *EventHandler) GetEventByID(c *gin.Context) {
	idEvent, err := utils.StringToInt64(c.Param("idEvent"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	event, err := h.EventService.GetEventByID(ctx, idEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": event, "message": "OK"})
}
