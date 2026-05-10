-- 000039_add_announcement_notes.up.sql

-- 1. Add announcement_notes setting
INSERT INTO system_settings (category, key, value, value_type, description) VALUES 
('dashboard', 'announcement_notes', '### Detail Pembaruan v1.2\n\n- **OCR Engine**: Peningkatan akurasi hingga 95% untuk dokumen tulisan tangan.\n- **Role Management**: Penambahan role Document Controller untuk manajemen alur dokumen yang lebih fleksibel.\n- **Performance**: Optimasi query dashboard untuk loading yang lebih cepat.\n- **UI/UX**: Pembaruan tema gelap (dark mode) yang lebih konsisten.\n\nJika Anda menemui kendala, silakan hubungi tim IT Support di internal nomor 123.', 'string', 'Detail catatan rilis di dashboard')
ON CONFLICT (category, key) DO UPDATE SET 
    value = EXCLUDED.value,
    description = EXCLUDED.description;

-- 2. Add Announcement Module under cat_config
INSERT INTO system_modules (id, name, category, path, icon, allowed_actions, sort_order, parent_id) VALUES
('announcements', 'Pengumuman Dashboard', 'Configuration', '/admin/announcements', 'LucideMegaphone', ARRAY['VIEW', 'CREATE', 'UPDATE', 'DELETE'], 100, 'cat_config')
ON CONFLICT (id) DO UPDATE SET 
    name = EXCLUDED.name,
    path = EXCLUDED.path,
    icon = EXCLUDED.icon,
    parent_id = EXCLUDED.parent_id;

-- 3. Grant permission to superadmin
INSERT INTO role_permissions (role_id, module_id, action)
SELECT r.id, 'announcements', UNNEST(ARRAY['VIEW', 'CREATE', 'UPDATE', 'DELETE'])
FROM roles r
WHERE r.name = 'superadmin'
ON CONFLICT DO NOTHING;
