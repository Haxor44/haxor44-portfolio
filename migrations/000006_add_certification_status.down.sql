DROP INDEX IF EXISTS idx_certifications_status;

-- In-progress rows have no issue_date; remove them before restoring NOT NULL.
DELETE FROM certifications WHERE issue_date IS NULL;
ALTER TABLE certifications ALTER COLUMN issue_date SET NOT NULL;
ALTER TABLE certifications DROP COLUMN IF EXISTS status;
