package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/grayjunzi/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		Id:   "1",
		Name: "James",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("")
}
