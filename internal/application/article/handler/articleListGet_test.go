package handler_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/api/article"
	"github.com/radityacandra/go-cms/internal/application/article/handler"
	"github.com/radityacandra/go-cms/internal/application/article/service"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	mockService "github.com/radityacandra/go-cms/mocks/internal_/application/article/service"
	jwtType "github.com/radityacandra/go-cms/pkg/jwt/types"
	"github.com/radityacandra/go-cms/pkg/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestHandler_ArticleListGet(t *testing.T) {
	type fields struct {
		Service service.IService
	}
	type args struct {
		params     article.ArticleListGetParams
		jwtContext map[string]interface{}
	}

	type expected struct {
		statusCode int
		response   interface{}
	}

	type test struct {
		name     string
		fields   fields
		args     args
		expected expected
		mock     func(tt test) test
	}

	tests := []test{
		{
			name: "should return article list successfully with default parameters",
			args: args{
				params: article.ArticleListGetParams{},
				jwtContext: map[string]interface{}{
					"sub": userId,
				},
			},
			expected: expected{
				statusCode: http.StatusOK,
				response: api.ArticleListGetResponse{
					Data:       []api.ArticleListGetResponseItem{},
					Pagination: api.PaginationSchema{},
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().ListArticle(mock.Anything, types.ListArticleInput{
					Page:     1,
					PageSize: 10,
					Status:   "all",
				}).Return(types.ListArticleOutput{
					Data:       []types.ListArticleItem{},
					Pagination: types.Pagination{},
				}, nil)

				return tt
			},
		},
		{
			name: "should return error when validation fails for page parameter",
			args: args{
				params: article.ArticleListGetParams{
					Page: func() *int { p := 0; return &p }(),
				},
				jwtContext: map[string]interface{}{
					"sub": "user-123",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				response: api.DefaultErrorResponse{
					Error: "minimum value for page is 1",
				},
			},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return error when validation fails for page-size parameter",
			args: args{
				params: article.ArticleListGetParams{
					PageSize: func() *int { ps := 0; return &ps }(),
				},
				jwtContext: map[string]interface{}{
					"sub": "user-123",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				response: api.DefaultErrorResponse{
					Error: "minimum value for page-size is 1",
				},
			},
			fields: fields{},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return error when validation fails for status parameter",
			args: args{
				params: article.ArticleListGetParams{
					Status: func() *string { s := "invalid-status"; return &s }(),
				},
				jwtContext: map[string]interface{}{
					"sub": "user-123",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				response: api.DefaultErrorResponse{
					Error: "invalid allowed value for status",
				},
			},
			fields: fields{},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return article list with custom pagination",
			args: args{
				params: article.ArticleListGetParams{
					Page:     func() *int { p := 2; return &p }(),
					PageSize: func() *int { ps := 5; return &ps }(),
				},
				jwtContext: map[string]interface{}{
					"sub": "user-123",
				},
			},
			expected: expected{
				statusCode: http.StatusOK,
				response: api.ArticleListGetResponse{
					Data:       []api.ArticleListGetResponseItem{},
					Pagination: api.PaginationSchema{},
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().ListArticle(mock.Anything, types.ListArticleInput{
					Page:     2,
					PageSize: 5,
					Status:   "all",
				}).Return(types.ListArticleOutput{
					Data:       []types.ListArticleItem{},
					Pagination: types.Pagination{},
				}, nil)

				return tt
			},
		},
		{
			name: "should return error when unauthenticated user requests non-published status",
			args: args{
				params: article.ArticleListGetParams{
					Status: func() *string { s := "drafted"; return &s }(),
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				response: api.DefaultErrorResponse{
					Error: "unauthenticated user cannot see unpublished data",
				},
			},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return published articles for unauthenticated user",
			args: args{
				params: article.ArticleListGetParams{},
			},
			expected: expected{
				statusCode: http.StatusOK,
				response: api.ArticleListGetResponse{
					Data:       []api.ArticleListGetResponseItem{},
					Pagination: api.PaginationSchema{},
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().ListArticle(mock.Anything, types.ListArticleInput{
					Page:     1,
					PageSize: 10,
					Status:   "published",
				}).Return(types.ListArticleOutput{
					Data:       []types.ListArticleItem{},
					Pagination: types.Pagination{},
				}, nil)

				return tt
			},
		},
		{
			name: "should return article list filtered by status",
			args: args{
				params: article.ArticleListGetParams{
					Status: func() *string { s := "drafted"; return &s }(),
				},
				jwtContext: map[string]interface{}{
					"sub": "user-123",
				},
			},
			expected: expected{
				statusCode: http.StatusOK,
				response: api.ArticleListGetResponse{
					Data:       []api.ArticleListGetResponseItem{},
					Pagination: api.PaginationSchema{},
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().ListArticle(mock.Anything, types.ListArticleInput{
					Page:     1,
					PageSize: 10,
					Status:   "drafted",
				}).Return(types.ListArticleOutput{
					Data:       []types.ListArticleItem{},
					Pagination: types.Pagination{},
				}, nil)

				return tt
			},
		},
		{
			name: "should return error when service returns error",
			args: args{
				params: article.ArticleListGetParams{},
				jwtContext: map[string]interface{}{
					"sub": "user-123",
				},
			},
			expected: expected{
				statusCode: http.StatusInternalServerError,
				response: api.DefaultErrorResponse{
					Error: "unknown error",
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().ListArticle(mock.Anything, types.ListArticleInput{
					Page:     1,
					PageSize: 10,
					Status:   "all",
				}).Return(types.ListArticleOutput{}, errors.New("some error"))

				return tt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			e := echo.New()
			e.Validator = validator.NewValidator()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/articles", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if tt.args.jwtContext != nil {
				c.Set(jwtType.CONTEXT_KEY, tt.args.jwtContext)
			}

			h := &handler.Handler{
				Service: tt.fields.Service,
				Logger:  zap.NewNop(),
			}
			err := h.ArticleListGet(c, tt.args.params)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.statusCode, rec.Code)

			body, _ := io.ReadAll(rec.Result().Body)
			if tt.expected.statusCode != http.StatusOK {
				var bodyStruct api.DefaultErrorResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.response, bodyStruct)
			} else {
				var bodyStruct api.ArticleListGetResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.response, bodyStruct)
			}
		})
	}
}
