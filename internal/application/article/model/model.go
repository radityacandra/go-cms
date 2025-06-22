package model

import (
	"slices"
	"time"

	tagModel "github.com/radityacandra/go-cms/internal/application/tag/model"
)

type Article struct {
	Id        string  `db:"id"`
	Content   string  `db:"content"`
	Title     string  `db:"title"`
	AuthorId  string  `db:"author_id"`
	ParentId  *string `db:"parent_id"`
	Status    string  `db:"status"`
	CreatedBy string  `db:"created_by"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt *int64  `db:"updated_at"`
	UpdatedBy *string `db:"updated_by"`

	Author           Author       `db:"-"`
	ArticleRevisions []Article    `db:"-"`
	Tags             []ArticleTag `db:"-"`
}

type Author struct {
	Id   string
	Name string
}

type ArticleTag struct {
	Id        string  `db:"id"`
	ArticleId string  `db:"article_id"`
	TagId     string  `db:"tag_id"`
	CreatedBy string  `db:"created_by"`
	CreatedAt int64   `db:"created_at"`
	UpdatedAt *int64  `db:"updated_at"`
	UpdatedBy *string `db:"updated_by"`

	Tag tagModel.Tag `db:"-"`
}

func NewArticle(id, content, title, authorId, status, createdBy string, createdAt int64) *Article {
	return &Article{
		Id:        id,
		Content:   content,
		Title:     title,
		AuthorId:  authorId,
		Status:    status,
		CreatedBy: createdBy,
		CreatedAt: createdAt,
	}
}

func (a *Article) DeactivateArticle(mainArticleId, userId string) {
	a.ParentId = &mainArticleId

	now := time.Now().UnixMilli()
	a.UpdatedAt = &now
	a.UpdatedBy = &userId
}

func (a *Article) TagCombinationPairs() [][]string {
	pairList := [][]string{}

	for index, tag := range a.Tags {
		if index+1 >= len(a.Tags) {
			break
		}

		for _, tagPair := range a.Tags[index+1:] {
			pair := []string{tag.TagId, tagPair.TagId}
			slices.Sort(pair)

			pairList = append(pairList, pair)
		}
	}

	return pairList
}
