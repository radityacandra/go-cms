package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radityacandra/go-cms/api/article"
	"github.com/radityacandra/go-cms/api/articlePrivate"
	"github.com/radityacandra/go-cms/api/auth"
	"github.com/radityacandra/go-cms/api/authPrivate"
	"github.com/radityacandra/go-cms/api/tag"
	"github.com/radityacandra/go-cms/api/user"
	articleHandler "github.com/radityacandra/go-cms/internal/application/article/handler"
	authHandler "github.com/radityacandra/go-cms/internal/application/auth/handler"
	tagHandler "github.com/radityacandra/go-cms/internal/application/tag/handler"
	"github.com/radityacandra/go-cms/internal/application/user/handler"
	"github.com/radityacandra/go-cms/internal/core"
	"github.com/radityacandra/go-cms/pkg/jwt"
	"github.com/radityacandra/go-cms/pkg/validator"

	"go.uber.org/zap"
)

func InitServer(ctx context.Context, deps *core.Dependency) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Validator = validator.NewValidator()

	e.Use(middleware.CORS())
	e.Use(jwt.OptionalAuthorize())

	deps.Echo = e

	ePrivate := e.Group("")
	ePrivate.Use(jwt.Authorize())

	userHandler := handler.NewHandler(deps)
	user.RegisterHandlers(e, userHandler)

	articleHandler := articleHandler.NewHandler(deps)
	article.RegisterHandlers(e, articleHandler)
	articlePrivate.RegisterHandlers(ePrivate, articleHandler)

	authHandler := authHandler.NewHandler(deps)
	auth.RegisterHandlers(e, authHandler)
	authPrivate.RegisterHandlers(ePrivate, authHandler)

	tagHandler := tagHandler.NewHandler(deps)
	tag.RegisterHandlers(ePrivate, tagHandler)

	deps.Logger.Info("Web server ready", zap.Int("port", 9000))
	go func() {
		if err := e.Start(":9000"); err != nil && err != http.ErrServerClosed {
			deps.Logger.Fatal("Failed to start web server", zap.Error(err))
		}
	}()
}
