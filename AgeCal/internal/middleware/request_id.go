package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Check if client already sent request id
		reqID := c.Get("X-Request-ID")

		// 2. Generate if missing
		if reqID == "" {
			reqID = uuid.NewString()
		}

		// 3. Store in context
		c.Locals(RequestIDKey, reqID)

		// 4. Send back in response header
		c.Set("X-Request-ID", reqID)

		return c.Next()
	}
}
