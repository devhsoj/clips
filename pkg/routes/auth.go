package routes

import (
	"clips/pkg/controllers"
	"clips/pkg/data"
	"github.com/gofiber/fiber/v2"
)

var AuthRouter fiber.Router

func SetupAuthRoutes() {
	AuthRouter = data.Application.Group("/auth")

	AuthRouter.Post("/login",controllers.Login)
	AuthRouter.Post("/signup",controllers.Signup)
}