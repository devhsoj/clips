package db

import (
	"clips/internal/config"
	"github.com/go-pg/pg/v10"
	"log"
)

var Database *pg.DB

// Setup tries to connect using the configured PostgreSQL connection URL.
func Setup() {
	options,err := pg.ParseURL(config.PostgresURL)

	if err != nil {
		log.Printf("Failed to parse PG Connection URL: %s",err)
	}

	Database = pg.Connect(options)
}