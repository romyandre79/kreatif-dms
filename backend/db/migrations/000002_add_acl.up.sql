-- 9. Document Role Access (ACL)
CREATE TABLE IF NOT EXISTS document_role_access (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    document_id UUID NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    UNIQUE(document_id, role_id)
);

CREATE INDEX IF NOT EXISTS idx_doc_role_access_doc ON document_role_access(document_id);
CREATE INDEX IF NOT EXISTS idx_doc_role_access_role ON document_role_access(role_id);
