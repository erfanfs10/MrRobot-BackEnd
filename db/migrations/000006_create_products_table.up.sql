CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    brand_id INT,
    category_id INT,
    product_type_id INT,
    title VARCHAR(255) NOT NULL,
    title_farsi VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'av' CHECK (status IN ('av', 'na', 'cs')),
    used BOOLEAN DEFAULT FALSE,
    list_price NUMERIC(13, 3) NOT NULL,
    tax NUMERIC(5, 2) DEFAULT 0,
    discount NUMERIC(5, 2) DEFAULT 0,
    net_price NUMERIC(13, 3) GENERATED ALWAYS AS (
        (
            list_price + (list_price * tax / 100)
        ) - (list_price * discount / 100)
    ) STORED,
    view SMALLINT DEFAULT 0,
    sell SMALLINT DEFAULT 0,
    stock SMALLINT DEFAULT 0,
    color_code VARCHAR(50),
    variant VARCHAR(255),
    variant_farsi VARCHAR(255),
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_brand FOREIGN KEY (brand_id) REFERENCES brands (id) ON DELETE SET NULL,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE SET NULL,
    CONSTRAINT fk_product_type FOREIGN KEY (product_type_id) REFERENCES product_types (id) ON DELETE SET NULL
);