package main

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func getHandler(c *fiber.Ctx) error {
	log.Infof("Recieved a request %s for items", c.Method())
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

	log.Info("Starting server on port 3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
