package queries

const (
	ProductTypeList = `
		SELECT id,category_id,title,title_farsi,image FROM product_types
		WHERE image IS NOT NULL
		ORDER BY created;
	`

	ProductTypesBasedOnCategory = `
		SELECT id,category_id,title,title_farsi,image FROM product_types
		WHERE category_id=$1;
	`
)
