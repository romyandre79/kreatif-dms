-- 000075_set_department_heads.up.sql

DO $$
DECLARE
    it_manager_id UUID;
    it_dept_id UUID;
BEGIN
    -- Get IT Manager ID
    SELECT id INTO it_manager_id FROM users WHERE email = 'manager@kreatif.id';
    -- Get IT Department ID
    SELECT id INTO it_dept_id FROM departments WHERE name = 'IT Department';

    IF it_manager_id IS NOT NULL AND it_dept_id IS NOT NULL THEN
        UPDATE departments SET head_id = it_manager_id WHERE id = it_dept_id;
    END IF;
END $$;
