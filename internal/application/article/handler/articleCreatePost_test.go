package handler_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
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

var userId = uuid.NewString()

func TestArticleCreatePost(t *testing.T) {
	type fields struct {
		Service service.IService
		Logger  *zap.Logger
	}
	type args struct {
		reqBody    api.ArticleCreatePostRequest
		reqBodyStr string
		jwtScope   []interface{}
	}

	type expected struct {
		statusCode int
		body       interface{}
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
			name: "should return error if failed to bind request body",
			args: args{
				reqBodyStr: "{\"tags\": true}",
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Error: "code=400, message=Unmarshal type error: expected=[]string, got=bool, field=tags, offset=13, internal=json: cannot unmarshal bool into Go struct field ArticleCreatePostRequest.tags of type []string",
				},
			},
			fields: fields{
				Logger: zap.NewNop(),
			},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return error if validation fails",
			args: args{
				reqBody: api.ArticleCreatePostRequest{
					Content: "some content",
					Status:  "drafted",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Error: "title must not be empty.",
				},
			},
			fields: fields{
				Logger: zap.NewNop(),
			},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return error if try to publish article without relevant access",
			args: args{
				reqBody: api.ArticleCreatePostRequest{
					Content: "some content",
					Status:  "published",
					Title:   "some title",
				},
			},
			expected: expected{
				statusCode: http.StatusForbidden,
				body: api.DefaultErrorResponse{
					Error: "you don't have permission to access this resource",
				},
			},
			fields: fields{
				Logger: zap.NewNop(),
			},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return error if service returns error",
			args: args{
				reqBody: api.ArticleCreatePostRequest{
					Content: "some content",
					Title:   "some title",
					Status:  "published",
				},
				jwtScope: []interface{}{"create-article-published"},
			},
			expected: expected{
				statusCode: http.StatusInternalServerError,
				body: api.DefaultErrorResponse{
					Error: "unknown error",
				},
			},
			fields: fields{
				Logger: zap.NewNop(),
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().CreateArticle(mock.Anything, types.CreateArticleInput{
					ArticleCreatePostRequest: tt.args.reqBody,
					UserId:                   userId,
				}).Return("", errors.New("some error")).Once()

				return tt
			},
		},
		{
			name: "should create article successfully",
			args: args{
				reqBody: api.ArticleCreatePostRequest{
					Content: "some content",
					Title:   "some title",
					Status:  "drafted",
				},
			},
			expected: expected{
				statusCode: http.StatusOK,
				body: api.IDOnlyResponseSchema{
					Id: uuid.NewString(),
				},
			},
			fields: fields{
				Logger: zap.NewNop(),
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().CreateArticle(mock.Anything, types.CreateArticleInput{
					ArticleCreatePostRequest: tt.args.reqBody,
					UserId:                   userId,
				}).Return(tt.expected.body.(api.IDOnlyResponseSchema).Id, nil).Once()

				return tt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			e := echo.New()
			e.Validator = validator.NewValidator()

			var reqBody string
			if tt.args.reqBodyStr != "" {
				reqBody = tt.args.reqBodyStr
			} else {
				bytes, _ := json.Marshal(tt.args.reqBody)
				reqBody = string(bytes)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/articles", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set(jwtType.CONTEXT_KEY, map[string]interface{}{
				"sub":    userId,
				"scopes": tt.args.jwtScope,
			})

			h := &handler.Handler{
				Service: tt.fields.Service,
				Logger:  tt.fields.Logger,
			}
			err := h.ArticleCreatePost(c)

			assert.NoError(t, err)

			assert.Equal(t, tt.expected.statusCode, rec.Code)

			body, _ := io.ReadAll(rec.Result().Body)
			if tt.expected.statusCode != http.StatusOK {
				var bodyStruct api.DefaultErrorResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.body, bodyStruct)
			} else {
				var bodyStruct api.IDOnlyResponseSchema
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.body, bodyStruct)
			}
		})
	}
}
