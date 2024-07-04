package routes

import (
	"github.com/TimothySubekti0322/first-web-with-golang-playground/controllers"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	// User Routes
	user := r.Group("/users", middleware.JWTMiddleware)

	user.Get("/", controllers.GetAllUsers)
	user.Post("/", controllers.Register)

	auth := r.Group("/auth")

	auth.Post("/", controllers.Login)
}