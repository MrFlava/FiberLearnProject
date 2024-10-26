package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/constraints/FiberLearnProject/Lesson10/handlers"
	"golang.org/x/exp/constraints/FiberLearnProject/Lesson10/services"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)

	itemsRepo := services.NewItemRepository()
	itemApi := handlers.NewItemApi(itemsRepo)

	api := app.Group("/api")
	api.Get("/items", itemApi.GetItems)
	api.Post("/items", itemApi.CreateItems)
	api.Get("/items/:id<int>", itemApi.GetItemById)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
