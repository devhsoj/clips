package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

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