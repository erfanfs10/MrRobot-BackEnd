package queries

const (
	CreateWishList = `
		INSERT INTO wishlists(user_id)
			VALUES($1);
	`
	GetWishListIDCurrentUser = `
		SELECT id FROM wishlists
			WHERE user_id=$1;
	`
	GetWishListItemsProductIDs = `
		SELECT product_id FROM wishlist_items
			WHERE wishlist_id=$1
	`

	GetWishListItems = `
		SELECT 
    		p.id AS product_id,
			p.title AS product_title,
			p.title_farsi AS product_title_farsi,
			p.description AS product_description,
			p.used AS product_used,
            p.status AS product_status,
            p.list_price AS product_list_price,
            p.tax AS product_tax,
            p.discount AS product_discount,
            p.net_price AS product_net_price,
            p.stock AS product_stock,
            p.variant AS product_variant,
            p.variant_farsi AS product_variant_farsi,
            p.color_code AS product_color_code,
			b.title AS brand,
			b.title_farsi AS brand_farsi,
			c.title AS category,
			c.title_farsi AS category_farsi,
			pt.title AS product_type,
			pt.title_farsi AS product_type_farsi,
			pi.image AS primary_image

		FROM products p
		LEFT JOIN brands b ON p.brand_id = b.id
		LEFT JOIN categories c ON p.category_id = c.id
		LEFT JOIN product_types pt ON p.product_type_id = pt.id
		LEFT JOIN LATERAL (
			SELECT image
			FROM product_images
			WHERE product_id = p.id AND is_primary = TRUE
			LIMIT 1
		) pi ON TRUE

		WHERE p.id IN (?);
		
	`

	CreateWishListItem = `
		INSERT INTO wishlist_items(wishlist_id, product_id) 
			VALUES($1, $2);
	`

	DeleteWishListItem = `
		DELETE FROM wishlist_items
		WHERE wishlist_id=$1 AND product_id=$2
	`
)
