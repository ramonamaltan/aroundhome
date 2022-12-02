package main

import (
	"log"

	"github.com/ramonamaltan/go-api/internal/api"
	"github.com/ramonamaltan/go-api/internal/db"
)

func main() {
	db := db.Init()
	r := api.SetupRoutes(db)

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal("failed to run")
	}

}
