package routes

import (
	"clips/internal/controllers"
	"clips/internal/data"
	"github.com/gofiber/fiber/v2"
)

var AuthRouter fiber.Router

// SetupAuthRoutes sets up all routing under /auth.
func SetupAuthRoutes() {
	AuthRouter = data.Application.Group("/auth")

	AuthRouter.Post("/login",controllers.Login)
	AuthRouter.Post("/signup",controllers.Signup)
}