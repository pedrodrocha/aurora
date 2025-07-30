-- Generated from basic-table.example.json
-- Basic users table with common patterns

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_active BOOLEAN NOT NULL DEFAULT true
);

-- Indexes
CREATE UNIQUE INDEX ON users (username);
CREATE INDEX ON users (email);