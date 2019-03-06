package databaselayer

import (
	"database/sql"
	// Use the mysql driver.
	_ "github.com/go-sql-driver/mysql"
)

// MySQLHandler embeds the SQLHandler type.
type MySQLHandler struct {
	*SQLHandler
}

// NewMySQLHandler constructor function returns a pointer to the MySQLHandler type.
func NewMySQLHandler(connection string) (*MySQLHandler, error) {
	db, err := sql.Open("mysql", connection)
	return &MySQLHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
