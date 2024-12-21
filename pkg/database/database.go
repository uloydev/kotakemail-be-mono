package database

type BaseDatabase struct {
	name string
}

func (d *BaseDatabase) Name() string {
	return d.name
}

func (d *BaseDatabase) SetName(name string) {
	d.name = name
}
