package main

import (
	`log`
	"strconv"

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

	// insert 10 records into
	for i := 0; i < 10; i++ {
		m := &models.M{Value: "value" + strconv.Itoa(i), UintValue: uint64(i)}
		if err = db.Put(m); err != nil {
			log.Fatal(err)
		}
		log.Println(*m)
	}
}
