package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type animal struct {
	gorm.Model
	Animaltype string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT`
	Zone       int    `gorm:type:INTEGER`
	Age        int    `gorm:"type:INTEGER`
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1  port=5432 user=mxb5594 dbname=project-dino sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.DropTableIfExists(&animal{})
	db.AutoMigrate(&animal{})
	a := animal{
		Animaltype: "Tyrannosaurus Rex",
		Nickname:   "rex",
		Zone:       1,
		Age:        11,
	}

	db.Save(&a)

	a = animal{
		Animaltype: "Velociraptor",
		Nickname:   "rapto",
		Zone:       2,
		Age:        11,
	}
	db.Create(&a)

	animals := []animal{}

	db.Find(&animals)
	fmt.Println(animals)

}
