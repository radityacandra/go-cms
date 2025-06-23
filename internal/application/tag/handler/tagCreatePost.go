package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/tag/types"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) TagCreatePost(ctx echo.Context) error {
	var reqBody api.TagCreatePostRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if err := ctx.Validate(reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	userId := util.GetLoggedUser(ctx)
	reqCtx := ctx.Request().Context()
	tagId, err := h.Service.CreateTag(reqCtx, types.CreateTagInput{
		TagCreatePostRequest: reqBody,
		UserId:               userId,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.IDOnlyResponseSchema{
		Id: tagId,
	})
}
