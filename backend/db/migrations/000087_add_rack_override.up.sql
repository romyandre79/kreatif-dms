ALTER TABLE racks 
ADD COLUMN is_full_override BOOLEAN DEFAULT FALSE,
ADD COLUMN override_reason TEXT;
