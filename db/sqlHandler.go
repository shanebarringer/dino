// Package databaselayer manages multiple databases.
package databaselayer

import (
	"database/sql"
	"fmt"
	"log"
)

// SQLHandler contains database information.
type SQLHandler struct {
	*sql.DB
}

// GetAvailableDynos returns all animals.
func (handler *SQLHandler) GetAvailableDynos() ([]Animal, error) {
	return handler.sendQuery("SELECT * FROM animals")
}

// GetDynoByNickname returns animal based on nickname.
func (handler *SQLHandler) GetDynoByNickname(nickname string) (Animal, error) {
	row := handler.QueryRow(fmt.Sprintf("SELECT * FROM aninals WHERE nickname = '%s'", nickname))

	a := Animal{}
	err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
	return a, err
}

// GetDynosByType returns all animals matching a specific animal_type.
func (handler *SQLHandler) GetDynosByType(dinoType string) ([]Animal, error) {
	return handler.sendQuery(fmt.Sprintf("SELECT * FROM animals WHERE animal_type = '%s'", dinoType))
}

// AddAnimal creates a new record in the database.
func (handler *SQLHandler) AddAnimal(a Animal) error {
	_, err := handler.Exec(fmt.Sprintf("INSERT INTO animals (animal_type,nickname,zone,age) VALUES ('%s', '%s', '%d', '%d')", a.AnimalType, a.Nickname, a.Zone, a.Age))
	return err
}

// UpdateAnimal will update the properties of the specified animal.
func (handler *SQLHandler) UpdateAnimal(a Animal, nickname string) error {
	_, err := handler.Exec("UPDATE animals SET animal_type = '%s', nickname = '%s', zone = '%d', age = '%d' WHERE nickname = '%s'", a.AnimalType, a.Nickname, a.Zone, a.Age, nickname)
	return err
}

func (handler *SQLHandler) sendQuery(q string) ([]Animal, error) {
	Animals := []Animal{}
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		a := Animal{}
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
		}
		Animals = append(Animals, a)
	}
	return Animals, rows.Err()
}
