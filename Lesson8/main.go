package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/constraints/FiberLearnProject/Lesson8/handlers"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	app.Get("/items", handlers.GetItems)
	app.Post("/items", handlers.CreateItems)
	app.Get("/items/:id<int>", handlers.GetItemById)
	app.Delete("/items/:id<int>", handlers.DeleteItemById)
	app.Put("/items/:id<int>", handlers.UpdateItemById)
	app.Get("/items/search", handlers.SearchItems)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
