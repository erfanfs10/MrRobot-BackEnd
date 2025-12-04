package models

type AttributeFilters struct {
	ID     int      `db:"id" json:"id"`
	Title  string   `db:"title" json:"title"`
	Values []string `json:"values"`
}

type AttributeFilterValue struct {
	AttributeID int    `db:"attribute_id"`
    Title       string `db:"title"`
}