package handlers

import "github.com/gofiber/fiber/v2"

// Health is a liveness/readiness probe target for Docker HEALTHCHECK,
// Kubernetes/ECS probes, and load balancer health checks.
func (h *Handler) Health(c *fiber.Ctx) error {
	sqlDB, err := h.DB.DB()
	if err != nil || sqlDB.Ping() != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":   "unhealthy",
			"database": "unreachable",
		})
	}
	return c.JSON(fiber.Map{"status": "ok"})
}
