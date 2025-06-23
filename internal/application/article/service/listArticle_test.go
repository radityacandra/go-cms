package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/radityacandra/go-cms/internal/application/article/repository"
	"github.com/radityacandra/go-cms/internal/application/article/service"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	repositoryMock "github.com/radityacandra/go-cms/mocks/internal_/application/article/repository"
	"github.com/stretchr/testify/assert"
)

func TestService_ListArticle(t *testing.T) {
	dummyError := errors.New("some error")

	type fields struct {
		Repository repository.IRepository
	}
	type args struct {
		ctx   context.Context
		input types.ListArticleInput
	}

	type expectation struct {
		output types.ListArticleOutput
		err    error
	}

	type test struct {
		name        string
		fields      fields
		args        args
		expectation expectation
		mock        func(tt test) test
	}

	tests := []test{
		{
			name: "should return data",
			args: args{
				ctx:   context.Background(),
				input: types.ListArticleInput{Page: 1, PageSize: 10, Status: "published"},
			},
			expectation: expectation{
				output: types.ListArticleOutput{
					Data: []types.ListArticleItem{
						{
							Id:         "article-1",
							Title:      "Test Article",
							Content:    "This is a test article.",
							AuthorId:   "author-1",
							AuthorName: "John Doe",
							Status:     "published",
						},
					},
					Pagination: types.Pagination{
						Page:      1,
						PageSize:  10,
						TotalData: 1,
					},
				},
			},
			mock: func(tt test) test {
				repo := repositoryMock.NewMockIRepository(t)
				tt.fields.Repository = repo

				repo.EXPECT().ListArticle(tt.args.ctx, tt.args.input).Return(types.ListArticleOutput{
					Data: []types.ListArticleItem{
						{
							Id:         "article-1",
							Title:      "Test Article",
							Content:    "This is a test article.",
							AuthorId:   "author-1",
							AuthorName: "John Doe",
							Status:     "published",
						},
					},
					Pagination: types.Pagination{
						Page:      1,
						PageSize:  10,
						TotalData: 1,
					},
				}, nil).Once()

				return tt
			},
		},
		{
			name: "should return error if repository return error",
			args: args{
				ctx:   context.Background(),
				input: types.ListArticleInput{Page: 1, PageSize: 10, Status: "published"},
			},
			expectation: expectation{
				err: dummyError,
			},
			mock: func(tt test) test {
				repo := repositoryMock.NewMockIRepository(t)
				tt.fields.Repository = repo

				repo.EXPECT().ListArticle(tt.args.ctx, tt.args.input).
					Return(types.ListArticleOutput{}, dummyError).Once()

				return tt
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt.mock(tt)
			s := &service.Service{
				Repository: tt.fields.Repository,
			}

			got, err := s.ListArticle(tt.args.ctx, tt.args.input)
			assert.ErrorIs(t, err, tt.expectation.err)
			assert.Equal(t, tt.expectation.output, got)
		})
	}
}
