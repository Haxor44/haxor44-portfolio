package handlers

import (
	"github.com/gofiber/fiber/v2"

	"portfolio/internal/models"
)

// Home renders the landing page: bio/headline plus a handful of featured
// projects, the latest articles, and certifications, to keep the page fast
// and avoid loading every record on first paint.
func (h *Handler) Home(c *fiber.Ctx) error {
	var featured []models.Project
	h.DB.Where("featured = ?", true).Order("display_order asc").Limit(3).Find(&featured)

	var certs []models.Certification
	h.DB.Where("status = ?", models.CertEarned).
		Order("issue_date desc, display_order asc").Limit(4).Find(&certs)

	var inProgress []models.Certification
	h.DB.Where("status = ?", models.CertInProgress).Order("display_order asc").Find(&inProgress)

	var articles []models.Article
	h.DB.Order("published_at desc").Limit(3).Find(&articles)

	return c.Render("index", fiber.Map{
		"Title":            "Home",
		"FeaturedProjects": featured,
		"RecentCerts":      certs,
		"InProgressCerts":  inProgress,
		"RecentArticles":   articles,
	}, "layout")
}
