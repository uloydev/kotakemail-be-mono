package routes

import (
	"github.com/gofiber/fiber/v2"
	"kotakemail.id/pkg/rest"
)

func HeathCheckRoute() *rest.RestRoute {
	return rest.NewRestRoute("health check", "", func(router fiber.Router) {
		router.Get("health", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"status": "up",
			})
		})
	})
}
