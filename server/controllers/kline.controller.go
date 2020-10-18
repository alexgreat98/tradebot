package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container"
	binance2 "github.com/webdelo/tradebot/pkg/binance/binance"
)

func KlineIndex(c *fiber.Ctx) error {
	var BinanceKlineRepo binance2.BinanceKlineRepository
	container.Make(&BinanceKlineRepo)
	return c.JSON(BinanceKlineRepo.Find())
}
