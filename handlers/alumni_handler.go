package handlers

import (
	"alumni-circle-api/services"
	"alumni-circle-api/utils"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AlumniHandler struct {
	AlumniService *services.AlumniService
}

func NewAlumniHandler(service *services.AlumniService) *AlumniHandler {
	return &AlumniHandler{AlumniService: service}
}

func (h *AlumniHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username cannot be empty"})
		return
	}

	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password cannot be empty"})
		return
	}

	if input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email cannot be empty"})
		return
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	_, err := h.AlumniService.GetAlumniByUsername(ctx, input.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exist"})
		return
	}

	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.AlumniService.Register(ctx, input.Username, hashPassword, input.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

func (h *AlumniHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username cannot be empty"})
		return
	}

	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password cannot be empty"})
		return
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	alumni, err := h.AlumniService.GetAlumniByUsername(ctx, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(alumni.Password)
	log.Println(input.Password)

	if !utils.CheckPasswordHash(input.Password, alumni.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password does not match!!"})
		return
	}

	token, expireTime, err := utils.GenerateJWT(alumni.IDAlumni, alumni.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "expireTime": expireTime, "message": "Login Successfully"})
}

func (h *AlumniHandler) GetAllAlumni(c *gin.Context) {
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

	alumni, totalItems, err := h.AlumniService.GetAllAlumni(ctx, limit, offset, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isNext := (offset + limit) < int(totalItems)

	c.JSON(http.StatusOK, gin.H{
		"datas":   alumni,
		"is_next": isNext,
		"message": "OK",
	})
}

func (h *AlumniHandler) GetAlumniByID(c *gin.Context) {
	idAlumni, err := utils.StringToInt64(c.Param("IDAlumni"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	alumni, err := h.AlumniService.GetAlumniByID(ctx, idAlumni)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": alumni, "message": "OK"})
}

func (h *AlumniHandler) DeleteAlumni(c *gin.Context) {
	idAlumni, err := utils.StringToInt64(c.Param("IDAlumni"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	if err := h.AlumniService.DeleteAlumni(ctx, idAlumni); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Delete Alumni!"})
}
