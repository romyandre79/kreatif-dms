-- Seed Document Types
INSERT INTO document_types (code, name, description) VALUES 
('INV', 'Invoice', 'Tagihan atau faktur dari vendor atau untuk pelanggan'),
('CON', 'Contract', 'Surat perjanjian atau kontrak kerjasama'),
('TAX', 'Tax Document', 'Dokumen perpajakan, PPh, PPN, dll'),
('LEG', 'Legal Document', 'Dokumen legalitas perusahaan, akta, ijin, dll'),
('EMP', 'Employee Record', 'Data karyawan, CV, kontrak kerja, dll'),
('OTH', 'Others', 'Dokumen lainnya')
ON CONFLICT (code) DO NOTHING;
