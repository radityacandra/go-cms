package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	"github.com/radityacandra/go-cms/pkg/jwt"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) ArticleCreatePost(ctx echo.Context) error {
	var reqBody api.ArticleCreatePostRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if reqBody.Status == "published" {
		if ok := jwt.ScopeCheck(ctx, []string{"create-article-published"}); !ok {
			return nil
		}
	}

	userId := util.GetLoggedUser(ctx)
	reqCtx := ctx.Request().Context()
	id, err := h.Service.CreateArticle(reqCtx, types.CreateArticleInput{
		ArticleCreatePostRequest: reqBody,
		UserId:                   userId,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.IDOnlyResponseSchema{
		Id: id,
	})
}
