package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var uReq CreateUserReq
	if err := c.ShouldBindJSON(&uReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var uResp *CreateUserResp
	uResp, err := h.Service.CreateUser(c, &uReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, uResp)
}
