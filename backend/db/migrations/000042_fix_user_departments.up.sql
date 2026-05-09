-- Assign IT Department to users that don't have one (for testing)
UPDATE users 
SET department_id = '550e8400-e29b-41d4-a716-446655440002' 
WHERE email IN ('admin@kreatif.id', 'controller@kreatif.id') 
AND department_id IS NULL;
