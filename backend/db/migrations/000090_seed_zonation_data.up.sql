-- 000090_seed_zonation_data.up.sql

-- 1. Rename existing departments for better matching with zonation names
UPDATE departments SET name = 'IT Infrastructure' WHERE id = '550e8400-e29b-41d4-a716-446655440002';
UPDATE departments SET name = 'Finance & Tax' WHERE id = '550e8400-e29b-41d4-a716-446655440003';

-- 2. Insert new departments
INSERT INTO departments (id, branch_id, name) VALUES 
('550e8400-e29b-41d4-a716-446655440101', '550e8400-e29b-41d4-a716-446655440001', 'Legal & Compliance'),
('550e8400-e29b-41d4-a716-446655440102', '550e8400-e29b-41d4-a716-446655440001', 'Human Resources'),
('550e8400-e29b-41d4-a716-446655440103', '550e8400-e29b-41d4-a716-446655440001', 'Procurement')
ON CONFLICT (id) DO NOTHING;

-- 3. Loop and populate racks for grid coordinates A1 to H8
DO $$
DECLARE
    cols TEXT[] := ARRAY['A','B','C','D','E','F','G','H'];
    col_idx INT;
    row_idx INT;
    col_letter TEXT;
    rack_name TEXT;
    dept_id UUID;
    is_full BOOLEAN;
    reason TEXT;
    rack_uuid UUID;
BEGIN
    FOR col_idx IN 1..8 LOOP
        col_letter := cols[col_idx];
        FOR row_idx IN 1..8 LOOP
            rack_name := 'RACK-' || col_letter || row_idx;
            is_full := FALSE;
            reason := NULL;
            
            -- Assign department and status based on mockup pattern:
            -- LOCKED:
            IF (col_letter = 'A' AND row_idx = 1) OR 
               (col_letter = 'A' AND row_idx = 7) OR
               (col_letter = 'B' AND row_idx = 5) OR
               (col_letter = 'B' AND row_idx = 6) OR
               (col_letter = 'C' AND row_idx = 2) OR
               (col_letter = 'C' AND row_idx = 3) OR
               (col_letter = 'C' AND row_idx = 4) OR
               (col_letter = 'C' AND row_idx = 7) OR
               (col_letter = 'D' AND row_idx = 1) OR
               (col_letter = 'D' AND row_idx = 6) OR
               (col_letter = 'D' AND row_idx = 7) OR
               (col_letter = 'E' AND row_idx = 2) OR
               (col_letter = 'E' AND row_idx = 3) OR
               (col_letter = 'E' AND row_idx = 4) OR
               (col_letter = 'F' AND row_idx = 3) OR
               (col_letter = 'F' AND row_idx = 5) OR
               (col_letter = 'G' AND row_idx IN (3,4,5,6,7,8)) OR
               (col_letter = 'H' AND row_idx IN (1,2,4,5,6))
            THEN
                -- LOCKED / FULL OVERRIDE (e.g. HR Department, but locked)
                dept_id := '550e8400-e29b-41d4-a716-446655440102'; -- HR
                is_full := TRUE;
                reason := 'Manual lock for permanent archives.';
            
            -- LEGAL (Navy Blue):
            ELSIF (col_letter = 'A' AND row_idx IN (3,4,6)) OR
                  (col_letter = 'B' AND row_idx IN (1,3,7,8)) OR
                  (col_letter = 'C' AND row_idx IN (5,6)) OR
                  (col_letter = 'D' AND row_idx IN (2,3,4,5,8)) OR
                  (col_letter = 'E' AND row_idx IN (1,7,8)) OR
                  (col_letter = 'F' AND row_idx IN (1,2,7))
            THEN
                dept_id := '550e8400-e29b-41d4-a716-446655440101'; -- Legal
            
            -- FINANCE (Light Blue / Teal):
            ELSIF (col_letter = 'C' AND row_idx = 1) OR
                  (col_letter = 'E' AND row_idx = 6) OR
                  (col_letter = 'H' AND row_idx = 8)
            THEN
                dept_id := '550e8400-e29b-41d4-a716-446655440003'; -- Finance
                
            -- GENERIC/PROCUREMENT (Brown):
            ELSIF (col_letter = 'A' AND row_idx IN (2,5,8)) OR
                  (col_letter = 'B' AND row_idx IN (2,4)) OR
                  (col_letter = 'C' AND row_idx = 8) OR
                  (col_letter = 'E' AND row_idx = 5) OR
                  (col_letter = 'F' AND row_idx IN (4,6)) OR
                  (col_letter = 'G' AND row_idx IN (1,2)) OR
                  (col_letter = 'H' AND row_idx IN (3,7))
            THEN
                dept_id := '550e8400-e29b-41d4-a716-446655440103'; -- Procurement
                
            -- IT Infrastructure:
            ELSE
                dept_id := '550e8400-e29b-41d4-a716-446655440002'; -- IT
            END IF;
            
            -- Insert or update rack
            SELECT id INTO rack_uuid FROM racks WHERE name = rack_name;
            IF rack_uuid IS NULL THEN
                INSERT INTO racks (department_id, name, location_detail, map_pos_x, map_pos_y, is_full_override, override_reason)
                VALUES (dept_id, rack_name, 'Lantai 2, Grid ' || col_letter || row_idx, col_idx, row_idx, is_full, reason);
            ELSE
                UPDATE racks SET 
                    department_id = dept_id, 
                    location_detail = 'Lantai 2, Grid ' || col_letter || row_idx, 
                    map_pos_x = col_idx, 
                    map_pos_y = row_idx,
                    is_full_override = is_full,
                    override_reason = reason
                WHERE id = rack_uuid;
            END IF;
            
        END LOOP;
    END LOOP;
END $$;
