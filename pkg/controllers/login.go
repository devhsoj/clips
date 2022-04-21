package controllers

import (
	"clips/pkg/data"
	"clips/pkg/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

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
	}

	return c.Status(res.Status).SendString(res.Message)
}
