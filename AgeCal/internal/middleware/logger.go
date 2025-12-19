package middleware

import (
	"time"

	"AgeCal/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ZapLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)
		reqID, _ := c.Locals(RequestIDKey).(string)

		if logger.Log != nil {
			logger.Log.Info("request",
				zap.String("request_id", reqID),
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.Int("status", c.Response().StatusCode()),
				zap.Duration("duration", duration),
			)
		}
		return err
	}
}
