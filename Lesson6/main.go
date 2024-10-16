package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	app.Get("/items/:id", addRequestId, requestLogger, getItemById).Name("Get single item")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
