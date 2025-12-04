package queries

const (
	GetFilterAttributes = `
	SELECT a.id, a.title
		FROM attributes a
		JOIN filters f ON f.attribute_id = a.id
		JOIN product_types pt ON f.product_type_id = pt.id
	WHERE pt.title = $1;
	`

	GetFilterAttributeValues = `
	
	SELECT DISTINCT a.id AS attribute_id, pa.title
		FROM product_attributes pa
		JOIN products p ON pa.product_id = p.id
		JOIN attributes a ON pa.attribute_id = a.id
		JOIN product_types pt ON p.product_type_id = pt.id
	WHERE pt.title = $1;
	`

	GetBrandsProductTypes = `
	SELECT DISTINCT b.id, b.title, b.title_farsi, b.image
		FROM brands b
		JOIN products p ON p.brand_id = b.id
		JOIN product_types pt ON p.product_type_id = pt.id
	WHERE pt.title = $1;
	`
)