package api

import (
	"clips/pkg/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	return c.JSON(fiber.Map{
		"message":fmt.Sprintf("Authenticated as %s",user.Username),
	})
}