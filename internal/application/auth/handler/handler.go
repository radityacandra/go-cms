package handler

import (
	"github.com/radityacandra/go-cms/internal/application/auth/service"
	"github.com/radityacandra/go-cms/internal/application/user/repository"
	userService "github.com/radityacandra/go-cms/internal/application/user/service"
	"github.com/radityacandra/go-cms/internal/core"
	"go.uber.org/zap"
)

type Handler struct {
	Logger      *zap.Logger
	Service     service.IService
	UserService userService.IService
}

func NewHandler(deps *core.Dependency) *Handler {
	repo := repository.NewRepository(deps.DB)
	service := service.NewService(repo, deps.Config.JwtPrivateKey)
	userService := userService.NewService(repo)

	return &Handler{
		Logger:      deps.Logger,
		Service:     service,
		UserService: userService,
	}
}
