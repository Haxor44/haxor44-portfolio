ALTER TABLE articles ADD COLUMN IF NOT EXISTS image_url VARCHAR(255);
ALTER TABLE articles ADD COLUMN IF NOT EXISTS tags      VARCHAR(255);

CREATE INDEX IF NOT EXISTS idx_articles_external_url ON articles (external_url);
