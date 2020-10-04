package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/webdelo/tradebot/server/routes"
)

func Run(ctx *context.Context) error {
	app := fiber.New()

	// define routes
	routes.Define(app)

	// start web-server
	err := app.Listen(":3000")
	if err != nil {
		return err
	}

	return nil
}
