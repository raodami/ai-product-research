package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"ai-product-research/internal/analyzer"
	"ai-product-research/internal/scraper"
	"ai-product-research/internal/store"
)

func SetupRoutes(r *gin.Engine, s *store.Store) {
	r.Use(corsMiddleware())

	// Scraper endpoints
	r.GET("/api/scraper/toolify", func(c *gin.Context) {
		products, err := scraper.FetchToolify("")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	r.POST("/api/scraper/toolify", func(c *gin.Context) {
		var req struct {
			Category string `json:"category"`
		}
		c.ShouldBindJSON(&req)

		products, err := scraper.FetchToolify(req.Category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, p := range products {
			p.ID = uuid.New().String()
			s.SaveProduct(&store.Product{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Website:     p.Website,
				Category:    p.Category,
				Price:       p.Price,
				Users:       p.Users,
			})
		}
		c.JSON(http.StatusOK, gin.H{"saved": len(products)})
	})

	r.GET("/api/scraper/producthunt", func(c *gin.Context) {
		products, err := scraper.FetchProductHunt()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	// Products endpoints
	r.GET("/api/products", func(c *gin.Context) {
		category := c.Query("category")
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
		products, err := s.GetProducts(category, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	r.GET("/api/products/:id", func(c *gin.Context) {
		product, err := s.GetProduct(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	})

	r.DELETE("/api/products/:id", func(c *gin.Context) {
		s.DeleteProduct(c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	})

	// Trends endpoint
	r.GET("/api/trends", func(c *gin.Context) {
		trends, err := scraper.FetchTrends()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, trends)
	})

	// Analytics endpoint
	r.GET("/api/analytics", func(c *gin.Context) {
		analytics, err := s.GetAnalytics()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, analytics)
	})

	// Competitors endpoint
	r.GET("/api/competitors", func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		competitors, err := s.GetCompetitors(limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, competitors)
	})

	// Analysis endpoints
	r.POST("/api/analyze/opportunity", func(c *gin.Context) {
		var req struct {
			Keyword string `json:"keyword" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		opp, err := analyzer.AnalyzeOpportunity(req.Keyword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, opp)
	})

	r.POST("/api/analyze/competitor", func(c *gin.Context) {
		var req struct {
			Company string `json:"company" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		comp, err := analyzer.AnalyzeCompetitor(req.Company)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, comp)
	})

	// Export endpoints
	r.GET("/api/export/csv", func(c *gin.Context) {
		products, _ := s.GetProducts("", 100)
		csv, err := analyzer.ExportCSV(products)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment; filename=products.csv")
		c.String(http.StatusOK, csv)
	})

	r.GET("/api/export/report", func(c *gin.Context) {
		products, _ := s.GetProducts("", 50)
		trends, _ := s.GetTrends(10)
		
		report, err := analyzer.GenerateReport(products, trends)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Header("Content-Type", "text/markdown")
		c.Header("Content-Disposition", "attachment; filename=report.md")
		c.String(http.StatusOK, report)
	})

	// Templates
	r.GET("/api/templates", func(c *gin.Context) {
		templates := []map[string]string{
			{"id": "saas", "name": "SaaS Tool", "description": "For subscription-based SaaS products"},
			{"id": "ai-tool", "name": "AI Tool", "description": "For AI-powered tools and apps"},
			{"id": "dev-tool", "name": "Developer Tool", "description": "For developer productivity tools"},
			{"id": "productivity", "name": "Productivity", "description": "For productivity and workflow tools"},
		}
		c.JSON(http.StatusOK, templates)
	})
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
