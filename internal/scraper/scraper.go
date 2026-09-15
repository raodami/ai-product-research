package scraper

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Website     string
	Category    string
	Price       string
	Users       string
	CreatedAt   time.Time
}

func FetchToolify(category string) ([]*Product, error) {
	// Mock data for MVP
	// In production, would scrape https://www.toolify.ai/category/{category}
	return []*Product{
		{ID: "1", Name: "AI Resume Builder", Description: "ATS-optimized resume builder with AI analysis", Website: "https://example.com/resume", Category: "HR", Price: "$29/mo", Users: "50K+"},
		{ID: "2", Name: "AI Analytics Dashboard", Description: "Natural language to SQL and charts", Website: "https://example.com/analytics", Category: "Analytics", Price: "$49/mo", Users: "30K+"},
		{ID: "3", Name: "AI Audio Tools", Description: "Speech-to-text, NLP, and TTS pipeline", Website: "https://example.com/audio", Category: "Audio", Price: "Free", Users: "100K+"},
	}, nil
}

func FetchProductHunt() ([]*Product, error) {
	// Mock data for MVP
	return []*Product{
		{ID: "ph1", Name: "Notion AI", Description: "AI-powered workspace", Website: "https://notion.so", Category: "Productivity", Price: "$10/mo", Users: "10M+"},
		{ID: "ph2", Name: "Arc Browser", Description: "Intelligent web browser", Website: "https://arc.net", Category: "Browser", Price: "Free", Users: "500K+"},
		{ID: "ph3", Name: "Gamma", Description: "AI presentation maker", Website: "https://gamma.app", Category: "Design", Price: "$20/mo", Users: "1M+"},
	}, nil
}

func FetchCompetitorAnalysis(company string) (string, error) {
	// Mock competitor analysis
	return fmt.Sprintf("Competitors for %s:\n1. Company A - Direct competitor, similar features\n2. Company B - Alternative approach\n3. Company C - Different market segment", company), nil
}

func FetchTrends() ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{"keyword": "AI writing assistant", "volume": 50000, "growth": 35.5, "competition": "medium"},
		{"keyword": "AI video generation", "volume": 30000, "growth": 85.2, "competition": "low"},
		{"keyword": "AI code reviewer", "volume": 20000, "growth": 45.8, "competition": "medium"},
		{"keyword": "AI marketing copy", "volume": 40000, "growth": 28.3, "competition": "high"},
	}, nil
}

func FetchCategoryKeywords(category string) ([]string, error) {
	keywords := map[string][]string{
		"HR": {"recruitment", "resume screening", "employee onboarding", "performance management"},
		"Analytics": {"data visualization", "business intelligence", "dashboard", "reporting"},
		"Audio": {"speech to text", "text to speech", "voice cloning", "audio transcription"},
		"Productivity": {"task management", "note taking", "calendar", "collaboration"},
		"Design": {"UI design", "prototyping", "wireframing", "design system"},
	}
	return keywords[strings.ToLower(category)], nil
}

func FetchWebsiteInfo(url string) (map[string]interface{}, error) {
	// Mock website analysis
	return map[string]interface{}{
		"domain_authority": 65,
		"monthly_traffic": "150K",
		"tech_stack": "Next.js, React, Node.js",
		"founded": "2023",
		"team_size": "10-50",
		"funding": "$2M Seed",
	}, nil
}

func SimpleFetch(url string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
