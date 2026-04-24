ALTER TABLE users ADD COLUMN status VARCHAR(20) NOT NULL DEFAULT 'approved';
UPDATE users SET status = 'approved';

-- Index for admin to find pending users quickly
CREATE INDEX idx_users_status ON users(status);
