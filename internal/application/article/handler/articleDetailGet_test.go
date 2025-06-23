package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/api/article"
	"github.com/radityacandra/go-cms/internal/application/article/handler"
	"github.com/radityacandra/go-cms/internal/application/article/service"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	mockService "github.com/radityacandra/go-cms/mocks/internal_/application/article/service"
	jwtType "github.com/radityacandra/go-cms/pkg/jwt/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestHandler_ArticleDetailGet(t *testing.T) {
	type fields struct {
		Service service.IService
	}
	type args struct {
		articleId  article.RequiredArticleIdParams
		withoutJwt bool
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

	userId := uuid.NewString()

	tests := []test{
		{
			name: "should return error when article not found",
			args: args{
				articleId: "non-existent-article",
			},
			expected: expected{
				statusCode: http.StatusNotFound,
				response: api.DefaultErrorResponse{
					Error: types.ErrArticleNotFound.Error(),
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().DetailArticle(mock.Anything, tt.args.articleId, userId).
					Return(types.DetailArticleOutput{}, types.ErrArticleNotFound).Once()

				return tt
			},
		},
		{
			name: "should return article detail successfully",
			args: args{
				articleId: "article-123",
			},
			expected: expected{
				statusCode: http.StatusOK,
				response: article.ArticleDetailGetResponse{
					Author: article.AuthorSchema{
						Id:   "some-author-id",
						Name: "John Doe",
					},
					Content: "some content",
					Id:      "article-123",
					RevisionHistories: []article.RevisionHistorySchema{{
						Id:           "some-revision-history-id",
						AuthorId:     "some-revision-author-id",
						AuthorName:   "John Smith",
						RevisionTime: 1,
					}},
					Title: "some title",
					Tags: api.ArticleTagsSchema{
						"Art",
					},
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().DetailArticle(mock.Anything, tt.args.articleId, userId).
					Return(types.DetailArticleOutput{
						Author: api.AuthorSchema{
							Id:   "some-author-id",
							Name: "John Doe",
						},
						Content: "some content",
						Id:      "article-123",
						RevisionHistories: []api.RevisionHistorySchema{{
							Id:           "some-revision-history-id",
							AuthorId:     "some-revision-author-id",
							AuthorName:   "John Smith",
							RevisionTime: 1,
						}},
						Title: "some title",
						Tags: api.ArticleTagsSchema{
							"Art",
						},
					}, nil).Once()

				return tt
			},
		},
		{
			name: "should return article detail with no jwt context",
			args: args{
				articleId:  "article-123",
				withoutJwt: true,
			},
			expected: expected{
				statusCode: http.StatusOK,
				response: article.ArticleDetailGetResponse{
					Author: article.AuthorSchema{
						Id:   "some-author-id",
						Name: "John Doe",
					},
					Content: "some content",
					Id:      "article-123",
					RevisionHistories: []article.RevisionHistorySchema{{
						Id:           "some-revision-history-id",
						AuthorId:     "some-revision-author-id",
						AuthorName:   "John Smith",
						RevisionTime: 1,
					}},
					Title: "some title",
					Tags: api.ArticleTagsSchema{
						"Art",
					},
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().DetailArticle(mock.Anything, tt.args.articleId, "").
					Return(types.DetailArticleOutput{
						Author: api.AuthorSchema{
							Id:   "some-author-id",
							Name: "John Doe",
						},
						Content: "some content",
						Id:      "article-123",
						RevisionHistories: []api.RevisionHistorySchema{{
							Id:           "some-revision-history-id",
							AuthorId:     "some-revision-author-id",
							AuthorName:   "John Smith",
							RevisionTime: 1,
						}},
						Title: "some title",
						Tags: api.ArticleTagsSchema{
							"Art",
						},
					}, nil).Once()

				return tt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/articles/"+string(tt.args.articleId), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Set JWT context if user is expected
			if !tt.args.withoutJwt {
				c.Set(jwtType.CONTEXT_KEY, map[string]interface{}{
					"sub": userId,
				})
			}

			h := &handler.Handler{
				Service: tt.fields.Service,
				Logger:  zap.NewNop(),
			}

			err := h.ArticleDetailGet(c, tt.args.articleId)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected.statusCode, rec.Code)

			body, _ := io.ReadAll(rec.Result().Body)
			if tt.expected.statusCode != http.StatusOK {
				var bodyStruct api.DefaultErrorResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.response, bodyStruct)
			} else {
				var bodyStruct article.ArticleDetailGetResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.response, bodyStruct)
			}
		})
	}
}
