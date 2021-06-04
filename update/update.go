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
		skyorm.Eq(models.MStore.Pk(), 2),
		skyorm.Eq(m.UintValueProp(), 4),
	)
	values := []skyorm.Val{
		skyorm.NewVal(m.ValueProp(), "new value"),
	}
	if err = db.Update(models.MStore, cond, values...); err != nil {
		log.Fatal(err)
	}
}
