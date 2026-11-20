package routes

import (
	"github.com/gofiber/fiber/v2"

	"portfolio/internal/handlers"
)

// Register wires every route to its handler.
//
// There's no admin/write API yet — every route here is a public,
// read-only page or the contact form. See JWT_ADMIN_API_GUIDE.md for
// how to add a JWT-protected /admin/api group with create/update/delete
// handlers for projects, certifications, and articles.
func Register(app *fiber.App, h *handlers.Handler) {
	app.Get("/health", h.Health)

	app.Get("/", h.Home)
	app.Get("/projects", h.ProjectsPage)
	app.Get("/certifications", h.CertificationsPage)
	app.Get("/blog", h.BlogPage)
	app.Get("/contact", h.ContactPage)
	app.Post("/contact", h.SubmitContact)

	app.Static("/static", "./web/static")

	api := app.Group("/api")
	api.Get("/projects", h.ListProjectsAPI)
}
