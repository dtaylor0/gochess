package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("board", fiber.Map{})
	})

	app.Get("/game", func(c *fiber.Ctx) error {
		return c.Render("game", fiber.Map{})
	})
	logger.Println("Hello from logger")

    app.Listen(":3000")

}
