package model

import (
	"math"
	"time"
)

type Tag struct {
	Id            string  `db:"id"`
	Name          string  `db:"name"`
	TrendingScore float64 `db:"trending_score"`
	UsageCount    int     `db:"usage_count"`
	CreatedBy     string  `db:"created_by"`
	CreatedAt     int64   `db:"created_at"`
	UpdatedBy     *string `db:"updated_by"`
	UpdatedAt     *int64  `db:"updated_at"`

	TagArticles []TagArticle `db:"-"`
}

type TagArticle struct {
	Id        string `db:"id"`
	TagId     string `db:"tag_id"`
	ArticleId string `db:"article_id"`
	CreatedAt int64  `db:"created_at"`
}

func NewTag(id, name, createdBy string, createdAt int64) *Tag {
	return &Tag{
		Id:        id,
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: createdAt,
	}
}

func (t *Tag) averageArticleTimeInHour() float64 {
	totalItem := len(t.TagArticles)
	now := time.Now().UnixMilli()
	totalTime := int64(0)
	for _, tagArticle := range t.TagArticles {
		totalTime += ((now - tagArticle.CreatedAt) / 1000) / 3600
	}

	return float64(totalTime) / float64(totalItem)
}

func (t *Tag) CalculateTrendingScore() float64 {
	totalItem := len(t.TagArticles)
	avgArticleLife := t.averageArticleTimeInHour()

	return float64(totalItem) / math.Pow(avgArticleLife+1, 1.5)
}

func (t *Tag) UpdateTrendingScore(score float64, updatedBy string) {
	now := time.Now().UnixMilli()
	t.TrendingScore = score
	t.UpdatedAt = &now
	t.UpdatedBy = &updatedBy
}
