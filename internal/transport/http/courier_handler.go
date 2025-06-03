package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
	"github.com/nemcs/flower-marketplace-api/internal/service"
	"net/http"
)

type CourierHandler struct {
	svc *service.CourierService
}

func NewCourierHandler(router *gin.Engine, svc *service.CourierService) {
	h := &CourierHandler{svc: svc}
	g := router.Group("/couriers")
	g.POST("/", h.Create)
	g.GET("/", h.List)
}

func (h *CourierHandler) Create(c *gin.Context) {
	var courier domain.Courier
	if err := c.ShouldBindJSON(&courier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Create(c.Request.Context(), &courier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create courier"})
		return
	}

	c.JSON(http.StatusCreated, courier)
}

func (h *CourierHandler) List(c *gin.Context) {
	couriers, err := h.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch couriers"})
		return
	}
	c.JSON(http.StatusOK, couriers)
}
