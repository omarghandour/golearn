package main

import (
	"log"
	"os"

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

	// 4. Determine Port (configurable via env, default 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 5. Start Server
	log.Printf("Server is running on port %s\n", port)
	log.Printf("Swagger UI is available at http://localhost:%s/swagger/\n", port)
	log.Fatal(app.Listen(":" + port))
}
