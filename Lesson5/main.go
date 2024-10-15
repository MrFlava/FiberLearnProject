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
	app.Get("/books", getAllBooksHnalder).Name("Get all boks")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
