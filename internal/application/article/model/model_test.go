package model_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/go-cms/internal/application/article/model"
	"github.com/stretchr/testify/assert"
)

func TestArticle_DeactivateArticle(t *testing.T) {
	t.Run("should deactivate current article", func(t *testing.T) {
		mainArticleId := uuid.NewString()
		userId := uuid.NewString()

		a := &model.Article{
			Id:       uuid.NewString(),
			ParentId: nil,
		}
		a.DeactivateArticle(mainArticleId, userId)

		assert.NotNil(t, a.ParentId)
		assert.Equal(t, mainArticleId, *a.ParentId)
		assert.NotNil(t, a.UpdatedAt)
		assert.True(t, time.Now().UnixMilli() >= *a.UpdatedAt)
		assert.Equal(t, userId, *a.UpdatedBy)
	})
}

func TestArticle_TagCombinationPairs(t *testing.T) {
	tag1 := uuid.NewString()
	tag2 := uuid.NewString()
	tag3 := uuid.NewString()

	// tag1, tag2, tag3 := uuid.NewString(), uuid.NewString(), uuid.NewString()
	a := &model.Article{
		Id:       uuid.NewString(),
		ParentId: nil,
		Tags: []model.ArticleTag{
			{
				Id:    uuid.NewString(),
				TagId: tag2,
			},
			{
				Id:    uuid.NewString(),
				TagId: tag1,
			},
			{
				Id:    uuid.NewString(),
				TagId: tag3,
			},
		},
	}

	pairsList := a.TagCombinationPairs()
	assert.Equal(t, 3, len(pairsList))

	// pair 1
	assert.Equal(t, tag1, pairsList[0][0])
	assert.Equal(t, tag2, pairsList[0][1])
}
