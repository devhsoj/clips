package db

import (
	"clips/pkg/config"
	"github.com/go-pg/pg/v10"
	"log"
)

var Database *pg.DB

func Setup() {
	var err error
	var options *pg.Options

	options,err = pg.ParseURL(config.PostgresURL)

	if err != nil {
		log.Printf("Failed to parse PG Connection URL: %s",err)
	}

	Database = pg.Connect(options)
}