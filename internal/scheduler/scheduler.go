package scheduler

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/radityacandra/go-cms/internal/application/tag/handler"
	"github.com/radityacandra/go-cms/internal/core"
	"go.uber.org/zap"
)

func InitScheduler(ctx context.Context, deps *core.Dependency) {
	deps.Logger.Info("Initiating scheduler...")
	s, err := gocron.NewScheduler()
	if err != nil {
		deps.Logger.Fatal("Failed to initiate scheduler", zap.Error(err))
		return
	}

	handler := handler.NewHandler(deps)

	_, err = s.NewJob(
		gocron.DurationJob(1*time.Hour),
		gocron.NewTask(handler.CalculateTrendingScore, ctx),
	)
	if err != nil {
		deps.Logger.Fatal("Failed to register task: CalculateTrendingScore", zap.Error(err))
		return
	}

	deps.Gocron = s

	s.Start()
	deps.Logger.Info("Scheduler initiation completed successfully")
}
