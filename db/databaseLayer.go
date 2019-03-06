package databaselayer

import "errors"

// Specifies what type of database is being used.
const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

// DinoDBHandler interface provides various db interactions.
type DinoDBHandler interface {
	GetAvailableDynos() ([]Animal, error)
	GetDynoByNickname(string) (Animal, error)
	GetDynosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

// Animal struct represents values of each animal.
type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

// ErrDBNotSupported gives friendly error message if dialect is not supported.
var ErrDBNotSupported = errors.New("The Database type provided is not supported")

// GetDatabaseHandler is a factory function that returns a db handler object.
func GetDatabaseHandler(dbtype uint8, connection string) (DinoDBHandler, error) {
	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongoDBHandler(connection)
	case SQLITE:
		return NewSQLiteHandler(connection)
	case POSTGRESQL:
		return NewPGHandler(connection)
	}
	return nil, ErrDBNotSupported
}
