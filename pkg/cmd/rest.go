package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"kotakemail.id/config"
	"kotakemail.id/pkg/logger"
	"kotakemail.id/pkg/rest"
)

type RestCommand struct {
	name        string
	cfg         *config.Config
	logger      *logger.Logger
	app         *fiber.App
	routes      []*rest.RestRoute
	middlewares []fiber.Handler
}

func NewRestCommand(
	cfg *config.Config,
	logger *logger.Logger,
	routes []*rest.RestRoute,
	middlewares ...fiber.Handler,

) Command {
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})
	return &RestCommand{
		name:        cfg.AppName + "-rest",
		cfg:         cfg,
		logger:      logger,
		app:         app,
		routes:      routes,
		middlewares: middlewares,
	}
}

func (r *RestCommand) Execute() error {
	r.logger.Info().Msg("starting rest server")
	router := r.app.Group(r.cfg.Rest.BasePath)

	for _, middleware := range r.middlewares {
		router.Use(middleware)
	}

	for _, route := range r.routes {
		route.Register(router)
	}

	return r.app.Listen(fmt.Sprintf("%s:%s", r.cfg.Rest.Host, r.cfg.Rest.Port))
}

func (r *RestCommand) Shutdown() error {
	r.logger.Info().Msg("shutting down rest server")
	return r.app.Shutdown()
}

func (r *RestCommand) App() interface{} {
	return r.app
}

func (r *RestCommand) Name() string {
	return r.name
}
