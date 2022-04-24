package routes

import (
	"clips/internal/api"
	"clips/internal/data"
	"github.com/gofiber/fiber/v2"
)

var APIRouter fiber.Router
var UserRouter fiber.Router
var ClipRouter fiber.Router

// SetupAPIRoutes sets up all routing under /api.
func SetupAPIRoutes() {
	// API index routing
	APIRouter = data.Application.Group("/api",api.Auth)

	// User API routing
	UserRouter = APIRouter.Group("/user")

	UserRouter.Get("/me",api.GetMe)
	UserRouter.Get("/:user_id",api.GetUser)

	// Clip API routing
	ClipRouter = APIRouter.Group("/clip")

	ClipRouter.Post("/new",api.NewClip)
	ClipRouter.Get("/get",api.GetClips)
	ClipRouter.Post("/views/:clip_id",api.IncrementViews)
	ClipRouter.Get("/view/:clip_id",api.ViewClip)
	ClipRouter.Get("/:clip_id",api.GetClip)
}