package dto

type ResourceDTO struct {
	Id          *[]byte `db:"id"`
	Title       *string `db:"title"`
	Ordering    *int    `db:"ordering"`
	Description *string `db:"description"`
	Link        *string `db:"link"`
}
