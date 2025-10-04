package models

import "time"

type Rates struct {
	ID        *int       `json:"id" db:"rate_id"`
	UserName  *string    `json:"username" db:"username"`
	Title     *string    `json:"title" db:"rate_title"`
	Body      *string    `json:"body" db:"rate_body"`
	Point     *int       `json:"point" db:"rate_point"`
	Created   *time.Time `json:"created" db:"rate_created"`
	ProductID *int       `json:"product_id" db:"rate_product_id"`
}
