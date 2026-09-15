package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"ai-product-research/internal/api"
	"ai-product-research/internal/scheduler"
	"ai-product-research/internal/store"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/products.db"
	}

	s, err := store.New(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer s.Close()

	// 启动爬虫调度器
	sc := scheduler.New(s)
	if err := sc.Start(); err != nil {
		log.Fatalf("Failed to start scheduler: %v", err)
	}
	defer sc.Stop()

	r := gin.Default()
	api.SetupRoutes(r, s, sc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("AI Product Research starting on :%s", port)
	log.Fatal(r.Run(":" + port))
}
