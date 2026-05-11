-- 000053_move_inbound_registration_to_registration_category.up.sql

UPDATE system_modules 
SET parent_id = 'cat_registration', sort_order = 45
WHERE id = 'inbound_registration';
