DROP INDEX IF EXISTS idx_articles_external_url;

ALTER TABLE articles DROP COLUMN IF EXISTS tags;
ALTER TABLE articles DROP COLUMN IF EXISTS image_url;
