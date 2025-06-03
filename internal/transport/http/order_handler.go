package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
	"github.com/nemcs/flower-marketplace-api/internal/service"
	"net/http"
)

type OrderHandler struct {
	svc *service.OrderService
}

func NewOrderHandler(router *gin.Engine, svc *service.OrderService) {
	h := &OrderHandler{svc: svc}
	g := router.Group("/orders")
	g.POST("/", h.Create)
	g.GET("/", h.List)
}

func (h *OrderHandler) Create(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Create(c.Request.Context(), &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) List(c *gin.Context) {
	orders, err := h.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
