CREATE TYPE flows AS ENUM ('create', 'review', 'update', 'collaborate', 'approve', 'release')

CREATE TABLE IF NOT EXISTS workflows (
    id BIGSERIAL PRIMARY KEY,
    flow_type flows NOT NULL DEFAULT 'create',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    initiated_by BIGSERIAL NOT NULL REFERENCES users (id),
    document_id BIGSERIAL NOT NULL REFERENCES documents (id)
);