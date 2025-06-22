package types

import "github.com/radityacandra/go-cms/api"

type ListArticleInput struct {
	Page     int
	PageSize int
	Status   string
}

type ListArticleOutput struct {
	Data       []ListArticleItem
	Pagination Pagination
}

type ListArticleItem struct {
	Id         string `db:"id"`
	Title      string `db:"title"`
	Content    string `db:"content"`
	AuthorId   string `db:"author_id"`
	AuthorName string `db:"author_name"`
	Status     string `db:"status"`
}

type Pagination struct {
	Page      int
	PageSize  int
	TotalData int64
}

type CreateArticleInput struct {
	api.ArticleCreatePostRequest
	UserId string
}

type DetailArticleOutput api.ArticleDetailGetResponse

type CreateArticleRevisionInput struct {
	api.ArticleUpdatePutRequest
	ArticleId string
	UserId    string
}

type DetailArticleRevisionOutput api.ArticleRevisionDetailGetResponse

type UpsertTagAssociationInput struct {
	Id        string  `db:"id"`
	Tag1Id    string  `db:"tag1_id"`
	Tag2Id    string  `db:"tag2_id"`
	Score     float64 `db:"score"`
	CreatedBy string  `db:"created_by"`
	CreatedAt int64   `db:"created_at"`
}
