package app

import (
	"clips/pkg/config"
	"clips/pkg/data"
	"clips/pkg/db"
	"clips/pkg/models"
	"clips/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
)

func Setup() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %s",err)
	}

	db.Setup()
	data.Setup()
	models.Setup(db.Database)
}

func Start() {
	data.Application = fiber.New(fiber.Config{
		BodyLimit: config.BodyLimit,
		DisableStartupMessage: true,
	})

	// Setup session store
	data.Store = session.New()

	// Setup static router for serving svelte bundles and other static files
	data.Application.Static("/","./public",fiber.Static{Compress: true})

	// Setup other routes
	routes.SetupAPIRoutes()
	routes.SetupAuthRoutes()

	var err error

	if config.HTTPS {
		log.Printf("Starting on https://%s",config.ListenAddress)
		err = data.Application.ListenTLS(config.ListenAddress,config.TLSCertPath,config.TLSKeyPath)
	} else {
		log.Printf("Starting on http://%s",config.ListenAddress)
		err = data.Application.Listen(config.ListenAddress)
	}

	if err != nil {
		log.Fatalf("Failed to start server: %s",err)

		db.Database.Close()
	}
}