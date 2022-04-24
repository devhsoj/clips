package controllers

import (
	"clips/internal/data"
	"clips/internal/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Login is called whenever a user hits POST /auth/login. It's passed the credentials from the form in Login.svelte, and
// tries to authenticate.
func Login(c *fiber.Ctx) error {

	sess,err := data.Store.Get(c)

	if err != nil {
		log.Printf("Failed to get session: %s",err)

		return c.Status(500).SendString("Unexpected Error")
	}

	login := data.Login{}

	if err := c.BodyParser(&login); err != nil {
		log.Printf("Failed to parse login body: %s",err)

		return c.Status(500).SendString("Unexpected Error")
	}

	res := services.UserAuth(login)

	if !res.Success {
		return c.Status(res.Status).SendString(res.Message)
	}

	sess.Set("user",res.User)
	sess.Set("active",true)

	if err = sess.Save(); err != nil {
		log.Printf("Failed to save session: %s",err)

		return c.Status(500).SendString("Unexpected Error")
	}

	return c.Status(res.Status).SendString(res.Message)
}
