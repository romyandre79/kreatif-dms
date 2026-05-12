-- 000060_add_storage_capacity_and_zonation.up.sql

-- 1. Add capacity limits to organizations
ALTER TABLE companies ADD COLUMN max_docs_capacity INT DEFAULT 1000000;
ALTER TABLE branches ADD COLUMN max_docs_capacity INT DEFAULT 100000;
ALTER TABLE departments ADD COLUMN max_docs_capacity INT DEFAULT 10000;

-- 2. Add capacity limits to storage hierarchy
ALTER TABLE racks ADD COLUMN max_boxes_capacity INT DEFAULT 50;
ALTER TABLE boxes ADD COLUMN max_docs_capacity INT DEFAULT 100;
ALTER TABLE ordners ADD COLUMN max_docs_capacity INT DEFAULT 20;

-- 3. Add zonation (allowed categories and types)
ALTER TABLE racks ADD COLUMN allowed_category_ids UUID[] DEFAULT '{}';
ALTER TABLE racks ADD COLUMN allowed_type_ids UUID[] DEFAULT '{}';

-- 4. Create a view to easily check current occupancy
CREATE OR REPLACE VIEW storage_occupancy AS
SELECT 
    r.id as rack_id,
    r.name as rack_name,
    r.max_boxes_capacity,
    (SELECT COUNT(*) FROM boxes WHERE rack_id = r.id) as current_boxes,
    bx.id as box_id,
    bx.name as box_name,
    bx.max_docs_capacity as box_max_docs,
    (SELECT COUNT(*) FROM documents WHERE box_id = bx.id) as current_docs
FROM racks r
LEFT JOIN boxes bx ON bx.rack_id = r.id;
