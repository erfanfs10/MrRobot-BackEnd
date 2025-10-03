package queries

const (
	ProductTypeList = `
		SELECT id,category_id,title,title_farsi,image FROM product_types;
	`

	ProductTypesBasedOnCategory = `
		SELECT id,category_id,title,title_farsi,image FROM product_types
		WHERE category_id=$1;
	`
)
