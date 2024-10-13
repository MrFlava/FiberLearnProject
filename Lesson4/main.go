package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
		CaseSensitive:     true,
	}

	app := fiber.New(appConfig)
	app.Get("/books", getAllBooksLCHnalder).Name("Get all boks lowercase")
	app.Get("/Books", getAllBooksUCHnalder).Name("Get all boks uppercase")
	app.Get("/books/:id", getBookByIdHandler).Name("Get book by ID")
	app.Get("/author/:id?", getAuthorsOrAuthorByIdHandler).Name("Get all authors or a specific author")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
