package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vocys/vocy/internal/job"
	"github.com/vocys/vocy/internal/service"
	"github.com/vocys/vocy/internal/storage"
	"gorm.io/gorm"
)

type services struct {
	appConfigService *service.AppConfigService
	fileStorage      storage.FileStorage

	appLockService *service.AppLockService
}

// Initializes all services
func initServices(ctx context.Context, db *gorm.DB, httpClient *http.Client, fileStorage storage.FileStorage, scheduler *job.Scheduler) (svc *services, err error) {
	svc = &services{}
	svc.appConfigService, err = service.NewAppConfigService(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to create app config service: %w", err)
	}

	svc.fileStorage = fileStorage

	svc.appLockService = service.NewAppLockService(db)

	return svc, nil
}
