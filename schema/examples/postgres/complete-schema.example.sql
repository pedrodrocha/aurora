-- Generated from complete-schema.example.json
-- Complete schema with multiple related tables

-- Categories table with self-referencing relationship
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    parent_id INTEGER
);

-- Self-referencing foreign key
ALTER TABLE categories ADD CONSTRAINT categories_parent_fk 
    FOREIGN KEY (parent_id) REFERENCES categories(id) 
    ON UPDATE CASCADE ON DELETE SET NULL;

-- Index on parent_id
CREATE INDEX ON categories USING btree (parent_id);

-- Users table with various data types
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash CHAR(60) NOT NULL,
    profile_data JSONB,
    ip_address INET,
    subnet CIDR,
    mac_address MACADDR,
    avatar BYTEA,
    date_of_birth DATE,
    last_login TIMESTAMPTZ,
    session_timeout INTERVAL,
    is_verified BOOLEAN NOT NULL DEFAULT false
);

-- User indexes
CREATE UNIQUE INDEX ON users USING btree (username);
CREATE UNIQUE INDEX ON users USING btree (email);
CREATE INDEX ON users USING gin (profile_data);

-- Orders table with UUID primary key
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    user_id BIGINT NOT NULL,
    total_amount NUMERIC(15,4) NOT NULL,
    order_date TIMESTAMPTZ NOT NULL,
    shipping_address POINT,
    delivery_window TSRANGE,
    item_ids UUID[]
);

-- Order indexes
CREATE INDEX ON orders USING btree (user_id);
CREATE INDEX ON orders USING brin (order_date);
CREATE INDEX ON orders USING gist (shipping_address);

-- Foreign key relationships
ALTER TABLE orders ADD CONSTRAINT orders_user_fk 
    FOREIGN KEY (user_id) REFERENCES users(id) 
    ON UPDATE CASCADE ON DELETE CASCADE;