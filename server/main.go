package main

import (
	"github.com/TimothySubekti0322/first-web-with-golang-playground/database"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
app := fiber.New()

// CORS middleware configuration
app.Use(cors.New(cors.Config{
	AllowOrigins:     "http://localhost:3000", // Specify your frontend origin here
	AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	AllowCredentials: true,
}))

// Initialize Database
database.DatabaseInit()

// Initialize Migration
// migrations.Migration()

// Initialize Routes

app.Get("/", func(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message" : "Welcome to backend Golang API",
	})
})

routes.RouteInit(app)

app.Listen(":8080");

}