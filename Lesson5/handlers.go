package main

import (
	"github.com/gofiber/fiber/v2"
)

func getAllBooksHnalder(c *fiber.Ctx) error {
	// err := fiber.NewError(http.StatusInternalServerError, "this is some kind of a error")
	person := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "John",
	}
	return c.JSON(person)
}
