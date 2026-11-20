-- Certifications can now be "in_progress", which have no issue date yet.
ALTER TABLE certifications ADD COLUMN IF NOT EXISTS status VARCHAR(20) NOT NULL DEFAULT 'earned';
ALTER TABLE certifications ALTER COLUMN issue_date DROP NOT NULL;

CREATE INDEX IF NOT EXISTS idx_certifications_status ON certifications (status);
