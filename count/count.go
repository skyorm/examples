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
	cond := skyorm.Or(
		skyorm.Eq(m.OrmPkProp(), 2),
		skyorm.Eq(m.UintValueProp(), 4),
	)
	size, err := db.Count(models.MStore, cond)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(size)
}
