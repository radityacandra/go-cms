package types

import "github.com/radityacandra/go-cms/api"

type ListTagInput struct {
	Page     int
	PageSize int
}

type ListTagOutput api.TagListGetResponse

type ListTagRepoItem struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type CreateTagInput struct {
	api.TagCreatePostRequest
	UserId string
}
