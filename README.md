# DevOps Portfolio

A server-rendered portfolio site built with Go, [Fiber](https://gofiber.io/), and PostgreSQL —
showcasing projects, certifications, and writing. Content is currently
read-only via the public routes / seeded directly in the database; there's no
admin API yet.

> **Docker & CI/CD:** The `Dockerfile` and `.github/workflows/` pipeline are intentionally
> **not** included here — see [`DOCKER_AND_CICD_GUIDE.md`](./DOCKER_AND_CICD_GUIDE.md) for a
> full reference to write them yourself.
>
> **Admin API:** also intentionally not included — see
> [`JWT_ADMIN_API_GUIDE.md`](./JWT_ADMIN_API_GUIDE.md) to build the JWT-protected
> create/update/delete endpoints yourself, and
> [`GORM_AND_MIGRATIONS_GUIDE.md`](./GORM_AND_MIGRATIONS_GUIDE.md) for how the
> GORM setup and `migrations/` folder in this repo work.

## Stack

- **Backend:** Go, [Fiber v2](https://gofiber.io/) (HTTP), [GORM](https://gorm.io/) (ORM), [godotenv](https://github.com/joho/godotenv)
- **Database:** PostgreSQL
- **Auth:** none yet — see `JWT_ADMIN_API_GUIDE.md`
- **Frontend:** Server-rendered HTML templates + vanilla CSS/JS, [Alpine.js](https://alpinejs.dev/) for the mobile nav
- **Email:** `net/smtp` for contact-form notifications (optional — see `.env.example`)

## Project structure

```
cmd/server/main.go          Entrypoint: config, DB, template engine, routes, server start
internal/
  config/                   Env var loading
  database/                 GORM connection + AutoMigrate
  models/                   Project, Certification, Article, ContactMessage
  handlers/                 Route handlers, grouped by resource
  middleware/               Request ID (JWT auth middleware not built yet)
  mailer/                   SMTP contact-form notifications
  routes/                   Route registration
migrations/                 Raw SQL migrations (golang-migrate format), mirroring the GORM models
web/
  templates/                HTML templates (layout + one per page)
  static/                   CSS, JS, images
docker-compose.yml          Local Postgres + Adminer for development
.env.example                All environment variables, documented
```

## Local development (without Docker for the app itself)

1. **Start Postgres** (and Adminer, a DB UI, at http://localhost:8081):

   ```bash
   docker compose up -d postgres adminer
   ```

2. **Copy env vars:**

   ```bash
   cp .env.example .env
   # edit SMTP_* as needed
   ```

3. **Run the app** (GORM's `AutoMigrate` creates tables on boot — no manual migration step needed for local dev):

   ```bash
   go run ./cmd/server
   ```

   Visit http://localhost:3000.

4. **Add content:** there's no admin API yet (see `JWT_ADMIN_API_GUIDE.md`), so
   for now insert rows directly, e.g. via Adminer (http://localhost:8081) or `psql`:

   ```sql
   INSERT INTO projects (title, description, tech_stack, github_url, deployment_platform, featured, created_at, updated_at)
   VALUES ('My Project', 'What it does', 'Go,Docker,AWS', 'https://github.com/you/repo', 'AWS ECS', true, now(), now());
   ```

## Database migrations

The app uses GORM `AutoMigrate` on boot, so nothing extra is required for local
development or a fresh deployment. The `migrations/` directory contains the
equivalent raw SQL migrations (in [golang-migrate](https://github.com/golang-migrate/migrate)
format) for teams that want explicit, reviewable migration history instead.
See [`GORM_AND_MIGRATIONS_GUIDE.md`](./GORM_AND_MIGRATIONS_GUIDE.md) for how
this is wired up and how to run/extend the migrations.

## Routes

| Method | Path             | Auth   | Purpose                        |
|--------|------------------|--------|---------------------------------|
| GET    | `/health`        | none   | Liveness/readiness probe        |
| GET    | `/`               | none   | Home page                       |
| GET    | `/projects`       | none   | Projects showcase               |
| GET    | `/certifications` | none   | Certifications                  |
| GET    | `/blog`           | none   | Articles list                   |
| GET    | `/contact`        | none   | Contact form                    |
| POST   | `/contact`        | none   | Submit contact form             |
| GET    | `/api/projects`   | none   | Projects as JSON                |

More routes (`/admin/login`, `/admin/api/*`) get added here once you follow
`JWT_ADMIN_API_GUIDE.md`.

## Environment variables

See [`.env.example`](./.env.example) for the full, documented list.

## Deployment

This app is stateless (all persistent state lives in Postgres) and reads its
entire configuration from environment variables, so it's ready to run behind
any container platform — AWS ECS/EKS, DigitalOcean App Platform, Railway, etc.
Once you've written the `Dockerfile`, see `DOCKER_AND_CICD_GUIDE.md` for
guidance on building, publishing, and deploying it via GitHub Actions.
