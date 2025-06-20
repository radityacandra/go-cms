package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	"github.com/radityacandra/go-cms/pkg/jwt"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) ArticleUpdatePut(ctx echo.Context, articleId api.RequiredArticleIdParams) error {
	var reqBody api.ArticleUpdatePutRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if reqBody.Status != nil && *reqBody.Status == "published" {
		if ok := jwt.ScopeCheck(ctx, []string{"update-article-published"}); !ok {
			return nil
		}
	}

	reqCtx := ctx.Request().Context()
	userId := util.GetLoggedUser(ctx)
	revId, err := h.Service.CreateArticleRevision(reqCtx, types.CreateArticleRevisionInput{
		ArticleUpdatePutRequest: reqBody,
		ArticleId:               articleId,
		UserId:                  userId,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.IDOnlyResponseSchema{
		Id: revId,
	})
}
