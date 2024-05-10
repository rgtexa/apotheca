CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    firstname text NOT NULL,
    lastname text NOT NULL,
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1,
    department bigint NOT NULL REFERENCES departments (id)
)