package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"ai-product-research/internal/api"
	"ai-product-research/internal/store"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/products.db"
	}

	store, err := store.NewDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer store.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	api.SetupRoutes(r, store)

	log.Printf("AI Product Research starting on :%s", port)
	log.Fatal(r.Run(":" + port))
}
