package datastore

type BaseDatastore struct {
	name string
}

func (d *BaseDatastore) Name() string {
	return d.name
}

func (d *BaseDatastore) SetName(name string) {
	d.name = name
}
