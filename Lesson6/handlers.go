package main

import (
	"log/slog"

	"github.com/beevik/guid"
	"github.com/gofiber/fiber/v2"
)

func requestLogger(c *fiber.Ctx) error {
	reqId := c.Request().Header.Peek("REQUEST-ID")
	slog.Info("Got request", "method", c.Method(), "path", c.Path(), "id", c.Params("id"), "requestId", reqId)
	return c.Next()
}

func addRequestId(c *fiber.Ctx) error {
	uid := guid.New()
	c.Request().Header.Add("REQUEST-ID", uid.String())
	return c.Next()
}

func getItemById(c *fiber.Ctx) error {
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
