package routes

import (
	"github.com/gofiber/fiber/v2"
	"kotakemail.id/shared/base/rest"
)

// HealthCheck godoc
// @Summary      health check
// @Description  check if the service is up and running
// @Tags         misc
// @Produce      json
// @Success      200 {string} string "ok"
// @Router       /health [get]

func HeathCheckRoute() *rest.RestRoute {
	route := rest.NewRestRoute().SetRoot()
	route.Handler(func(r fiber.Router) {
		r.Get("/health", func(c *fiber.Ctx) error {
			return c.SendString("ok")
		})
	})

	return route
}
