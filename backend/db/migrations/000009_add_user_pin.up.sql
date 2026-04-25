-- Add pin field to users table for Level 2 approvals
ALTER TABLE users ADD COLUMN IF NOT EXISTS pin VARCHAR(255);

-- Optional: Set a comment to explain its purpose
COMMENT ON COLUMN users.pin IS '6-digit transaction PIN for high-security approvals (L2)';
