package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

const grId = "ABCDEFGH"

var idx int = 0

func getHandlerId() string {
	c := grId[idx%9]
	idx++
	return fmt.Sprintf("grID-%v-%c", idx, c)
}

func getHandler(c *fiber.Ctx) error {
	ccId := getHandlerId()
	n := c.Params("name")

	go func() {
		slog.Info("starting handler", "id", ccId, "name", n)
		t := time.After(10 * time.Second)

		for {
			select {
			case <-t:
				slog.Info("handler done", "ccId", ccId, "name", n)
				return

			default:
				slog.Info("handler running", "ccId", ccId, "name", n)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	slog.Info("request revieved", "name", n)
	return nil
}

func main() {
	appConfig := fiber.Config{
		AppName:           "Some interesting API",
		EnablePrintRoutes: true,
		ServerHeader:      "Test API 1",
		Prefork:           true,
	}
	app := fiber.New(appConfig)
	app.Get("/:name", getHandler).Name("Get default")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
