package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Desc  string `json:"desc"`
}

var db = make(map[string]Item)

func SearchItems(c *fiber.Ctx) error {
	slog.Info("search items request recieved", "path", c.Path(), "method", c.Method())

	q := c.Query("q")

	if q == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	var items = []Item{}
	for _, item := range db {
		if strings.Contains(item.Name, q) || strings.Contains(item.Desc, q) {
			items = append(items, item)
		}
	}

	return c.JSON(items)
}

func GetItems(c *fiber.Ctx) error {
	slog.Info("get items request recieved", "path", c.Path(), "method", c.Method())

	var items = []Item{}
	for _, item := range db {
		items = append(items, item)
	}

	return c.JSON(items)
}

func GetItemById(c *fiber.Ctx) error {
	slog.Info("get item request recieved", "path", c.Path(), "method", c.Method())
	id := c.Params("Id")
	item := db[id]

	if item.ID == "" {
		return c.SendStatus(http.StatusNotFound)
	}
	return c.JSON(item)
}

func DeleteItemById(c *fiber.Ctx) error {
	slog.Info("delete item request recieved", "path", c.Path(), "method", c.Method())
	id := c.Params("Id")
	item := db[id]

	if item.ID == "" {
		return c.SendStatus(http.StatusNotFound)
	}

	delete(db, id)

	return c.SendStatus(http.StatusOK)
}

func UpdateItemById(c *fiber.Ctx) error {
	slog.Info("update item request recieved", "path", c.Path(), "method", c.Method())
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	type ItemCreateRequest struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
		Desc  string `json:"desc"`
	}
	var reqItem ItemCreateRequest
	if err := c.BodyParser(&reqItem); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	item := db[id]
	if item.ID == "" {
		return c.SendStatus(http.StatusNotFound)
	}

	item.Name = reqItem.Name
	item.Price = reqItem.Price
	item.Desc = reqItem.Desc
	db[id] = item // update db

	return c.JSON(item)
}

func CreateItems(c *fiber.Ctx) error {
	slog.Info("create item request recieved", "path", c.Path(), "method", c.Method())
	type ItemCreateReqest struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
		Desc  string `json:"desc"`
	}

	var reqItem ItemCreateReqest
	if err := c.BodyParser(&reqItem); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	item := Item{
		Name:  reqItem.Name,
		Price: reqItem.Price,
		Desc:  reqItem.Desc,
	}

	item.ID = fmt.Sprintf("%v", len(db)+1)
	db[item.ID] = item

	slog.Info("item created", "id", item.ID)
	return c.JSON(item)
}
