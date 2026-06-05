-- 000092_add_master_floors.down.sql

ALTER TABLE racks DROP COLUMN IF EXISTS floor_id;
DROP TABLE IF EXISTS floors;
