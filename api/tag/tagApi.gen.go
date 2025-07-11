// Package tag provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package tag

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	"github.com/radityacandra/go-cms/pkg/jwt"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// DefaultErrorResponse defines model for DefaultErrorResponse.
type DefaultErrorResponse struct {
	// Error error description
	Error string `json:"error"`
}

// IDOnlyResponseSchema defines model for IDOnlyResponseSchema.
type IDOnlyResponseSchema struct {
	// Id ID of created object
	Id string `json:"id"`
}

// PaginationSchema defines model for PaginationSchema.
type PaginationSchema struct {
	// Page current active page number (default to 1)
	Page int64 `json:"page"`

	// PageSize total data displayed on the page (default to 10)
	PageSize int64 `json:"pageSize"`

	// TotalData total number of data available. with formula of ceil(totalData / pageSize), can determine total page available
	TotalData int64 `json:"totalData"`
}

// TagCreatePostRequest defines model for TagCreatePostRequest.
type TagCreatePostRequest struct {
	// Name name of the tag
	Name string `json:"name" validate:"required"`
}

// TagListGetResponse defines model for TagListGetResponse.
type TagListGetResponse struct {
	Data       []TagListGetResponseItem `json:"data"`
	Pagination PaginationSchema         `json:"pagination"`
}

// TagListGetResponseItem defines model for TagListGetResponseItem.
type TagListGetResponseItem struct {
	// Id id of the tag
	Id string `json:"id"`

	// Name name of the tag
	Name string `json:"name"`
}

// OptionalPageParams defines model for OptionalPageParams.
type OptionalPageParams = int

// OptionalPageSizeParams defines model for OptionalPageSizeParams.
type OptionalPageSizeParams = int

// TagListGetParams defines parameters for TagListGet.
type TagListGetParams struct {
	// Page active page in pagination. default to 1
	Page *OptionalPageParams `form:"page,omitempty" json:"page,omitempty" validate:"omitempty,min=1"`

	// PageSize max number of data in the active page. default to 10
	PageSize *OptionalPageSizeParams `form:"page-size,omitempty" json:"page-size,omitempty" validate:"omitempty,min=1"`
}

// TagCreatePostJSONRequestBody defines body for TagCreatePost for application/json ContentType.
type TagCreatePostJSONRequestBody = TagCreatePostRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get tag list
	// (GET /api/v1/tags)
	TagListGet(ctx echo.Context, params TagListGetParams) error
	// Create a new tag
	// (POST /api/v1/tags)
	TagCreatePost(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// TagListGet converts echo context to params.
func (w *ServerInterfaceWrapper) TagListGet(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{"list-tag"})

	if ok := jwt.ScopeCheck(ctx, []string{"list-tag"}); !ok {
		return nil
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params TagListGetParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "page-size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page-size", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page-size: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.TagListGet(ctx, params)
	return err
}

// TagCreatePost converts echo context to params.
func (w *ServerInterfaceWrapper) TagCreatePost(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{"create-tag"})

	if ok := jwt.ScopeCheck(ctx, []string{"create-tag"}); !ok {
		return nil
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.TagCreatePost(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/v1/tags", wrapper.TagListGet)
	router.POST(baseURL+"/api/v1/tags", wrapper.TagCreatePost)

}
