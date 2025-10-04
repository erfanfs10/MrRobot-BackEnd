package models

type ProductImages struct {
	ID        *int    `json:"id" db:"product_image_id"`
	Image     *string `json:"image" db:"product_image_image"`
	IsPrimary *bool   `json:"is_primary" db:"product_image_is_primary"`
	ProductID *int    `json:"product_id" db:"product_image_product_id"`
}
