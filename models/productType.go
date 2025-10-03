package models

type ProductType struct {
	ID         int    `json:"id" db:"id"`
	CategoryID int    `json:"category_id" db:"category_id"`
	Title      string `json:"title" db:"title"`
	TitleFarsi string `json:"title_farsi" db:"title_farsi"`
	Image      string `json:"image" db:"image"`
}
