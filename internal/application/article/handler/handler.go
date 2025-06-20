package handler

import (
	"github.com/radityacandra/go-cms/internal/application/article/repository"
	"github.com/radityacandra/go-cms/internal/application/article/service"
	"github.com/radityacandra/go-cms/internal/core"
	"go.uber.org/zap"
)

type Handler struct {
	Service service.IService
	Logger  *zap.Logger
}

func NewHandler(deps *core.Dependency) *Handler {
	repository := repository.NewRepository(deps.DB)
	service := service.NewService(repository)

	return &Handler{
		Service: service,
		Logger:  deps.Logger,
	}
}
