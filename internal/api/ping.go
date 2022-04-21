package api

import (
	"clips/pkg/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	return c.Status(200).SendString(fmt.Sprintf("Authenticated as %s",user.Username))
}