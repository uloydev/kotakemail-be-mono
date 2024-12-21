package cmd

type Command interface {
	Name() string
	Execute() error
	Shutdown() error
}
