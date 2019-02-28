package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("mysql", "root@/dino")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from dino.animals where age > ?", 10)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	animals := []animal{}

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

	fmt.Println(animals)

	row := db.QueryRow("select * from dino.animals where age > ?", 10)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.age, &a.zone)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)
}
