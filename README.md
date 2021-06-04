# examples

This repository contains examples for using [skyorm](https://github.com/skyorm/skyorm).

## Packages

package|Description
---|---
[models](models)|Models definitions, used in examples.
[count](count/count.go)|Count method example
[delete](delete/delete.go)|Delete method example.
[find](find/find.go)|Find method example.
[populate](populate/populate.go)|Populate method example.
[put](put/put.go)|Put method example.
[update](update/update.go)|Update method example.

## Basic sample

```go
package main

import (
	`log`
	"os"

	"github.com/skyorm/examples/models"
	`github.com/skyorm/postgres`
	`github.com/skyorm/skyorm`
)

func main() {
	p, err := postgres.New(os.Getenv("DSN"), log.Default())
	if err != nil {
		log.Fatal(err)
	}
	db := skyorm.NewDB(p)

	m := &models.M{}
	cond := skyorm.Or(
		skyorm.Eq(m.OrmPkProp(), 2),
		skyorm.Eq(m.UintValueProp(), 4),
	)
	list, err := db.Find(models.MStore, cond, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range list {
		model := v.(*models.M)
		log.Println(*model)
	}
}
```
