package main

import (
	"log"

	"github.com/mohamadafzal06/simple-chat/internal/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("cannot initialize database connection: %v\n", err)
	}
}
