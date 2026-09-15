package analyzer

import (
	"fmt"
	"math/rand"
	"time"

	"ai-product-research/internal/store"
)

type Opportunity struct {
	Keyword     string  `json:"keyword"`
	MarketSize  string  `json:"market_size"`
	Competition string  `json:"competition"`
	Score       float64 `json:"score"`
	Suggestion  string  `json:"suggestion"`
}

type CompetitorAnalysis struct {
	Company     string  `json:"company"`
	Strengths   string  `json:"strengths"`
	Weaknesses  string  `json:"weaknesses"`
	Differentiation string `json:"differentiation"`
}

func AnalyzeOpportunity(keyword string) (*Opportunity, error) {
	// Mock AI analysis
	scores := rand.Float64()*40 + 60 // 60-100
	
	return &Opportunity{
		Keyword:    keyword,
		MarketSize: fmt.Sprintf("$%dM+% TAM", int(rand.Float64()*50+10)),
		Competition: map[int]string{5: "low", 6: "medium", 7: "high"}[rand.Intn(3)+5],
		Score:      scores,
		Suggestion: fmt.Sprintf("Target %s niche with AI-powered solution. Focus on automation and personalization.", keyword),
	}, nil
}

func AnalyzeCompetitor(company string) (*CompetitorAnalysis, error) {
	return &CompetitorAnalysis{
		Company:       company,
		Strengths:     "Strong brand, large user base, good UX",
		Weaknesses:    "Limited customization, expensive pricing, slow feature updates",
		Differentiation: "Focus on AI-first approach with better personalization",
	}, nil
}

func GenerateReport(products []*store.Product, opportunities []*Opportunity) (string, error) {
	report := "# AI Product Research Report\n\n"
	report += fmt.Sprintf("Generated: %s\n\n", time.Now().Format("2006-01-02 15:04"))
	
	report += "## Top Opportunities\n\n"
	for _, opp := range opportunities {
		report += fmt.Sprintf("### %s (Score: %.1f)\n\n", opp.Keyword, opp.Score)
		report += fmt.Sprintf("- **Market Size**: %s\n", opp.MarketSize)
		report += fmt.Sprintf("- **Competition**: %s\n", opp.Competition)
		report += fmt.Sprintf("- **Suggestion**: %s\n\n", opp.Suggestion)
	}
	
	report += "## Product Landscape\n\n"
	for _, p := range products {
		report += fmt.Sprintf("### %s (%s)\n", p.Name, p.Category)
		report += fmt.Sprintf("%s\n", p.Description)
		report += fmt.Sprintf("**Price**: %s | **Users**: %s\n\n", p.Price, p.Users)
	}
	
	return report, nil
}

func ExportCSV(products []*store.Product) (string, error) {
	csv := "Name,Category,Price,Users,Score,Website,Description\n"
	for _, p := range products {
		csv += fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\",%.1f,\"%s\",\"%s\"\n",
			p.Name, p.Category, p.Price, p.Users, p.Score, p.Website, p.Description)
	}
	return csv, nil
}
