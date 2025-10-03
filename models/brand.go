package models

type Brand struct {
	ID         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	TitleFarsi string `json:"title_farsi" db:"title_farsi"`
	Image      string `json:"image" db:"image"`
}
