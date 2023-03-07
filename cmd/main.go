package main

import (
	"log"

	"github.com/mohamadafzal06/simple-chat/internal/db"
	"github.com/mohamadafzal06/simple-chat/internal/user"
	"github.com/mohamadafzal06/simple-chat/router"
)

func main() {
	conn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("cannot initialize database connection: %v\n", err)
	}

	db, _ := conn.GetDB()
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	router.InitRouter(userHandler)
	err = router.Start(":8080")
	if err != nil {
		log.Fatalf("cannot initialize the router: %v\n", err)
	}
}
