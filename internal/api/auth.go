package api

import (
	"clips/internal/data"
	"clips/internal/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Auth is used to set the authenticated user for every request that hits /api/*
func Auth(c *fiber.Ctx) error {
	sess,err := data.Store.Get(c)

	if err != nil {
		log.Printf("Failed to get session: %s",err)

		return c.Status(500).JSON(fiber.Map{
			"error":"Unknown Error",
		})
	}

	if sess.Get("active") == true {
		c.Locals("active",true)
		c.Locals("user",sess.Get("user"))
	}

	// If there is an API key in the header, then set the authenticated user with the specified API key

	apiKey := c.GetReqHeaders()["X-Api-Key"]

	if len(apiKey) != 0 {
		if res := services.APIKeyAuth(apiKey); res.Success {
			c.Locals("active",true)
			c.Locals("user",res.User)
		}
	}

	return c.Next()
}