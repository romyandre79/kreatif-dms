-- Migration: Create document_categories table and link it to document_types
CREATE TABLE IF NOT EXISTS document_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Add category_id to document_types
ALTER TABLE document_types ADD COLUMN IF NOT EXISTS category_id UUID REFERENCES document_categories(id) ON DELETE SET NULL;

-- Seed initial categories
INSERT INTO document_categories (code, name, description) VALUES 
('LEG', 'Legal Document', 'Dokumen hukum, akta, ijin, dsb'),
('FIN', 'Financial Document', 'Dokumen keuangan, invoice, pajak, dsb'),
('HR', 'HR Document', 'Dokumen kepegawaian, CV, kontrak kerja, dsb'),
('OPR', 'Operational Document', 'Dokumen operasional harian'),
('OTH', 'Others', 'Dokumen lainnya')
ON CONFLICT (code) DO NOTHING;

-- Map existing types to categories (optional, based on seed data)
UPDATE document_types SET category_id = (SELECT id FROM document_categories WHERE code = 'FIN') WHERE code = 'INV';
UPDATE document_types SET category_id = (SELECT id FROM document_categories WHERE code = 'LEG') WHERE code IN ('CON', 'LEG');
UPDATE document_types SET category_id = (SELECT id FROM document_categories WHERE code = 'HR') WHERE code = 'EMP';
UPDATE document_types SET category_id = (SELECT id FROM document_categories WHERE code = 'OTH') WHERE code = 'OTH';
