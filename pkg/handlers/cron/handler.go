package cronHandler

import (
	"context"
	"log"
	"time"
)

type (
	cronHandler struct {
	}

	CronHandler interface {
		StartCron(ctx context.Context)
	}
)

// StartCron implements CronHandler.
func (c *cronHandler) StartCron(ctx context.Context) {
	ticker := time.NewTicker(50 * time.Second)
	defer ticker.Stop()

	log.Println("INFO: Cron job scheduler started. Press Ctrl+C to stop.")

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			return
		}
	}
}

// NewCronHandler creates a new CronHandler instance.
func NewCronHandler() CronHandler {
	return &cronHandler{}
}
