package handler

import (
	"github.com/radityacandra/go-cms/internal/application/user/repository"
	"github.com/radityacandra/go-cms/internal/application/user/service"
	"github.com/radityacandra/go-cms/internal/core"
	"go.uber.org/zap"
)

type Handler struct {
	Service service.IService
	Logger  *zap.Logger
}

func NewHandler(deps *core.Dependency) *Handler {
	userRepo := repository.NewRepository(deps.DB)
	userService := service.NewService(userRepo)

	return &Handler{
		Service: userService,
		Logger:  deps.Logger,
	}
}
