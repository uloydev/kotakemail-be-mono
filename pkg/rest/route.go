package rest

import "github.com/gofiber/fiber/v2"

type RestRoute struct {
	handle func(router fiber.Router)
	name   string
	prefix string
}

func NewRestRoute(name, prefix string, handle func(router fiber.Router)) *RestRoute {
	return &RestRoute{
		handle: handle,
		name:   name,
		prefix: prefix,
	}
}

func (r *RestRoute) Name() string {
	return r.name
}

func (r *RestRoute) Prefix() string {
	return r.prefix
}

func (r *RestRoute) Register(router fiber.Router) {
	r.handle(router)
}
