CREATE TABLE IF NOT EXISTS projects (
    id                   BIGSERIAL PRIMARY KEY,
    title                VARCHAR(150) NOT NULL,
    description          TEXT NOT NULL,
    tech_stack           VARCHAR(255) NOT NULL,
    github_url           VARCHAR(255),
    live_url             VARCHAR(255),
    blog_url             VARCHAR(255),
    deployment_platform  VARCHAR(100),
    image_url            VARCHAR(255),
    featured             BOOLEAN NOT NULL DEFAULT FALSE,
    display_order        INTEGER NOT NULL DEFAULT 0,
    created_at           TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_projects_featured ON projects (featured);
CREATE INDEX IF NOT EXISTS idx_projects_display_order ON projects (display_order);
