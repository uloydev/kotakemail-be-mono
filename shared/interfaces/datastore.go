package interfaces

type Datastore[T any] interface {
	GetConnection() T
	Name() string
	SetName(name string)
	Shutdown() error
}
