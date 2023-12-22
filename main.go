package main

import (
	"fmt"

	configuration "matsukana.cloud/go-marketing/config"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	*fiber.App

	DB *database.Database
}

func main() {
	config := configuration.New()

	app := App{
		App: fiber.New(*config.GetFiberConfig()),
	}

	// Initialize database
	db, err := database.New(&database.DatabaseConfig{
		Driver:   config.GetString("DB_DRIVER"),
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
		Database: config.GetString("DB_DATABASE"),
	})

	if err != nil || db == nil {
		fmt.Println("failed to connect to database:", err.Error())
	}

	app.registerMiddlewares()

	// Handle Register All Route in Router Folder
	api := app.Group("/api")
	router.RegisterRoute(api)

	// Start listening on the specified address
	err = app.Listen(config.GetString("APP_ADDR"))
	if err != nil {
		app.exit()
	}
}

func (app *App) registerMiddlewares() {
	// Handle Panic
	app.Use(recover.New())
	app.Use(logger.New())
}

// Stop the Fiber application
func (app *App) exit() {
	_ = app.Shutdown()
}
