package queries

import "fmt"

func BuildProductPTCBQuery(filterBy string) string {
	validFilters := map[string]string{
		"product_type": "pt.title",
		"category":     "c.title",
		"brand":        "b.title",
	}

	column := validFilters[filterBy]

	baseQuery := `
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
			p.view AS product_view,
			p.sell AS product_sell,
            p.variant AS product_variant,
            p.variant_farsi AS product_variant_farsi,
            p.color_code AS product_color_code,
			p.created AS product_created,
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
        WHERE %s = $1
		ORDER BY p.discount DESC, p.status='av' DESC;
    `

	return fmt.Sprintf(baseQuery, column)
}

func BuildProductListQuery(filterBy string) string {
	validFilters := map[string]string{
		"new":  "p.created",
		"view": "p.view",
		"sell": "p.sell",
	}

	column := validFilters[filterBy]

	baseQuery := `
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
			p.view AS product_view,
			p.sell AS product_sell,
            p.variant AS product_variant,
            p.variant_farsi AS product_variant_farsi,
            p.color_code AS product_color_code,
			p.created AS product_created,
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

		ORDER BY %s DESC
		LIMIT 10;
    `

	return fmt.Sprintf(baseQuery, column)
}

const (
	ProductUpdateView = `
		UPDATE products SET view=view+1
		WHERE title=$1;
	`

	ProductUpdateSell = `
		UPDATE products SET sell=sell+$1
		WHERE id=$2;
	`

	ProductDetail = `
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
			p.view AS product_view,
			p.sell AS product_sell,
            p.variant AS product_variant,
            p.variant_farsi AS product_variant_farsi,
            p.color_code AS product_color_code,
			p.created AS product_created,
			b.title AS brand,
			b.title_farsi AS brand_farsi,
			c.title AS category,
			c.title_farsi AS category_farsi,
			pt.title AS product_type,
			pt.title_farsi AS product_type_farsi,
			pi.image AS primary_image,
			COALESCE(AVG(r.point), 0.0) AS point
			
		FROM products p
		LEFT JOIN 
			brands b ON p.brand_id = b.id
		LEFT JOIN 
			categories c ON p.category_id = c.id
		LEFT JOIN 
			product_types pt ON p.product_type_id = pt.id
		LEFT JOIN 
			rates r ON p.id = r.product_id
		LEFT JOIN LATERAL (
			SELECT image
			FROM product_images
			WHERE product_id = p.id AND is_primary = TRUE
			LIMIT 1
		) pi ON TRUE
		WHERE 
			p.title = $1
		GROUP BY p.id, p.title, p.title_farsi, p.description, p.used, p.status, 
				p.list_price, p.tax, p.discount, p.net_price, p.stock, 
				p.variant, p.variant_farsi, p.color_code,
				b.title, b.title_farsi, c.title, c.title_farsi, 
				pt.title, pt.title_farsi, pi.image;
	`

	GetProduct = `
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

		WHERE p.id = $1;
		
	`
)
