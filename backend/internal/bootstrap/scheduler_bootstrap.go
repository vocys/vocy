package bootstrap

import (
	"context"
	"net/http"

	"github.com/vocys/vocy/internal/job"
	"gorm.io/gorm"
)

func registerScheduledJobs(ctx context.Context, db *gorm.DB, svc *services, httpClient *http.Client, scheduler *job.Scheduler) error {

	return nil
}
