# Haxor44 — DevOps Portfolio

Personal portfolio of a DevOps & Cloud Infrastructure Engineer. It's a server-rendered
web app that showcases projects, certifications and writing, and takes messages
through a contact form.

## Features

- **Projects:** featured work with tech stack, GitHub/live links and deployment platform.
- **Certifications:** earned and in-progress certifications, with links to verify each credential.
- **Blog:** articles synced automatically from Medium via its RSS feed, on a configurable interval.
- **Contact form:** messages are stored in PostgreSQL, with optional email notifications over SMTP.
- **JSON API:** `GET /api/projects` returns the projects as JSON.
- **Health check:** `GET /health` for container liveness/readiness probes.
- **UI:** dark/light theme, animated starfield and cluster topology hero, eagle footer landscape.

## Tech stack

| Layer     | Tools                                                                 |
|-----------|-----------------------------------------------------------------------|
| Backend   | Go, [Fiber v2](https://gofiber.io/), [GORM](https://gorm.io/)          |
| Database  | PostgreSQL 16                                                          |
| Frontend  | Go HTML templates, vanilla CSS/JS, [Alpine.js](https://alpinejs.dev/) |
| Container | Multi-stage Docker build (Alpine runtime, non-root user)               |
| CI/CD     | GitHub Actions: lint (golangci-lint), build, Docker image push         |

## Project structure

```
cmd/server/          Entrypoint: config, database, templates, routes
internal/
  config/            Environment variable loading
  database/          GORM connection and AutoMigrate
  models/            Project, Certification, Article, ContactMessage
  handlers/          Page and API handlers
  medium/            Medium RSS fetch, parse and background sync
  mailer/            SMTP contact-form notifications
  middleware/        Request logging
  routes/            Route registration
migrations/          SQL migrations mirroring the GORM models
seeds/               Initial content (projects and certifications)
web/templates/       HTML layout, pages and shared partials
web/static/          CSS, JS, favicon and images
.github/workflows/   CI pipeline
```

## Running locally

```bash
docker compose up -d postgres adminer   # Postgres + Adminer UI on :8081
cp .env.example .env                    # then fill in values
go run ./cmd/server                     # http://localhost:3000
```

Tables are created automatically on startup. To load the initial content:

```bash
docker compose exec -T postgres psql -U portfolio -d portfolio < seeds/001_cv_content.sql
```


