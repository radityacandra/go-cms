package handler_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestHandler_ArticleUpdatePut(t *testing.T) {
	type fields struct {
		Service service.IService
	}
	type args struct {
		articleId  article.RequiredArticleIdParams
		reqBody    api.ArticleUpdatePutRequest
		reqBodyStr string
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
			name: "should return error if request body is invalid",
			args: args{
				articleId:  "some-article-id",
				reqBodyStr: "{\"tags\": true}",
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				response: api.DefaultErrorResponse{
					Error: "code=400, message=Unmarshal type error: expected=[]string, got=bool, field=tags, offset=13, internal=json: cannot unmarshal bool into Go struct field ArticleUpdatePutRequest.tags of type []string",
				},
			},
			mock: func(tt test) test { return tt },
		},
		{
			name: "should return error if validation fails",
			args: args{
				articleId: "some-article-id",
				reqBody: api.ArticleUpdatePutRequest{
					Status: func() *string { status := "any"; return &status }(),
				},
				jwtContext: map[string]interface{}{"sub": "user-123"},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				response: api.DefaultErrorResponse{
					Error: "invalid allowed value for status",
				},
			},
			mock: func(tt test) test { return tt },
		},
		{
			name: "should return error if service returns error",
			args: args{
				articleId:  "some-article-id",
				reqBody:    api.ArticleUpdatePutRequest{},
				jwtContext: map[string]interface{}{"sub": userId},
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

				service.EXPECT().CreateArticleRevision(mock.Anything, types.CreateArticleRevisionInput{
					ArticleUpdatePutRequest: api.ArticleUpdatePutRequest{},
					ArticleId:               tt.args.articleId,
					UserId:                  userId,
				}).Return("", errors.New("some error")).Once()

				return tt
			},
		},
		{
			name: "should update article successfully",
			args: args{
				articleId: "some-article-id",
				reqBody: api.ArticleUpdatePutRequest{
					Title: func() *string { title := "some title"; return &title }(),
				},
				jwtContext: map[string]interface{}{"sub": userId},
			},
			expected: expected{
				statusCode: http.StatusOK,
				response: api.IDOnlyResponseSchema{
					Id: "generated-id",
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().CreateArticleRevision(mock.Anything, types.CreateArticleRevisionInput{
					ArticleUpdatePutRequest: api.ArticleUpdatePutRequest{
						Title: tt.args.reqBody.Title,
					},
					ArticleId: tt.args.articleId,
					UserId:    userId,
				}).Return("generated-id", nil).Once()

				return tt
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			var reqBody string
			if tt.args.reqBodyStr != "" {
				reqBody = tt.args.reqBodyStr
			} else {
				bytes, _ := json.Marshal(tt.args.reqBody)
				reqBody = string(bytes)
			}

			e := echo.New()
			e.Validator = validator.NewValidator()
			req := httptest.NewRequest(http.MethodPut, "/api/v1/articles/:article_id", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if tt.args.jwtContext != nil {
				c.Set(jwtType.CONTEXT_KEY, tt.args.jwtContext)
			}

			h := &handler.Handler{
				Service: tt.fields.Service,
				Logger:  zap.NewNop(),
			}
			err := h.ArticleUpdatePut(c, tt.args.articleId)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.statusCode, rec.Code)

			body, _ := io.ReadAll(rec.Result().Body)
			if tt.expected.statusCode != http.StatusOK {
				var bodyStruct api.DefaultErrorResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.response, bodyStruct)
			} else {
				var bodyStruct api.IDOnlyResponseSchema
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.response, bodyStruct)
			}
		})
	}
}
