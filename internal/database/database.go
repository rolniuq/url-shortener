package database

type Database interface {
	Connect() error
}

type database struct {
}

func NewDatabase() Database {
	return &database{}
}

func (d *database) Connect() error {
	return nil
}
