# AI Product Research

AI-powered product research tool for discovering opportunities and analyzing competitors.

## Features

- **Toolify Scraper**: Track new products on Toolify.io
- **Product Hunt Monitor**: Discover trending products
- **Competitor Analysis**: Price/features/reviews comparison
- **Trend Detection**: AI-powered opportunity scoring
- **Export Reports**: CSV/PDF/Multi-page analysis

## Tech Stack

- Backend: Go (Gin) + SQLite (modernc.org)
- Frontend: Next.js 14
- AI: DeepSeek API
- Scraping: HTTP client + HTML parsing

## API Endpoints

```
GET  /api/scraper/toolify
POST /api/scraper/toolify
GET  /api/scraper/producthunt
POST /api/scraper/producthunt
GET  /api/products
GET  /api/products/:id
GET  /api/trends
GET  /api/analytics
GET  /api/reports
GET  /api/templates
POST /api/export/csv
POST /api/export/pdf
```

## Analysis Criteria

- **Market Size**: Search volume, category growth
- **Competition**: Number of similar tools
- **Monetization**: Pricing models, revenue estimates
- **Tech Stack**: Technologies used
- **User Reviews**: Sentiment analysis
- **SEO Metrics**: Domain authority, backlinks
