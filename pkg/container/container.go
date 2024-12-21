package container

import (
	"kotakemail.id/pkg/cmd"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/database"
	"kotakemail.id/pkg/logger"
)

type Container[T any] struct {
	commands  []cmd.Command
	ctx       *appcontext.AppContext
	cfg       T
	databases map[string]database.Database
	logger    *logger.Logger
}

func NewContainer[T any](ctx *appcontext.AppContext, cfg T, logger *logger.Logger) *Container[T] {
	return &Container[T]{
		ctx:    ctx,
		cfg:    cfg,
		logger: logger,
	}
}

func (c *Container[T]) AddCommand(command cmd.Command) {
	c.commands = append(c.commands, command)
}

func (c *Container[T]) AddDatabase(db database.Database) {
	if c.databases == nil {
		c.databases = make(map[string]database.Database)
	}
	c.databases[db.Name()] = db
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
	for _, command := range c.commands {
		command.Shutdown()
	}
	c.logger.Shutdown()
}
