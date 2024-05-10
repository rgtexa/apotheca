CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    email CITEXT UNIQUE NOT NULL,
    password_hash BYTEA NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    activated BOOL NOT NULL,
    version INTEGER NOT NULL DEFAULT 1,
    department BIGINT NOT NULL REFERENCES departments (id)
)