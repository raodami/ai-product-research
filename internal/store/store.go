package store

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	Category    string    `json:"category"`
	Price       string    `json:"price"`
	Users       string    `json:"users"`
	CreatedAt   time.Time `json:"created_at"`
}

type Competitor struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Website    string    `json:"website"`
	Strengths  string    `json:"strengths"`
	Weaknesses string    `json:"weaknesses"`
	Ranking    int       `json:"ranking"`
	CreatedAt  time.Time `json:"created_at"`
}

type Trend struct {
	ID           string    `json:"id"`
	Keyword      string    `json:"keyword"`
	Month        string    `json:"month"`
	SearchVolume int       `json:"search_volume"`
	GrowthRate   float64   `json:"growth_rate"`
	Competition  string    `json:"competition"`
	Opportunity  float64   `json:"opportunity"`
	CreatedAt    time.Time `json:"created_at"`
}

type Store struct {
	db        *sql.DB
	productMu sync.RWMutex
}

func New(dbPath string) (*Store, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	store := &Store{db: db}
	if err := store.initSchema(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) initSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		website TEXT,
		category TEXT,
		price TEXT,
		users INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS competitors (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		website TEXT,
		strengths TEXT,
		weaknesses TEXT,
		ranking INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS trends (
		id TEXT PRIMARY KEY,
		keyword TEXT NOT NULL,
		month TEXT,
		search_volume INTEGER,
		growth_rate REAL,
		competition TEXT,
		opportunity REAL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := s.db.Exec(schema)
	return err
}

func (s *Store) SaveProduct(p *Product) error {
	s.productMu.Lock()
	defer s.productMu.Unlock()
	_, err := s.db.Exec(
		"INSERT OR REPLACE INTO products (id, name, description, website, category, price, users) VALUES (?, ?, ?, ?, ?, ?, ?)",
		p.ID, p.Name, p.Description, p.Website, p.Category, p.Price, p.Users,
	)
	return err
}

func (s *Store) GetProducts(category string, limit int) ([]*Product, error) {
	return s.SearchProducts("", category, limit)
}

func (s *Store) SearchProducts(query, category string, limit int) ([]*Product, error) {
	s.productMu.RLock()
	defer s.productMu.RUnlock()

	queryStr := "SELECT id, name, description, website, category, price, users FROM products WHERE 1=1"
	args := []interface{}{}

	if query != "" {
		queryStr += " AND (name LIKE ? OR description LIKE ?)"
		args = append(args, "%"+query+"%", "%"+query+"%")
	}
	if category != "" {
		queryStr += " AND category = ?"
		args = append(args, category)
	}
	queryStr += " ORDER BY created_at DESC LIMIT ?"
	args = append(args, limit)

	rows, err := s.db.Query(queryStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Website, &p.Category, &p.Price, &p.Users); err != nil {
			continue
		}
		products = append(products, &p)
	}
	return products, rows.Err()
}

func (s *Store) GetAnalytics() (map[string]interface{}, error) {
	var productCount int
	var categoryCount int
	s.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&productCount)
	s.db.QueryRow("SELECT COUNT(DISTINCT category) FROM products").Scan(&categoryCount)

	return map[string]interface{}{
		"total_products":   productCount,
		"total_categories": categoryCount,
		"avg_users":        0,
	}, nil
}

func (s *Store) GetCompetitors(limit int) ([]*Competitor, error) {
	rows, err := s.db.Query("SELECT id, name, website, strengths, weaknesses, ranking FROM competitors ORDER BY ranking ASC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []*Competitor
	for rows.Next() {
		var c Competitor
		if err := rows.Scan(&c.ID, &c.Name, &c.Website, &c.Strengths, &c.Weaknesses, &c.Ranking); err != nil {
			continue
		}
		competitors = append(competitors, &c)
	}
	return competitors, rows.Err()
}

func (s *Store) GetTrends(limit int) ([]*Trend, error) {
	rows, err := s.db.Query("SELECT id, keyword, month, search_volume, growth_rate, competition, opportunity FROM trends ORDER BY created_at DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trends []*Trend
	for rows.Next() {
		var t Trend
		if err := rows.Scan(&t.ID, &t.Keyword, &t.Month, &t.SearchVolume, &t.GrowthRate, &t.Competition, &t.Opportunity); err != nil {
			continue
		}
		trends = append(trends, &t)
	}
	return trends, rows.Err()
}

func (s *Store) DeleteProduct(id string) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func (s *Store) Close() error {
	return s.db.Close()
}
