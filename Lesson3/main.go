package main

import (
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func getHandler(c *fiber.Ctx) error {
	n := c.Params("name")
	slog.Info("request recieved", "name", n)
	return nil
}

func main() {
	appConfig := fiber.Config{
		AppName:           "Some interesting API",
		EnablePrintRoutes: true,
		ServerHeader:      "Test API 1",
		Prefork:           true,
	}
	app := fiber.New(appConfig)
	app.Get("/", getHandler).Name("Get default")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
