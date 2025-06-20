package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/api/article"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) ArticleListGet(ctx echo.Context, params article.ArticleListGetParams) error {
	userId := util.GetLoggedUser(ctx)
	if params.Status != nil && userId == "" && *params.Status != "published" {
		return util.ReturnBadRequest(ctx, errors.New("unauthenticated user cannot see unpublished data"), h.Logger)
	}

	input := types.ListArticleInput{}
	if params.Status == nil && userId != "" {
		input.Status = "all"
	} else if params.Status == nil && userId == "" {
		input.Status = "published"
	} else {
		input.Status = *params.Status
	}

	input.Page = 1
	if params.Page != nil {
		input.Page = *params.Page
	}

	input.PageSize = 10
	if params.PageSize != nil {
		input.PageSize = *params.PageSize
	}

	reqCtx := ctx.Request().Context()
	output, err := h.Service.ListArticle(reqCtx, input)
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	response := api.ArticleListGetResponse{
		Pagination: api.PaginationSchema{
			Page:      int64(output.Pagination.Page),
			PageSize:  int64(output.Pagination.PageSize),
			TotalData: output.Pagination.TotalData,
		},
		Data: []api.ArticleListGetResponseItem{},
	}

	for _, item := range output.Data {
		response.Data = append(response.Data, api.ArticleListGetResponseItem{
			Id:      item.Id,
			Title:   item.Title,
			Content: item.Content,
			Status:  item.Status,
			Author: api.AuthorSchema{
				Id:   item.AuthorId,
				Name: item.AuthorName,
			},
		})
	}

	return ctx.JSON(http.StatusOK, response)
}
