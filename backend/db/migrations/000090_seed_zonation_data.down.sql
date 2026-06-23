-- 000090_seed_zonation_data.down.sql

-- 1. Remove seeded racks
DELETE FROM racks WHERE name LIKE 'RACK-%';

-- 2. Remove seeded departments
DELETE FROM departments WHERE id IN (
    '550e8400-e29b-41d4-a716-446655440101', -- Legal & Compliance
    '550e8400-e29b-41d4-a716-446655440102', -- Human Resources
    '550e8400-e29b-41d4-a716-446655440103'  -- Procurement
);

-- 3. Restore original department names
UPDATE departments SET name = 'IT Department' WHERE id = '550e8400-e29b-41d4-a716-446655440002';
UPDATE departments SET name = 'Finance Department' WHERE id = '550e8400-e29b-41d4-a716-446655440003';
