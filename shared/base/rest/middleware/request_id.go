package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// func to return fiber handler to set X-REQUEST-ID to response header
func RequestIDMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {
		c.Set("X-REQUEST-ID", uuid.Must(uuid.NewV7()).String())
		return c.Next()
	}

}
