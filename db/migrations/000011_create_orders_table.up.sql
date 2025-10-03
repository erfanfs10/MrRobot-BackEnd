CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    address_id INT,
    amount NUMERIC(13, 3) NOT NULL,
    shipping_price NUMERIC(10, 3) DEFAULT 0,
    total_amount NUMERIC(13, 3) GENERATED ALWAYS AS ((amount + shipping_price)) STORED,
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (
        status IN (
            'pending',
            'payed',
            'sending',
            'delivered',
            'cancelled'
        )
    ),
    tracking_number VARCHAR(255) NOT NULL UNIQUE,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT fk_address FOREIGN KEY (address_id) REFERENCES addresses (id) ON DELETE SET NULL
);