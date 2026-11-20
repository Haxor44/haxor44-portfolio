CREATE TABLE IF NOT EXISTS certifications (
    id              BIGSERIAL PRIMARY KEY,
    name            VARCHAR(150) NOT NULL,
    issuing_org     VARCHAR(150) NOT NULL,
    issue_date      TIMESTAMPTZ NOT NULL,
    expiry_date     TIMESTAMPTZ,
    credential_url  VARCHAR(255),
    badge_icon_url  VARCHAR(255),
    display_order   INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_certifications_issue_date ON certifications (issue_date DESC);
