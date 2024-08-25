package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/grayjunzi/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()

	v1 := app.Group("/api/v1")

	v1.Get("/users", api.HandleGetUsers)
	v1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
