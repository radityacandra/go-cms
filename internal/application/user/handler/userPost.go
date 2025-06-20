package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/user/types"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) UserPost(ctx echo.Context) error {
	var reqBody api.UserPostRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if err := ctx.Validate(reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	reqCtx := ctx.Request().Context()

	output, err := h.Service.RegisterUser(reqCtx, types.RegisterUserInput{
		Username: reqBody.Username,
		Password: reqBody.Password,
		FullName: reqBody.FullName,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.UserPostResponse{
		Id: output.Id,
	})
}
