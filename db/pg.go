package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connStr := "user=mxb5594 dbname=project-dino sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	errorHandler(err)
	defer db.Close()

	stmt, err := db.Prepare(" SELECT * FROM animals WHERE age > $1 ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(5)
	animals := handleRows(rows, err)
	fmt.Println(animals)

	rows, err = stmt.Query(10)
	animals = handleRows(rows, err)
	fmt.Println(animals)

}

func handleRows(rows *sql.Rows, err error) (animals []animal) {
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.age, &a.zone)

		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return animals
}
