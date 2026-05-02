-- 000022_add_mfa_and_cleanup_users.up.sql

-- Add MFA fields if not exists
ALTER TABLE users ADD COLUMN IF NOT EXISTS is_mfa_enabled BOOLEAN DEFAULT false;
ALTER TABLE users ADD COLUMN IF NOT EXISTS mfa_secret VARCHAR(255);

-- Ensure department_id is handled (already exists in init_schema, but making sure)
-- Ensure avatar_url and signature_url exist (already in 000007)

-- Add comments for clarity
COMMENT ON COLUMN users.is_mfa_enabled IS 'Flag to indicate if Multi-Factor Authentication is enabled for the user';
COMMENT ON COLUMN users.mfa_secret IS 'TOTP secret for MFA verification';
