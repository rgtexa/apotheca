CREATE TABLE IF NOT EXISTS audittrails (
    id BIGSERIAL PRIMARY KEY,
    document_id BIGSERIAL NOT NULL REFERENCES documents (id),
    flow_id BIGSERIAL NOT NULL REFERENCES workflows (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    participants TEXT[][]
);