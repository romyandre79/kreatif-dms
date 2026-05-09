-- Add delivery_instructions to companies
ALTER TABLE companies ADD COLUMN IF NOT EXISTS delivery_instructions TEXT;

-- Update existing companies with a default instruction
UPDATE companies SET delivery_instructions = 'Harap lampirkan manifest ini pada bagian luar kotak pengiriman. Pastikan semua dokumen fisik disusun sesuai ID Registrasi.' WHERE delivery_instructions IS NULL;
