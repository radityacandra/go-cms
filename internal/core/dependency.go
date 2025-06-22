package core

import (
	"context"
	"errors"
	"syscall"

	"github.com/go-co-op/gocron/v2"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/pkg/database"
	"go.uber.org/zap"
)

type Dependency struct {
	Logger *zap.Logger
	DB     *database.DB
	Config *Config
	Echo   *echo.Echo
	Gocron gocron.Scheduler
}

func NewDependency(logger *zap.Logger, db *database.DB, config *Config) *Dependency {
	return &Dependency{
		Logger: logger,
		DB:     db,
		Config: config,
	}
}

func (d *Dependency) GracefulShutdown(ctx context.Context) int {
	<-ctx.Done()
	code := 0
	d.Logger.Info("Gracefully shutting down web server...")
	err := d.Echo.Shutdown(ctx)
	if err != nil {
		d.Logger.Error("failed to close server", zap.Error(err))
		code = 1
	} else {
		d.Logger.Info("web server shutted down")
	}

	d.Logger.Info("Gracefully shutting down db connection...")
	err = d.DB.Close()
	if err != nil {
		d.Logger.Error("failed to close database connection", zap.Error(err))
		code = 1
	} else {
		d.Logger.Info("success to close database connection")
	}

	err = d.Gocron.Shutdown()
	if err != nil {
		d.Logger.Error("failed to shutdown cron scheduler", zap.Error(err))
		code = 1
	} else {
		d.Logger.Info("success to close cron scheduler")
	}

	err = d.Logger.Sync()
	if err != nil && !errors.Is(err, syscall.ENOTTY) && !errors.Is(err, syscall.EINVAL) {
		d.Logger.Error("failed to flush log", zap.Error(err))
		code = 1
	} else {
		d.Logger.Info("success to flush log")
	}

	d.Logger.Info("shutted down", zap.Any("code", code))
	return code
}
