// HTTP API (Gin)
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
	"github.com/nemcs/flower-marketplace-api/internal/service"
	"net/http"
)

type ShopHandler struct {
	svc *service.ShopService
}

func NewShopHandler(router *gin.Engine, svc *service.ShopService) {
	h := &ShopHandler{svc: svc}
	g := router.Group("/shops")
	g.POST("/", h.Create)
	g.GET("/", h.List)
}

func (h *ShopHandler) Create(c *gin.Context) {
	var shop domain.Shop
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Create(c.Request.Context(), &shop); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create shop"})
		return
	}

	c.JSON(http.StatusCreated, shop)
}

func (h *ShopHandler) List(c *gin.Context) {
	shops, err := h.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch shops"})
		return
	}
	c.JSON(http.StatusOK, shops)
}
