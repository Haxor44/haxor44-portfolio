package handlers

import (
	"github.com/gofiber/fiber/v2"

	"portfolio/internal/models"
)

// ProjectsPage renders the projects showcase page.
func (h *Handler) ProjectsPage(c *fiber.Ctx) error {
	var projects []models.Project
	if err := h.DB.Order("display_order asc, created_at desc").Find(&projects).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to load projects")
	}
	return c.Render("projects", fiber.Map{
		"Title":    "Projects",
		"Projects": projects,
	}, "layout")
}

// ListProjectsAPI returns all projects as JSON.
func (h *Handler) ListProjectsAPI(c *fiber.Ctx) error {
	var projects []models.Project
	if err := h.DB.Order("display_order asc, created_at desc").Find(&projects).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to load projects")
	}
	return c.JSON(projects)
}

// Create/Update/Delete handlers for projects live behind the admin API,
// which isn't built yet — see JWT_ADMIN_API_GUIDE.md.
