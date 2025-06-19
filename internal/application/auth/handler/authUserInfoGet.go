package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/pkg/util"
)

func (h *Handler) AuthUserInfoGet(ctx echo.Context) error {
	userId := util.GetLoggedUser(ctx)
	h.Logger.Info(userId)

	reqCtx := ctx.Request().Context()
	output, err := h.UserService.DetailUser(reqCtx, userId)
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.AuthUserInfoGetResponse{
		Username: output.Username,
		Scopes:   output.Acls,
	})
}
