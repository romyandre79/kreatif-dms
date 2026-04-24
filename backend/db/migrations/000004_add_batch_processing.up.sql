-- 1. Create Processing Batches Table
CREATE TABLE IF NOT EXISTS processing_batches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    total_files INT NOT NULL,
    processed_files INT DEFAULT 0,
    status VARCHAR(50) DEFAULT 'processing', -- 'processing', 'completed', 'failed'
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- 2. Add batch_id to documents
ALTER TABLE documents ADD COLUMN IF NOT EXISTS batch_id UUID REFERENCES processing_batches(id) ON DELETE SET NULL;

-- 3. Index for performance
CREATE INDEX IF NOT EXISTS idx_docs_batch ON documents(batch_id);
