package handlers

import (
	"github.com/gofiber/fiber/v2"

	"portfolio/internal/models"
)

// CertificationsPage renders the certifications section, split into
// earned certifications and the ones currently in progress.
func (h *Handler) CertificationsPage(c *fiber.Ctx) error {
	var earned []models.Certification
	if err := h.DB.Where("status = ?", models.CertEarned).
		Order("issue_date desc, display_order asc").Find(&earned).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to load certifications")
	}

	var inProgress []models.Certification
	if err := h.DB.Where("status = ?", models.CertInProgress).
		Order("display_order asc").Find(&inProgress).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to load certifications")
	}

	return c.Render("certifications", fiber.Map{
		"Title":      "Certifications",
		"Earned":     earned,
		"InProgress": inProgress,
	}, "layout")
}

// Create/Delete handlers for certifications live behind the admin API,
// which isn't built yet — see JWT_ADMIN_API_GUIDE.md.
