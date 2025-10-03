package queries

const (
	BrandList = `
		SELECT id,title,title_farsi,image FROM brands;
	`
	BrandGet = `
		SELECT id,title,title_farsi,image FROM brands
		WHERE title=$1
	`
)
