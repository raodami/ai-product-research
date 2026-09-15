package analyzer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type DeepSeekClient struct {
	APIKey string
	BaseURL string
}

type DeepSeekRequest struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
	Temperature float64 `json:"temperature,omitempty"`
	MaxTokens int `json:"max_tokens,omitempty"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type DeepSeekResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func NewDeepSeekClient() *DeepSeekClient {
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	baseURL := os.Getenv("DEEPSEEK_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.deepseek.com/v1"
	}
	return &DeepSeekClient{
		APIKey: apiKey,
		BaseURL: baseURL,
	}
}

func (c *DeepSeekClient) AnalyzeOpportunity(keyword string) (string, error) {
	if c.APIKey == "" {
		return "", fmt.Errorf("DEEPSEEK_API_KEY not configured")
	}

	messages := []Message{
		{
			Role: "system",
			Content: "You are an AI product research analyst. Analyze market opportunities and provide detailed reports.",
		},
		{
			Role: "user",
			Content: fmt.Sprintf("Analyze the market opportunity for: %s\n\nProvide a detailed analysis including:\n1. Market size and growth potential\n2. Competition analysis\n3. Target audience\n4. Key differentiators\n5. Revenue model suggestions\n6. Go-to-market strategy\n\nFormat your response as JSON with these fields: market_size, competition_level, opportunity_score (0-100), reasoning, suggestions (array of strings)", keyword),
		},
	}

	reqBody := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: messages,
		Temperature: 0.7,
		MaxTokens: 1000,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/chat/completions", c.BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var deepSeekResp DeepSeekResponse
	if err := json.Unmarshal(body, &deepSeekResp); err != nil {
		return string(body), err
	}

	if len(deepSeekResp.Choices) > 0 {
		return deepSeekResp.Choices[0].Message.Content, nil
	}

	return string(body), nil
}

func (c *DeepSeekClient) AnalyzeCompetitor(productName string) (string, error) {
	if c.APIKey == "" {
		return "", fmt.Errorf("DEEPSEEK_API_KEY not configured")
	}

	messages := []Message{
		{
			Role: "system",
			Content: "You are a competitive analysis expert. Analyze competitors and identify strengths/weaknesses.",
		},
		{
			Role: "user",
			Content: fmt.Sprintf("Analyze the competitor: %s\n\nProvide analysis of their:\n1. Strengths\n2. Weaknesses\n3. Market positioning\n4. Pricing strategy\n5. Differentiation opportunities\n\nFormat as JSON: strengths, weaknesses, differentiation", productName),
		},
	}

	reqBody := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: messages,
		Temperature: 0.7,
		MaxTokens: 800,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/chat/completions", c.BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var deepSeekResp DeepSeekResponse
	if err := json.Unmarshal(body, &deepSeekResp); err != nil {
		return string(body), err
	}

	if len(deepSeekResp.Choices) > 0 {
		return deepSeekResp.Choices[0].Message.Content, nil
	}

	return string(body), nil
}

func (c *DeepSeekClient) GenerateStrategy(topic string) (string, error) {
	if c.APIKey == "" {
		return "", fmt.Errorf("DEEPSEEK_API_KEY not configured")
	}

	messages := []Message{
		{
			Role: "system",
			Content: "You are a business strategy consultant. Help create go-to-market strategies for AI products.",
		},
		{
			Role: "user",
			Content: fmt.Sprintf("Create a go-to-market strategy for an AI product in the '%s' niche.\n\nInclude:\n1. Target customer segments\n2. Value proposition\n3. Pricing strategy\n4. Marketing channels\n5. Growth tactics\n6. Monetization approach", topic),
		},
	}

	reqBody := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: messages,
		Temperature: 0.8,
		MaxTokens: 1200,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/chat/completions", c.BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var deepSeekResp DeepSeekResponse
	if err := json.Unmarshal(body, &deepSeekResp); err != nil {
		return string(body), err
	}

	if len(deepSeekResp.Choices) > 0 {
		return deepSeekResp.Choices[0].Message.Content, nil
	}

	return string(body), nil
}
