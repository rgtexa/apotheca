CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    username text NOT NULL,
    password text NOT NULL,
    firstname text NOT NULL,
    lastname text NOT NULL,
    department bigint NOT NULL REFERENCES departments (id)
)