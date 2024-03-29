package main

import (
	"os"
	"os/signal"

	cache "matsukana.cloud/go-marketing/cache/redis"
	configuration "matsukana.cloud/go-marketing/config"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/message"
	"matsukana.cloud/go-marketing/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	*fiber.App
	Config *configuration.Config
	Db     *database.Database
	Redis  *cache.RedisCache
}

func NewApp(config *configuration.Config, Db *database.Database, Redis *cache.RedisCache) *App {
	return &App{fiber.New(*config.GetFiberConfig()), config, Db, Redis}
}

func main() {
	app, err := InitializedServer()

	if err != nil {
		panic(err.Error())
	}

	app.registerMiddlewares()

	mq := message.NewMessage(app.Config)

	if mq != nil {
		if mq.Connection != nil {
			defer mq.Connection.Close()
		}

		if mq.Channel != nil {
			defer mq.Channel.Close()
		}
	}

	go func() {
		mq.Consumer()
	}()

	// Handle Register All Route in Router Folder
	appRouter := InitializedRouter(app.Db, app.Config)
	app.Mount("/api", appRouter)

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		if err := c.SendStatus(fiber.StatusNotFound); err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "NOK",
			Data:   nil,
		})
	})

	// Close any connections on interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		app.exit()
	}()

	// Start listening on the specified address
	err = app.Listen(app.Config.GetString("APP_ADDR"))
	if err != nil {
		app.exit()
	}
}

func (app *App) registerMiddlewares() {
	// Handle Panic
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())
}

// Stop the Fiber application
func (app *App) exit() {
	_ = app.Shutdown()
}
