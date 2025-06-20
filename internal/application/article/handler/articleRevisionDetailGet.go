package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) ArticleRevisionDetailGet(ctx echo.Context, articleId api.RequiredArticleIdParams, revisionId api.RequiredRevisionIdParams) error {
	reqCtx := ctx.Request().Context()
	output, err := h.Service.DetailArticleRevision(reqCtx, articleId, revisionId)
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, output)
}
