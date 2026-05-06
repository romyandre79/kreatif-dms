-- 000035_configure_user_menu.down.sql
-- Note: Down migration is best effort as it involves reversing ID renames

-- Optional: Revert IDs if needed, but usually down migrations just drop or revert state
-- For simplicity, we just leave the IDs as they are as they are more descriptive now.
-- But if we must revert to match 000019:
UPDATE system_modules SET id = 'new_doc' WHERE id = 'submit';
UPDATE system_modules SET id = 'checkout' WHERE id = 'loans';
UPDATE system_modules SET id = 'labels' WHERE id = 'tracking';
