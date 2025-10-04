package queries

const (
	ProductImages = `
		SELECT
			id AS product_image_id,
			image AS product_image_image,
			is_primary AS product_image_is_primary,
			product_id AS product_image_product_id
		FROM product_images
		WHERE product_images.product_id=$1; 
	`
)
