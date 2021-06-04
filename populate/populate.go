package main

import (
	`log`

	"github.com/skyorm/examples/models"
	`github.com/skyorm/postgres`
	`github.com/skyorm/skyorm`
)

func main() {
	p, err := postgres.New(
		"postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable",
		log.Default())
	if err != nil {
		log.Fatal(err)
	}
	db := skyorm.NewDB(p)

	m := &models.M{}
	if err = db.Populate(m, 2); err != nil {
		if db.IsNotFound(err) {
			log.Printf("not found")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println(*m)
	}
}
