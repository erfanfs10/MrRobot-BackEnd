package models

type Attributes struct {
	ID             *int    `json:"id" db:"product_attribute_id"`
	ProductID      *int    `json:"product_id" db:"product_attribute_product_id"`
	AttributeID    *int    `json:"attribute_id" db:"product_attribute_attribute_id"`
	Value          *string `json:"value" db:"product_attribute_value"`
	AttributeTitle *string `json:"attribute_title" db:"attribute_title"`
}
