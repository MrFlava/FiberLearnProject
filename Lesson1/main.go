package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

var items []Item = []Item{
	{1, "Item 1", 10.99},
	{2, "Item 2", 5.99},
	{3, "Item 3", 25.99},
}

func getAllItemsHandler(c *fiber.Ctx) error {
	log.Infof("Recieved a request %s for items", c.Method())
	return c.JSON(items)
}

func getItemHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	for _, item := range items {
		if item.ID == id {
			return c.JSON(item)
		}
	}

	return c.Status(http.StatusNotFound).SendString("Item not found")
}

func postItemsHandler(c *fiber.Ctx) error {
	var item Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
	}

	item.ID = len(items) + 1
	items = append(items, item)
	return c.JSON(item)
}

func main() {
	app := fiber.New()
	app.Get("/items", getAllItemsHandler)
	app.Get("/items/:id", getItemHandler)
	app.Post("/items", postItemsHandler)

	log.Info("Starting server on port 3000")
	err := app.Listen(":3000")

	if err != nil {
		log.Fatal(err)
	}
}
