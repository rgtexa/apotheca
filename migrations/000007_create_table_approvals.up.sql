CREATE TABLE IF NOT EXISTS approvals (
    id BIGSERIAL PRIMARY KEY,
    document_id BIGSERIAL NOT NULL REFERENCES documents (id),
    approvers text[][] NOT NULL,
    doc_majorver INTEGER NOT NULL,
    doc_minorver INTEGER NOT NULL,
    approved BOOLEAN DEFAULT false,
    workflow BIGSERIAL NOT NULL REFERENCES workflows (id),
    approved_on TIMESTAMPTZ,
    flow_id BIGSERIAL NOT NULL REFERENCES workflows (id)
);