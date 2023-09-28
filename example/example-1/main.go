package main

import (
	"github.com/Taiwrash/selfserveruploader"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Handle file uploads at the "/upload" endpoint
	app.Post("/upload", selfserveruploader.Upload)

	app.Listen(":4000")
}
