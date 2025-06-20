package model

import "time"

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

	Author           Author    `db:"-"`
	ArticleRevisions []Article `db:"-"`
}

type Author struct {
	Id   string
	Name string
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
