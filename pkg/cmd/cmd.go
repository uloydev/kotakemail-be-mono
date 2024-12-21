package cmd

type Command interface {
	App() interface{}
	Name() string
	Execute() error
	Shutdown() error
}
