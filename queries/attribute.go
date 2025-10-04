package queries

const (
	Attributes = `
		SELECT
			product_attributes.id AS product_attribute_id,
			product_attributes.product_id AS product_attribute_product_id,
			product_attributes.attribute_id AS product_attribute_attribute_id,
			product_attributes.title AS product_attribute_value,
			attributes.title AS attribute_title
		FROM product_attributes
		LEFT JOIN
			attributes ON product_attributes.attribute_id = attributes.id
		WHERE product_attributes.product_id=$1
		ORDER BY attributes.created;
	`
)
