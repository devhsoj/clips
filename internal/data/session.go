package data

import (
	"clips/pkg/models"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func Setup() {
	Store.RegisterType(models.User{})
}