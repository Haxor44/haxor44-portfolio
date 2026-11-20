package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberhtml "github.com/gofiber/template/html/v2"

	"portfolio/internal/config"
	"portfolio/internal/database"
	"portfolio/internal/handlers"
	"portfolio/internal/mailer"
	"portfolio/internal/medium"
	"portfolio/internal/middleware"
	"portfolio/internal/routes"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("main: failed to connect to database: %v", err)
	}

	// Pull Medium stories into the articles table in the background.
	if cfg.MediumUsername != "" {
		medium.StartSync(context.Background(), db, cfg.MediumUsername, cfg.MediumSyncInterval)
	}

	m := mailer.New(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPassword, cfg.ContactToEmail)

	engine := fiberhtml.New("./web/templates", ".html")
	engine.Reload(!cfg.IsProduction())

	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: false,
		ErrorHandler:          errorHandler,
	})

	app.Use(recover.New())
	app.Use(middleware.RequestID())
	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} ${latency} ${method} ${path} reqid=${locals:request_id}\n",
	}))
	app.Use(compress.New())

	h := handlers.New(db, m)
	routes.Register(app, h)

	log.Printf("main: starting server on :%s (env=%s)", cfg.Port, cfg.AppEnv)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("main: server error: %v", err)
	}
}

// errorHandler renders a minimal JSON error for API routes and a plain
// message for page routes, so failures never leak stack traces.
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	log.Printf("error: %s %s -> %d: %v", c.Method(), c.OriginalURL(), code, err)
	return c.Status(code).JSON(fiber.Map{"error": err.Error()})
}
