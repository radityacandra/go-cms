package model

type Tag struct {
	Id              string  `db:"id"`
	Name            string  `db:"name"`
	PopularityScore float64 `db:"popularity_score"`
	CreatedBy       string  `db:"created_by"`
	CreatedAt       int64   `db:"created_at"`
	UpdatedBy       *string `db:"updated_by"`
	UpdatedAt       *int64  `db:"updated_at"`
}

func NewTag(id, name, createdBy string, createdAt int64) *Tag {
	return &Tag{
		Id:        id,
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: createdAt,
	}
}
