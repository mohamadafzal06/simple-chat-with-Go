package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamadafzal06/simple-chat/internal/user"
)

//TODO: add struct for it
var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r := gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
}

func Start(addr string) error {
	return r.Run(addr)
}
