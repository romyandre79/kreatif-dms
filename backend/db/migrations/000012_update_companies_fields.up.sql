-- Add fields to companies table for PT Management
ALTER TABLE companies ADD COLUMN IF NOT EXISTS entity_id VARCHAR(50) UNIQUE;
ALTER TABLE companies ADD COLUMN IF NOT EXISTS npwp_status VARCHAR(50) DEFAULT 'PENDING';
ALTER TABLE companies ADD COLUMN IF NOT EXISTS location VARCHAR(255);
ALTER TABLE companies ADD COLUMN IF NOT EXISTS status VARCHAR(50) DEFAULT 'Aktif';

-- Update existing rows with some default entity IDs if any
UPDATE companies SET entity_id = 'PT-GEN-' || id::text WHERE entity_id IS NULL;
