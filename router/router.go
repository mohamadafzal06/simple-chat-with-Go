package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamadafzal06/simple-chat/internal/user"
	"github.com/mohamadafzal06/simple-chat/internal/ws"
)

//TODO: add struct for it
var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r := gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/createRoom", wsHandler.CreateRoom)
}

func Start(addr string) error {
	return r.Run(addr)
}
