CREATE TABLE IF NOT EXISTS product_attributes (
    id SERIAL PRIMARY KEY,
    product_id INT,
    attribute_id INT,
    title VARCHAR(255) NOT NULL,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE SET NULL,
    CONSTRAINT fk_attribute FOREIGN KEY (attribute_id) REFERENCES attributes (id) ON DELETE SET NULL,
    UNIQUE (product_id, attribute_id)
);