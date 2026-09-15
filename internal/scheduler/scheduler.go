package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"ai-product-research/internal/scraper"
	"ai-product-research/internal/store"
)

type Manager struct {
	cron   *cron.Cron
	store  *store.Store
}

func New(s *store.Store) *Manager {
	return &Manager{
		cron: cron.New(),
		store: s,
	}
}

func (m *Manager) Start() error {
	// 每天凌晨2点抓取Toolify
	_, err := m.cron.AddFunc("0 2 * * *", func() {
		log.Println("[Scheduler] Running Toolify scraper...")
		products, err := scraper.FetchToolify("")
		if err != nil {
			log.Printf("[Scheduler] Toolify fetch error: %v", err)
			return
		}
		for _, p := range products {
			m.store.SaveProduct(&store.Product{
				ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
				Name:        p.Name,
				Description: p.Description,
				Website:     p.Website,
				Category:    p.Category,
				Price:       p.Price,
				Users:       p.Users,
			})
		}
		log.Printf("[Scheduler] Toolify: saved %d products", len(products))
	})
	if err != nil {
		return fmt.Errorf("failed to add Toolify cron job: %w", err)
	}

	// 每天凌晨3点抓取Product Hunt
	_, err = m.cron.AddFunc("0 3 * * *", func() {
		log.Println("[Scheduler] Running Product Hunt scraper...")
		products, err := scraper.FetchProductHunt()
		if err != nil {
			log.Printf("[Scheduler] Product Hunt fetch error: %v", err)
			return
		}
		for _, p := range products {
			m.store.SaveProduct(&store.Product{
				ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
				Name:        p.Name,
				Description: p.Description,
				Website:     p.Website,
				Category:    p.Category,
				Price:       p.Price,
				Users:       p.Users,
			})
		}
		log.Printf("[Scheduler] Product Hunt: saved %d products", len(products))
	})
	if err != nil {
		return fmt.Errorf("failed to add Product Hunt cron job: %w", err)
	}

	m.cron.Start()
	log.Println("[Scheduler] Scheduler started")
	return nil
}

func (m *Manager) Stop() {
	m.cron.Stop()
	log.Println("[Scheduler] Scheduler stopped")
}

func (m *Manager) GetSchedule() []string {
	return []string{
		"Toolify: Daily at 02:00",
		"Product Hunt: Daily at 03:00",
	}
}

func (m *Manager) GetNextRun() (string, error) {
	next := m.cron.Entries()
	if len(next) == 0 {
		return "No jobs scheduled", nil
	}
	return next[0].Next.Format(time.RFC3339), nil
}
