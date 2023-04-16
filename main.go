package main

import (
	"log"

	frontend "github.com/alejandro31118/sveltekit-embeded-golang/sveltekit-front"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	app := fiber.New()

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello from GO server"})
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         frontend.BuildHTTPFS(),
		NotFoundFile: "index.html",
	}))

	log.Fatal(app.Listen(":3000"))
}
