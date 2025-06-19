package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/auth/types"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) AuthLoginPost(ctx echo.Context) error {
	var reqBody api.AuthLoginRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if err := ctx.Validate(reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	reqCtx := ctx.Request().Context()
	output, err := h.Service.Login(reqCtx, types.LoginInput{
		Username: reqBody.Username,
		Password: reqBody.Password,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.AuthLoginResponse{
		Token: output.Token,
	})
}
