-- 000054_create_physical_manifests.up.sql

CREATE TABLE IF NOT EXISTS physical_manifests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    manifest_no VARCHAR(100) UNIQUE NOT NULL,
    sender_id UUID NOT NULL REFERENCES users(id),
    department_id UUID NOT NULL REFERENCES departments(id),
    total_items INT NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending', 'received', 'partially_received', 'rejected'
    received_by UUID REFERENCES users(id),
    received_at TIMESTAMPTZ,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS physical_manifest_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    manifest_id UUID NOT NULL REFERENCES physical_manifests(id) ON DELETE CASCADE,
    document_id UUID NOT NULL REFERENCES documents(id),
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending', 'received', 'missing', 'damaged'
    verified_at TIMESTAMPTZ,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_manifest_no ON physical_manifests(manifest_no);
CREATE INDEX IF NOT EXISTS idx_manifest_status ON physical_manifests(status);
CREATE INDEX IF NOT EXISTS idx_manifest_items_id ON physical_manifest_items(manifest_id);

-- Update documents table to include manifest_id for tracking
ALTER TABLE documents ADD COLUMN IF NOT EXISTS current_manifest_id UUID REFERENCES physical_manifests(id);
