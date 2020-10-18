package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/webdelo/tradebot/server/controllers"
)

// Function define all web-routes
func Define(app *fiber.App) {
	app.Get("/", controllers.IndexPage)
	app.Get("/kline", controllers.KlineIndex)
}
