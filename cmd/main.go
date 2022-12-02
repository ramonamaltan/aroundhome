package main

import (
	"github.com/ramonamaltan/go-api/internal/api"
	"log"
)

func main() {
	r := api.SetupRoutes()

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal("failed to run")
	}
}
