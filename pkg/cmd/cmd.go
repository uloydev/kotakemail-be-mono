package cmd

type Command interface {
	Execute()
	Shutdown()
}
