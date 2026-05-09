-- Add logo_url to companies
ALTER TABLE companies ADD COLUMN IF NOT EXISTS logo_url TEXT;

-- Update existing companies with a placeholder logo if needed
UPDATE companies SET logo_url = 'https://ui-avatars.com/api/?name=' || REPLACE(name, ' ', '+') || '&background=1E3A5F&color=fff&size=128' WHERE logo_url IS NULL;
