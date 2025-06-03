package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nemcs/flower-marketplace-api/internal/domain"
	"github.com/nemcs/flower-marketplace-api/internal/service"
	"net/http"
)

type ClientHandler struct {
	svc *service.ClientService
}

func NewClientHandler(router *gin.Engine, svc *service.ClientService) {
	h := &ClientHandler{svc: svc}
	g := router.Group("/clients")
	g.POST("/", h.Create)
	g.GET("/", h.List)
}

func (h *ClientHandler) Create(c *gin.Context) {
	var client domain.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Create(c.Request.Context(), &client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create client"})
		return
	}

	c.JSON(http.StatusCreated, client)
}

func (h *ClientHandler) List(c *gin.Context) {
	clients, err := h.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, clients)
}
