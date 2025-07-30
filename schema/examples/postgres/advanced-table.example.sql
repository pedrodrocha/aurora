-- Generated from advanced-table.example.json
-- Advanced products table showcasing complex features

CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    weight REAL,
    metadata JSONB,
    tags TEXT[],
    availability_period TSRANGE,
    location POINT,
    sku UUID NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- Named primary key constraint
ALTER TABLE products ADD CONSTRAINT products_pkey PRIMARY KEY (id);

-- Indexes with names
CREATE INDEX idx_products_category ON products USING btree (category_id);
CREATE INDEX idx_products_name_search ON products USING gin (name);
CREATE INDEX idx_products_metadata ON products USING gin (metadata);
CREATE INDEX idx_products_location ON products USING gist (location);
CREATE INDEX idx_products_price_range ON products USING btree (price, category_id);

-- Foreign key relationship
ALTER TABLE products ADD CONSTRAINT fk_products_category 
    FOREIGN KEY (category_id) REFERENCES categories(id) 
    ON UPDATE CASCADE ON DELETE RESTRICT;