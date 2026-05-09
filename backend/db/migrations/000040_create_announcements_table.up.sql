-- 000040_create_announcements_table.up.sql

CREATE TABLE announcements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    message TEXT NOT NULL,
    notes TEXT,
    is_active BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by UUID REFERENCES users(id)
);

-- Index for active announcement
CREATE INDEX idx_announcements_active ON announcements(is_active) WHERE is_active = true;

-- Seed initial data from system_settings if possible
INSERT INTO announcements (title, message, notes, is_active)
SELECT 
    COALESCE(MAX(CASE WHEN key = 'announcement_title' THEN value END), 'System Update'),
    COALESCE(MAX(CASE WHEN key = 'announcement_message' THEN value END), 'Sistem telah diperbarui.'),
    MAX(CASE WHEN key = 'announcement_notes' THEN value END),
    COALESCE(MAX(CASE WHEN key = 'announcement_active' THEN value END) = 'true', false)
FROM system_settings 
WHERE category = 'dashboard';
