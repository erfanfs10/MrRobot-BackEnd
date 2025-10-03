CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    content TEXT NOT NULL,
    excerpt TEXT,
    image VARCHAR(255),
    status VARCHAR(20) NOT NULL DEFAULT 'draft' CHECK (
        status IN ('draft', 'published')
    ),
    published_at TIMESTAMPTZ,
    meta_title VARCHAR(255),
    meta_description VARCHAR(255),
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);