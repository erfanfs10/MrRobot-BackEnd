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

	OrderList = `
		SELECT
			o.id AS order_id,
			o.amount AS order_amount,
			o.shipping_price AS order_shipping_price,
			o.total_amount AS order_total_amount,
			o.status AS order_status,
			o.tracking_number AS order_tracking_number,
			o.created AS order_created,

			a.title AS address_title,
			a.address AS address_address,

			COALESCE(oi.items, '[]'::json) AS order_items
		FROM orders o
		LEFT JOIN addresses a ON a.id = o.address_id

		LEFT JOIN LATERAL (
			SELECT json_agg(
				json_build_object(
					'order_item_id', oi.id,
					'quantity', oi.quantity,
					'price', oi.price,
					'product_id', p.id,
					'product_title', p.title,
					'primary_image', pi.image
				)
			) AS items
			FROM order_items oi
			LEFT JOIN products p ON p.id = oi.product_id
			LEFT JOIN LATERAL (
				SELECT image
				FROM product_images
				WHERE product_id = p.id
				AND is_primary = TRUE
				LIMIT 1
			) pi ON TRUE
			WHERE oi.order_id = o.id
		) oi ON TRUE

		WHERE o.user_id = $1
		ORDER BY o.created DESC;

	`
)
