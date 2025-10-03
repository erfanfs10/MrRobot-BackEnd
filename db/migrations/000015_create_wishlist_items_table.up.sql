CREATE TABLE IF NOT EXISTS wishlist_items (
    id SERIAL PRIMARY KEY,
    wishlist_id INT,
    product_id INT,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_wishlist FOREIGN KEY (wishlist_id) REFERENCES wishlists (id) ON DELETE SET NULL,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE SET NULL,
    UNIQUE (wishlist_id, product_id)
);