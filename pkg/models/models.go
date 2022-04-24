package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Setup is used to set up & created the needed tables in the Postgres database with the specified models.
func Setup(db *pg.DB) {

	models := []interface{}{
		(*User)(nil),
		(*Clip)(nil),
	}

	for _,model := range models {
		if err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		}); err != nil {
			panic(err)
		}
	}
}