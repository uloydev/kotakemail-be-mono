package container

import (
	"kotakemail.id/pkg/cmd"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/database"
	"kotakemail.id/pkg/logger"
	"kotakemail.id/pkg/storage"
)

type Container[T any] struct {
	commands  []cmd.Command
	ctx       *appcontext.AppContext
	cfg       T
	databases map[string]database.Database
	storages  map[string]storage.Storage
	logger    *logger.Logger
}

func NewContainer[T any](ctx *appcontext.AppContext, cfg T, logger *logger.Logger) *Container[T] {
	return &Container[T]{
		ctx:       ctx,
		cfg:       cfg,
		logger:    logger,
		databases: make(map[string]database.Database),
		storages:  make(map[string]storage.Storage),
	}
}

func (c *Container[T]) AddCommand(command ...cmd.Command) {
	c.commands = append(c.commands, command...)
}

func (c *Container[T]) AddDatabase(db ...database.Database) {
	for _, d := range db {
		c.databases[d.Name()] = d
	}
}

func (c *Container[T]) AddStorage(storage ...storage.Storage) {
	for _, s := range storage {
		c.storages[s.Name()] = s
	}
}

func (c *Container[T]) GetDatabase(name string) database.Database {
	return c.databases[name]
}

func (c *Container[T]) Run() {
	for _, command := range c.commands {
		command.Execute()
	}
}

func (c *Container[T]) Context() *appcontext.AppContext {
	return c.ctx
}

func (c *Container[T]) Config() T {
	return c.cfg
}

func (c *Container[T]) Logger() *logger.Logger {
	return c.logger
}

func (c *Container[T]) Shutdown() {
	c.logger.Info().Msg("Shutting down app")
	for _, command := range c.commands {
		c.logger.Info().Msgf("Shutting down command %s", command.Name())
		if err := command.Shutdown(); err != nil {
			c.logger.Error().Err(err).Msgf("Failed to shutdown command %s", command.Name())
		}
	}
	for _, db := range c.databases {
		c.logger.Info().Msgf("Shutting down database %s", db.Name())
		if err := db.Shutdown(); err != nil {
			c.logger.Error().Err(err).Msgf("Failed to shutdown database %s", db.Name())
		}
	}

	for _, storage := range c.storages {
		c.logger.Info().Msgf("Shutting down storage %s", storage.Name())
		if err := storage.Shutdown(); err != nil {
			c.logger.Error().Err(err).Msgf("Failed to shutdown storage %s", storage.Name())
		}
	}

	c.logger.Info().Msg("Finished shutting down app")
}
