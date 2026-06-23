-- 000092_add_master_floors.up.sql

-- 1. Create floors table
CREATE TABLE IF NOT EXISTS floors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. Seed initial floors
INSERT INTO floors (id, name, code) VALUES 
('550e8400-e29b-41d4-a716-446655440201', 'Lantai 02-B (Warehouse Main)', 'Floor 02-B'),
('550e8400-e29b-41d4-a716-446655440202', 'Lantai 01-A (Annex Archive)', 'Floor 01-A')
ON CONFLICT (id) DO NOTHING;

-- 3. Add floor_id column to racks
ALTER TABLE racks ADD COLUMN floor_id UUID REFERENCES floors(id) ON DELETE SET NULL;

-- 4. Update existing racks to Floor 02-B by default
UPDATE racks SET floor_id = '550e8400-e29b-41d4-a716-446655440201';
