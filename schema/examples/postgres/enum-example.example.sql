-- Generated from enum-example.example.json
-- User roles table with custom enum type

-- Create custom enum type first (required in PostgreSQL)
CREATE TYPE user_role_type AS ENUM ('admin', 'user', 'moderator');

CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    role user_role_type NOT NULL,
    granted_at TIMESTAMPTZ NOT NULL
);

-- Unique composite index
CREATE UNIQUE INDEX ON user_roles USING btree (user_id, role);