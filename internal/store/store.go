package store

import (
	"database/sql"
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
	Founded     string    `json:"founded"`
	Features    string    `json:"features"`
	Score       float64   `json:"score"`
	Source      string    `json:"source"` // toolify/producthunt
	CreatedAt   time.Time `json:"created_at"`
}

type Competitor struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Website     string    `json:"website"`
	Price       string    `json:"price"`
	Strengths   string    `json:"strengths"`
	Weaknesses  string    `json:"weaknesses"`
	Ranking     int       `json:"ranking"`
	CreatedAt   time.Time `json:"created_at"`
}

type Trend struct {
	ID           string    `json:"id"`
	Keyword      string    `json:"keyword"`
	SearchVolume int       `json:"search_volume"`
	GrowthRate   float64   `json:"growth_rate"`
	Competition  string    `json:"competition"`
	Opportunity  float64   `json:"opportunity"`
	Category     string    `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
}

type Store struct {
	db *sql.DB
}

func NewDB(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	store := &Store{db: db}
	if err := store.createTables(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) createTables() error {
	schema := `
	CREATE TABLE IF NOT EXISTS products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		website TEXT,
		category TEXT,
		price TEXT DEFAULT '',
		users TEXT DEFAULT '',
		founded TEXT DEFAULT '',
		features TEXT,
		score REAL DEFAULT 0,
		source TEXT DEFAULT 'toolify',
		created_at INTEGER DEFAULT 0
	);
	CREATE TABLE IF NOT EXISTS competitors (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		website TEXT,
		price TEXT DEFAULT '',
		strengths TEXT,
		weaknesses TEXT,
		ranking INTEGER DEFAULT 0,
		created_at INTEGER DEFAULT 0
	);
	CREATE TABLE IF NOT EXISTS trends (
		id TEXT PRIMARY KEY,
		keyword TEXT NOT NULL,
		search_volume INTEGER DEFAULT 0,
		growth_rate REAL DEFAULT 0,
		competition TEXT DEFAULT 'medium',
		opportunity REAL DEFAULT 0,
		category TEXT DEFAULT '',
		created_at INTEGER DEFAULT 0
	);
	CREATE INDEX IF NOT EXISTS idx_products_category ON products(category);
	CREATE INDEX IF NOT EXISTS idx_products_score ON products(score DESC);
	CREATE INDEX IF NOT EXISTS idx_trends_keyword ON trends(keyword);
	`
	_, err := s.db.Exec(schema)
	return err
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) SaveProduct(p *Product) error {
	_, err := s.db.Exec(
		"INSERT OR REPLACE INTO products (id, name, description, website, category, price, users, founded, features, score, source, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		p.ID, p.Name, p.Description, p.Website, p.Category, p.Price, p.Users, p.Founded, p.Features, p.Score, p.Source, time.Now().Unix(),
	)
	return err
}

func (s *Store) GetProducts(category string, limit int) ([]*Product, error) {
	var sqlStr string
	var args []interface{}
	
	if category != "" {
		sqlStr = "SELECT id, name, description, website, category, price, users, founded, features, score, source, created_at FROM products WHERE category = ? ORDER BY score DESC LIMIT ?"
		args = append(args, category, limit)
	} else {
		sqlStr = "SELECT id, name, description, website, category, price, users, founded, features, score, source, created_at FROM products ORDER BY score DESC LIMIT ?"
		args = append(args, limit)
	}

	rows, err := s.db.Query(sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var p Product
		var createdAt int64
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Website, &p.Category, &p.Price, &p.Users, &p.Founded, &p.Features, &p.Score, &p.Source, &createdAt); err != nil {
			return nil, err
		}
		p.CreatedAt = time.Unix(createdAt, 0)
		products = append(products, &p)
	}
	return products, nil
}

func (s *Store) GetProduct(id string) (*Product, error) {
	var p Product
	var createdAt int64
	err := s.db.QueryRow("SELECT id, name, description, website, category, price, users, founded, features, score, source, created_at FROM products WHERE id = ?", id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Website, &p.Category, &p.Price, &p.Users, &p.Founded, &p.Features, &p.Score, &p.Source, &createdAt)
	if err != nil {
		return nil, err
	}
	p.CreatedAt = time.Unix(createdAt, 0)
	return &p, nil
}

func (s *Store) DeleteProduct(id string) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func (s *Store) SaveCompetitor(c *Competitor) error {
	_, err := s.db.Exec(
		"INSERT OR REPLACE INTO competitors (id, name, website, price, strengths, weaknesses, ranking, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		c.ID, c.Name, c.Website, c.Price, c.Strengths, c.Weaknesses, c.Ranking, time.Now().Unix(),
	)
	return err
}

func (s *Store) GetCompetitors(limit int) ([]*Competitor, error) {
	rows, err := s.db.Query("SELECT id, name, website, price, strengths, weaknesses, ranking, created_at FROM competitors ORDER BY ranking ASC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []*Competitor
	for rows.Next() {
		var c Competitor
		var createdAt int64
		if err := rows.Scan(&c.ID, &c.Name, &c.Website, &c.Price, &c.Strengths, &c.Weaknesses, &c.Ranking, &createdAt); err != nil {
			return nil, err
		}
		c.CreatedAt = time.Unix(createdAt, 0)
		competitors = append(competitors, &c)
	}
	return competitors, nil
}

func (s *Store) SaveTrend(t *Trend) error {
	_, err := s.db.Exec(
		"INSERT OR REPLACE INTO trends (id, keyword, search_volume, growth_rate, competition, opportunity, category, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		t.ID, t.Keyword, t.SearchVolume, t.GrowthRate, t.Competition, t.Opportunity, t.Category, time.Now().Unix(),
	)
	return err
}

func (s *Store) GetTrends(limit int) ([]*Trend, error) {
	rows, err := s.db.Query("SELECT id, keyword, search_volume, growth_rate, competition, opportunity, category, created_at FROM trends ORDER BY opportunity DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trends []*Trend
	for rows.Next() {
		var t Trend
		var createdAt int64
		if err := rows.Scan(&t.ID, &t.Keyword, &t.SearchVolume, &t.GrowthRate, &t.Competition, &t.Opportunity, &t.Category, &createdAt); err != nil {
			return nil, err
		}
		t.CreatedAt = time.Unix(createdAt, 0)
		trends = append(trends, &t)
	}
	return trends, nil
}

func (s *Store) GetAnalytics() (map[string]interface{}, error) {
	var totalProducts int
	s.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&totalProducts)

	var avgScore float64
	s.db.QueryRow("SELECT COALESCE(AVG(score), 0) FROM products").Scan(&avgScore)

	var categoryCounts []struct {
		Category string
		Count    int
	}
	rows, _ := s.db.Query("SELECT category, COUNT(*) as count FROM products GROUP BY category ORDER BY count DESC LIMIT 5")
	defer rows.Close()
	for rows.Next() {
		var c struct{ Category string; Count int }
		rows.Scan(&c.Category, &c.Count)
		categoryCounts = append(categoryCounts, c)
	}

	return map[string]interface{}{
		"total_products":   totalProducts,
		"avg_score":        avgScore,
		"top_categories":   categoryCounts,
		"total_competitors": len(categoryCounts),
	}, nil
}
