package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api/article"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) ArticleDetailGet(ctx echo.Context, articleId article.RequiredArticleIdParams) error {
	userId := util.GetLoggedUser(ctx)
	reqCtx := ctx.Request().Context()
	output, err := h.Service.DetailArticle(reqCtx, articleId, userId)
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, output)
}
