package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute() *RestRoute {
	return NewRestRoute().
		SetRoot().
		Handler(func(router fiber.Router) {
			router.Get("/swagger/swagger.json", func(c *fiber.Ctx) error {
				c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
				return c.SendFile("./docs/swagger.json")
			})
			router.Get("/swagger/*", swagger.New(swagger.Config{
				URL:         "/swagger/swagger.json",
				DeepLinking: true,
			}))

		})
}
