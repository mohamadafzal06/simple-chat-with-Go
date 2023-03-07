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

func (h *Handler) Login(c *gin.Context) {
	var user LoginUserReq

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.SetCookie("jwt", u.accessToken, 3600, "/", "localhost", false, true)

	resp := &LoginUserResp{
		Username: u.Username,
		ID:       u.ID,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"messgae": "logout successful"})
}
