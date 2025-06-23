package jwt

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/pkg/jwt/types"
)

func ScopeCheck(ctx echo.Context, allowedScopes []string) bool {
	data := ctx.Get(types.CONTEXT_KEY).(map[string]interface{})
	if data["scopes"] == nil && len(allowedScopes) > 0 {
		ctx.JSON(http.StatusForbidden, api.DefaultErrorResponse{
			Error: "you don't have permission to access this resource",
		})

		return false
	}

	jwtScopes := make([]string, len(data["scopes"].([]interface{})))

	for i, scope := range data["scopes"].([]interface{}) {
		jwtScopes[i] = scope.(string)
	}

	for _, jwtScope := range jwtScopes {
		for _, allowedScope := range allowedScopes {
			if jwtScope == allowedScope {
				return true
			}
		}
	}

	ctx.JSON(http.StatusForbidden, api.DefaultErrorResponse{
		Error: "you don't have permission to access this resource",
	})

	return false
}
