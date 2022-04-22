package controllers

import (
	"clips/internal/config"
	"clips/internal/data"
	"clips/internal/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Signup(c *fiber.Ctx) error {

	if !config.AllowSignup {
		return c.Status(401).SendString("Signup is disabled")
	}

	sess,err := data.Store.Get(c)

	if err != nil {
		log.Printf("Failed to get session: %s",err)

		return c.Status(500).SendString("Unexpected Error")
	}

	signup := data.Login{}

	if err := c.BodyParser(&signup); err != nil {
		return c.Status(500).SendString("Unexpected Error")
	}

	res,err := services.Signup(signup)

	if err != nil {
		log.Printf("Failed to signup: %s",err)

		return c.Status(500).SendString("Unexpected Error")
	}

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

