package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/omarghandour/test1/db"
	"github.com/omarghandour/test1/routes"
)

// @title           Books CRUD API
// @version         1.0
// @description     A sample RESTful CRUD API built with Go, Fiber v3, SQLite, and GORM.
// @host            localhost:8080
// @BasePath        /
func main() {
	// 1. Initialize SQLite Database
	db.InitDB()

	// 2. Initialize Fiber app
	app := fiber.New()

	// Middleware: logs every HTTP request to console
	app.Use(logger.New())

	// 3. Register Routes (including Swagger UI)
	routes.SetupRoutes(app)

	// 4. Start Server on port 8080
	log.Println("Server is running on http://localhost:8080")
	log.Println("Swagger UI is available at http://localhost:8080/swagger/")
	log.Fatal(app.Listen(":8080"))
}
