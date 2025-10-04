package queries

const (
	Rates = `
		SELECT
			rates.id AS rate_id,
			users.name AS username,
			rates.title AS rate_title,
			rates.body AS rate_body,
			rates.point AS rate_point,
			rates.created AS rate_created,
			rates.product_id AS rate_product_id
		FROM rates
		LEFT JOIN 
			users ON users.id = rates.user_id
		WHERE rates.product_id=$1;
	`
)
