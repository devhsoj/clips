package data

import (
	"clips/pkg/models"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

// Setup registers the needed types to the fiber middleware session.
func Setup() {
	Store.RegisterType(models.User{})
}