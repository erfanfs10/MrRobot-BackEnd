CREATE TABLE IF NOT EXISTS filters (
    product_type_id BIGINT NOT NULL REFERENCES product_types (id) ON DELETE CASCADE,
    attribute_id BIGINT NOT NULL REFERENCES attributes (id) ON DELETE CASCADE,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_type_id, attribute_id)
);