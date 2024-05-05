CREATE TABLE IF NOT EXISTS documents (
    id bigserial PRIMARY KEY,  
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    subject text NOT NULL,
    summary text NOT NULL,
    filename text NOT NULL,
    majorver smallint NOT NULL DEFAULT 0,
    minorver smallint NOT NULL DEFAULT 0,
    author bigint NOT NULL REFERENCES users (id)
);