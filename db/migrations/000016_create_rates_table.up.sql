CREATE TABLE IF NOT EXISTS rates (
    id SERIAL PRIMARY KEY,
    user_id INT,
    product_id INT,
    title TEXT,
    body TEXT,
    point SMALLINT CHECK (
        point >= 1
        AND point <= 5
    ),
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE SET NULL,
    UNIQUE (user_id, product_id)
);