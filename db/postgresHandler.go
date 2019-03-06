package databaselayer

import (
	"database/sql"
	// import pq db driver.
	_ "github.com/lib/pq"
)

// PGHandler embeds SQLHandler
type PGHandler struct {
	*SQLHandler
}

// NewPGHandler constructor function takes a connection string and returns a pointer to the PGHandler.
func NewPGHandler(connection string) (*PGHandler, error) {
	db, err := sql.Open("postgres", connection)

	return &PGHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
