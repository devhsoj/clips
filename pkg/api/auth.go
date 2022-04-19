package api

import (
	"clips/pkg/data"
	"clips/pkg/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Auth(c *fiber.Ctx) error {
	sess,err := data.Store.Get(c)

	if err != nil {
		log.Printf("Failed to get session: %s",err)
	}

	apiKey := c.GetReqHeaders()["Authorization"]

	// We keep these two auth checks separate so one doesn't get called if the other passes

	if sess.Get("active") == true {
		c.Locals("user",sess.Get("user"))

		return c.Next()
	}

	if res := services.APIKeyAuth(apiKey); res.Success {
		c.Locals("user",res.User)

		return c.Next()
	}

	return c.SendStatus(401)
}