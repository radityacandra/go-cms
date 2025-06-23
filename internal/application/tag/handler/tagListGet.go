package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api/tag"
	"github.com/radityacandra/go-cms/internal/application/tag/types"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) TagListGet(ctx echo.Context, params tag.TagListGetParams) error {
	if err := ctx.Validate(params); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if params.Page == nil {
		defaultPage := 1
		params.Page = &defaultPage
	}

	if params.PageSize == nil {
		defaultPageSize := 10
		params.PageSize = &defaultPageSize
	}

	reqCtx := ctx.Request().Context()
	output, err := h.Service.ListTag(reqCtx, types.ListTagInput{
		Page:     *params.Page,
		PageSize: *params.PageSize,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, output)
}
