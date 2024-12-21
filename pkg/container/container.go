package container

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"kotakemail.id/pkg/cmd"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/database"
	"kotakemail.id/pkg/logger"
	"kotakemail.id/pkg/storage" // Ensure this import path is correct
)

type Container struct {
	commands  []cmd.Command
	ctx       *appcontext.AppContext
	databases map[string]database.Database
	storages  map[string]storage.Storage
	logger    *logger.Logger
}

func NewContainer(ctx *appcontext.AppContext, logger *logger.Logger) *Container {
	return &Container{
		ctx:       ctx,
		logger:    logger,
		databases: make(map[string]database.Database),
		storages:  make(map[string]storage.Storage),
	}
}

func (c *Container) AddCommand(command ...cmd.Command) {
	c.commands = append(c.commands, command...)
}

func (c *Container) AddDatabase(db ...database.Database) {
	for _, d := range db {
		c.databases[d.Name()] = d
	}
}

func (c *Container) AddStorage(storage ...storage.Storage) {
	for _, s := range storage {
		c.storages[s.Name()] = s
	}
}

func (c *Container) GetDatabase(name string) database.Database {
	return c.databases[name]
}

func (c *Container) Context() *appcontext.AppContext {
	return c.ctx
}

func (c *Container) Logger() *logger.Logger {
	return c.logger
}

func (c *Container) GetStorage(name string) storage.Storage {
	return c.storages[name]
}

func (c *Container) Run() {
	var wg sync.WaitGroup
	stop := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

	for _, command := range c.commands {
		wg.Add(1)
		go func(cmd cmd.Command) {
			defer wg.Done()
			if err := cmd.Execute(); err != nil {
				c.logger.Error().Err(err).Msgf("Failed to execute command %s", cmd.Name())
			}
		}(command)
	}

	wg.Add(1)
	go func() {
		<-sigint // Wait for a signal to stop
		close(stop)
		c.Shutdown()
		wg.Done()
	}()

	wg.Wait()
}

func (c *Container) Shutdown() {
	c.logger.Info().Msg("Shutting down app")

	var wg sync.WaitGroup

	for _, command := range c.commands {
		wg.Add(1)
		go func(cmd cmd.Command) {
			defer wg.Done()
			c.logger.Info().Msgf("Shutting down command %s", cmd.Name())
			if err := cmd.Shutdown(); err != nil {
				c.logger.Error().Err(err).Msgf("Failed to shutdown command %s", cmd.Name())
			}
		}(command)
	}

	for _, db := range c.databases {
		wg.Add(1)
		go func(db database.Database) {
			defer wg.Done()
			c.logger.Info().Msgf("Shutting down database %s", db.Name())
			if err := db.Shutdown(); err != nil {
				c.logger.Error().Err(err).Msgf("Failed to shutdown database %s", db.Name())
			}
		}(db)
	}

	for _, storage := range c.storages {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.logger.Info().Msgf("Shutting down storage %s", storage.Name())
			if err := storage.Shutdown(); err != nil {
				c.logger.Error().Err(err).Msgf("Failed to shutdown storage %s", storage.Name())
			}
		}()
	}

	wg.Wait()
	c.logger.Info().Msg("Finished shutting down app")
}
