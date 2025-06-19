package jwt

import (
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/pkg/jwt/types"
	"github.com/radityacandra/go-cms/pkg/util"
)

func Authorize() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			data, err := AuthorizeToken(c.Request().Header.Get("Authorization"))
			if err != nil {
				return util.ReturnError(c, err, nil)
			}

			c.Set(types.CONTEXT_KEY, data)

			return next(c)
		}
	}
}
