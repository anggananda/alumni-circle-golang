package handlers

import (
	"alumni-circle-api/models"
	"alumni-circle-api/services"
	"alumni-circle-api/utils"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EventListResponse struct {
	Datas   []models.Event `json:"datas"`
	IsNext  bool           `json:"is_next"`
	Message string         `json:"message"`
}

type EventSingleResponse struct {
	Datas   models.Event `json:"datas"`
	Message string       `json:"message"`
}

type EventHandler struct {
	EventService *services.EventService
}

func NewEventHandler(service *services.EventService) *EventHandler {
	return &EventHandler{
		EventService: service,
	}
}

// GetAllEvent godoc
// @Summary List Event
// @Description Mendapatkan semua event dengan pagination dan optional search
// @Tags event
// @Accept  json
// @Produce  json
// @Param page query int false "Page number"
// @Param search query string false "Search keyword"
// @Success 200 {object} EventListResponse
// @Failure 400 {object} map[string]string
// @Router /private/event [get]
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

// GetEventByCategory godoc
// @Summary List Event By Category
// @Description Mendapatkan semua event berdasarkan kategori event dengan pagination dan optional search
// @Tags event
// @Accept  json
// @Produce  json
// @Param page query int false "Page number"
// @Param search query string false "Search keyword"
// @Success 200 {object} EventListResponse
// @Failure 400 {object} map[string]string
// @Router /private/event/category/:idCategory [get]
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

// GetEventByID godoc
// @Summary List Event By ID
// @Description Mendapatkan semua event berdasarkan id event
// @Tags event
// @Accept  json
// @Produce  json
// @Success 200 {object} EventSingleResponse
// @Failure 400 {object} map[string]string
// @Router /private/event/:idEvent [get]
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
