package handlers

import (
	"github.com/gofiber/fiber/v2"

	"portfolio/internal/models"
)

// BlogPage renders the articles/blog listing, linking out to Medium/dev.to
// or to a self-hosted article detail page.
func (h *Handler) BlogPage(c *fiber.Ctx) error {
	var articles []models.Article
	if err := h.DB.Order("published_at desc").Find(&articles).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to load articles")
	}
	return c.Render("blog", fiber.Map{
		"Title":    "Blog",
		"Articles": articles,
	}, "layout")
}

// Create/Delete handlers for articles live behind the admin API,
// which isn't built yet — see JWT_ADMIN_API_GUIDE.md.
