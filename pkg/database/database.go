package database

type Database interface {
	Name() string
	SetName(name string)
	GetConnection() interface{}
	Shutdown() error
}

type BaseDatabase struct {
	name string
}

func (d *BaseDatabase) Name() string {
	return d.name
}

func (d *BaseDatabase) SetName(name string) {
	d.name = name
}
