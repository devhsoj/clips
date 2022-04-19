package controllers

import (
	"clips/pkg/data"
	"clips/pkg/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Signup(c *fiber.Ctx) error {

	sess,err := data.Store.Get(c)

	if err != nil {
		log.Printf("Failed to get session: %s",err)
		return c.SendStatus(500)
	}

	signup := data.Login{}

	if err := c.BodyParser(&signup); err != nil {
		return c.SendStatus(500)
	}

	res,err := services.Signup(signup)

	if err != nil {
		log.Printf("Failed to signup: %s",err)
		return c.SendStatus(500)
	}

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

