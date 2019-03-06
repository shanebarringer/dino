package databaselayer

import (
	"database/sql"
	// import sqlite driver for NewSQLiteHandler function.
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteHandler embeds the SQLHandler.
type SQLiteHandler struct {
	*SQLHandler
}

// NewSQLiteHandler constructor function takes a connection string and returns a pointer to the SQLiteHandler.
func NewSQLiteHandler(connection string) (*SQLiteHandler, error) {
	db, err := sql.Open("sqlite3", connection)
	return &SQLiteHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err

}
