package api

import (
	"clips/internal/data"
	"clips/internal/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Auth is used to authenticate every request that hits /api/*
func Auth(c *fiber.Ctx) error {
	sess,err := data.Store.Get(c)

	if err != nil {
		log.Printf("Failed to get session: %s",err)

		return c.Status(500).JSON(fiber.Map{
			"error":"Unknown Error",
		})
	}

	if sess.Get("active") == true {
		c.Locals("user",sess.Get("user"))

		return c.Next()
	}

	// We keep these auth checks separate so one doesn't get called if the other passes

	apiKey := c.GetReqHeaders()["X-Api-Key"]

	if len(apiKey) == 0 {
		return c.SendStatus(401)
	}

	if res := services.APIKeyAuth(apiKey); res.Success {
		c.Locals("user",res.User)

		return c.Next()
	}

	return c.SendStatus(401)
}