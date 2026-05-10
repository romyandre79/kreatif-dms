-- 000038_seed_dashboard_announcements.up.sql

INSERT INTO system_settings (category, key, value, value_type, description) VALUES 
('dashboard', 'announcement_title', 'Sistem Update v1.2', 'string', 'Judul pengumuman di dashboard'),
('dashboard', 'announcement_message', 'Kami telah memperbarui sistem dengan fitur OCR yang lebih cepat dan manajemen role yang lebih ketat. Silakan hubungi admin jika ada kendala.', 'string', 'Pesan pengumuman di dashboard'),
('dashboard', 'announcement_active', 'true', 'boolean', 'Status aktif pengumuman di dashboard')
ON CONFLICT (category, key) DO UPDATE SET 
    value = EXCLUDED.value,
    description = EXCLUDED.description;
