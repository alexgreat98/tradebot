package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/webdelo/tradebot/server/routes"
)

func Run(ctx *context.Context) error {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,X-Requested-With, Content-Type, Accept",
	}))
	// define routes
	routes.Define(app)

	// start web-server
	err := app.Listen(":3000")
	if err != nil {
		return err
	}

	return nil
}
