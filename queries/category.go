package queries

const (
	CategoryList = `
		SELECT id,title,title_farsi,image FROM categories;
	`
	CategoryGet = `
		SELECT id,title,title_farsi,image FROM categories
		WHERE title=$1
	`
)
