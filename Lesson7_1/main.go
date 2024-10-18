package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"golang.org/x/exp/constraints/FiberLearnProject/Lesson7/handlers"
	"golang.org/x/exp/constraints/FiberLearnProject/Lesson7/middlewares"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	app.Use(requestid.New(), logger.New())
	app.Get("/items", middlewares.VerifyLogin, handlers.GetItems)
	app.Post("/login", handlers.DoLogin)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
