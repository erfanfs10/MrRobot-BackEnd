package queries

const (
	OrderCreate = `
		INSERT INTO orders(user_id,address_id,amount,
			shipping_price, status, tracking_number)
		VALUES($1,$2,$3,$4,$5,$6)
		RETURNING id
	`

	OrderItemsCreate = `
		INSERT INTO order_items(order_id,product_id,quantity,price)
		VALUES($1,$2,$3,$4)
	`
)