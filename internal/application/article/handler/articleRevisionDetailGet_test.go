package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/api/articlePrivate"
	"github.com/radityacandra/go-cms/internal/application/article/handler"
	"github.com/radityacandra/go-cms/internal/application/article/service"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	mockService "github.com/radityacandra/go-cms/mocks/internal_/application/article/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestHandler_ArticleRevisionDetailGet(t *testing.T) {
	type fields struct {
		Service service.IService
	}
	type args struct {
		articleId  articlePrivate.RequiredArticleIdParams
		revisionId articlePrivate.RequiredRevisionIdParams
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
			name: "should return error when revision not found",
			args: args{
				articleId:  "some-article-id",
				revisionId: "not-found-revision-id",
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

				service.EXPECT().DetailArticleRevision(mock.Anything, tt.args.articleId, tt.args.revisionId).
					Return(types.DetailArticleRevisionOutput{}, types.ErrArticleNotFound).Once()

				return tt
			},
		},
		{
			name: "should return article revision detail successfully",
			args: args{
				articleId:  "some-article-id",
				revisionId: "some-revision-id",
			},
			expected: expected{
				statusCode: http.StatusOK,
				response: articlePrivate.ArticleRevisionDetailGetResponse{
					Author: articlePrivate.AuthorSchema{
						Id:   "author-id",
						Name: "John Doe",
					},
					Content: "some content",
					Id:      "some-revision-id",
					Title:   "some title",
				},
			},
			fields: fields{},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().DetailArticleRevision(mock.Anything, tt.args.articleId, tt.args.revisionId).
					Return(types.DetailArticleRevisionOutput{
						Author: api.AuthorSchema{
							Id:   "author-id",
							Name: "John Doe",
						},
						Content: "some content",
						Id:      "some-revision-id",
						Title:   "some title",
					}, nil).Once()

				return tt
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/articles/:article_id/revisions/:revision_id", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			h := &handler.Handler{
				Service: tt.fields.Service,
				Logger:  zap.NewNop(),
			}
			err := h.ArticleRevisionDetailGet(c, tt.args.articleId, tt.args.revisionId)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.statusCode, rec.Code)

			body, _ := io.ReadAll(rec.Result().Body)
			if tt.expected.statusCode != http.StatusOK {
				var bodyStruct api.DefaultErrorResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.response, bodyStruct)
			} else {
				var bodyStruct articlePrivate.ArticleRevisionDetailGetResponse
				json.Unmarshal(body, &bodyStruct)
				assert.Equal(t, tt.expected.response, bodyStruct)
			}
		})
	}
}
