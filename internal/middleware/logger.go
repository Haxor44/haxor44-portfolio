package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestID attaches a unique ID to every request (as a response header and
// a Locals value) so log lines and error responses can be correlated back
// to one request. Actual request logging is handled by Fiber's built-in
// logger middleware (see cmd/server/main.go), which reads this ID.
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := uuid.NewString()
		c.Locals("request_id", id)
		c.Set("X-Request-ID", id)
		return c.Next()
	}
}
