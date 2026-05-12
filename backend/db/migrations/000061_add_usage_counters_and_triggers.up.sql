-- 000061_add_usage_counters_and_triggers.up.sql

-- 1. Add current usage columns to hierarchy
ALTER TABLE companies ADD COLUMN current_docs_count INT DEFAULT 0;
ALTER TABLE branches ADD COLUMN current_docs_count INT DEFAULT 0;
ALTER TABLE departments ADD COLUMN current_docs_count INT DEFAULT 0;
ALTER TABLE racks ADD COLUMN current_boxes_count INT DEFAULT 0;
ALTER TABLE boxes ADD COLUMN current_docs_count INT DEFAULT 0;

-- 2. Create Trigger Function for Document Capacity
CREATE OR REPLACE FUNCTION fn_sync_document_capacity()
RETURNS TRIGGER AS $$
BEGIN
    -- Handle DELETE or moving OUT of a box/dept
    IF (TG_OP = 'DELETE') OR (TG_OP = 'UPDATE') THEN
        -- Only decrement if the old status was 'active' or 'archived'
        IF (OLD.status IN ('active', 'archived')) AND (OLD.box_id IS NOT NULL) THEN
            -- Decrement Box
            UPDATE boxes SET current_docs_count = current_docs_count - 1 WHERE id = OLD.box_id;
            -- Decrement Org Hierarchy
            UPDATE departments SET current_docs_count = current_docs_count - 1 WHERE id = OLD.department_id;
            UPDATE branches SET current_docs_count = current_docs_count - 1 WHERE id = OLD.branch_id;
            UPDATE companies SET current_docs_count = current_docs_count - 1 WHERE id = OLD.company_id;
        END IF;
    END IF;

    -- Handle INSERT or moving INTO a box/dept
    IF (TG_OP = 'INSERT') OR (TG_OP = 'UPDATE') THEN
        -- Only increment if the new status is 'active' or 'archived'
        IF (NEW.status IN ('active', 'archived')) AND (NEW.box_id IS NOT NULL) THEN
            -- Increment Box
            UPDATE boxes SET current_docs_count = current_docs_count + 1 WHERE id = NEW.box_id;
            -- Increment Org Hierarchy
            UPDATE departments SET current_docs_count = current_docs_count + 1 WHERE id = NEW.department_id;
            UPDATE branches SET current_docs_count = current_docs_count + 1 WHERE id = NEW.branch_id;
            UPDATE companies SET current_docs_count = current_docs_count + 1 WHERE id = NEW.company_id;
        END IF;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- 3. Create Trigger Function for Box Capacity (on Racks)
CREATE OR REPLACE FUNCTION fn_sync_box_capacity()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        UPDATE racks SET current_boxes_count = current_boxes_count - 1 WHERE id = OLD.rack_id;
    ELSIF (TG_OP = 'INSERT') THEN
        UPDATE racks SET current_boxes_count = current_boxes_count + 1 WHERE id = NEW.rack_id;
    ELSIF (TG_OP = 'UPDATE') AND (OLD.rack_id <> NEW.rack_id) THEN
        UPDATE racks SET current_boxes_count = current_boxes_count - 1 WHERE id = OLD.rack_id;
        UPDATE racks SET current_boxes_count = current_boxes_count + 1 WHERE id = NEW.rack_id;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- 4. Assign Triggers
CREATE TRIGGER trg_document_capacity_sync
AFTER INSERT OR UPDATE OR DELETE ON documents
FOR EACH ROW EXECUTE FUNCTION fn_sync_document_capacity();

CREATE TRIGGER trg_box_capacity_sync
AFTER INSERT OR UPDATE OR DELETE ON boxes
FOR EACH ROW EXECUTE FUNCTION fn_sync_box_capacity();

-- 5. Initialize Existing Data
UPDATE boxes b SET current_docs_count = (SELECT COUNT(*) FROM documents WHERE box_id = b.id AND status IN ('active', 'archived'));
UPDATE racks r SET current_boxes_count = (SELECT COUNT(*) FROM boxes WHERE rack_id = r.id);
UPDATE departments d SET current_docs_count = (SELECT COUNT(*) FROM documents WHERE department_id = d.id AND status IN ('active', 'archived'));
UPDATE branches br SET current_docs_count = (SELECT COUNT(*) FROM documents WHERE branch_id = br.id AND status IN ('active', 'archived'));
UPDATE companies c SET current_docs_count = (SELECT COUNT(*) FROM documents WHERE company_id = c.id AND status IN ('active', 'archived'));
