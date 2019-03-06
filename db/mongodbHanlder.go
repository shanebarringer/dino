package databaselayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDBHandler embeds mgo.Session.
type MongoDBHandler struct {
	*mgo.Session
}

// NewMongoDBHandler constructor takes a connection and returns a pointer to the MongoDBHandler.
func NewMongoDBHandler(connection string) (*MongoDBHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongoDBHandler{
		Session: s,
	}, err
}

// GetAvailableDynos returns all animals.
func (handler *MongoDBHandler) GetAvailableDynos() ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(nil).All(&animals)
	return animals, err
}

// GetDynoByNickname returns an animal based on nickname.
func (handler *MongoDBHandler) GetDynoByNickname(nickname string) (Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	a := Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"nickname": nickname}).One(&a)
	return a, err
}

// GetDynosByType returns all dinos that match the specified animal_type.
func (handler *MongoDBHandler) GetDynosByType(dinoType string) ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"animal_type": dinoType}).All(&animals)
	return animals, err
}

// AddAnimal adds a new record to the database.
func (handler *MongoDBHandler) AddAnimal(a Animal) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("Dino").C("animals").Insert(a)
}

// UpdateAnimal updates an existing record in the db.
func (handler *MongoDBHandler) UpdateAnimal(a Animal, nickname string) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("Dino").C("animals").Update(bson.M{"nickname": nickname}, a)
}

func (handler *MongoDBHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}
