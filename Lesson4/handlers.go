package main

import (
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func getAllBooksLCHnalder(c *fiber.Ctx) error {
	slog.Info("request to get all books - lovwercase")
	return nil
}

func getAllBooksUCHnalder(c *fiber.Ctx) error {
	slog.Info("request to get all books - uppercase")
	return nil
}

func getBookByIdHandler(c *fiber.Ctx) error {
	bookId := c.Params("id")
	slog.Info("request to get book by ID", "bookId", bookId)
	return nil
}

func getAuthorsOrAuthorByIdHandler(c *fiber.Ctx) error {
	authorId := c.Params("id")

	if authorId == "" {
		slog.Info("request to get all authors")
	}

	slog.Info("request to get author by ID", "authorId", authorId)
	return nil
}
func getItemsHandler(c *fiber.Ctx) error {
	itemPath := c.Params("*")

	if itemPath == "" {
		slog.Info("request all items")
		return nil
	}

	slog.Info("request to get an item with sub-paths:", "itemPath", itemPath)
	subPaths := strings.Split(itemPath, "/")

	for _, subPath := range subPaths {
		slog.Info("\titem subPath", "subPath", subPath)
	}

	return nil
}
