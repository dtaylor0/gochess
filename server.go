package main

import (
	// "math/rand"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// perspective := rand.Intn(1)
	// fen_start := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	// move := "e4"

	logger := log.New(os.Stdout, "", 0)
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("game", fiber.Map{})
	})

	app.Get("/board", func(c *fiber.Ctx) error {
		return c.Render("board", fiber.Map{})
	})

	app.Get("/board/update", func(c *fiber.Ctx) error {
		payload := struct {
			FEN         string `json:"fen"`
			Perspective string `json:"perspective"`
			Move        string `json:"move"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		// updatedBoard, err := updateBoard(
		// 	payload.Move,
		// 	payload.FEN,
		// 	payload.Perspective,
		// )
		// if err != nil {
		//     return err
		// }

		return c.Render("board", fiber.Map{})
	})
	logger.Println("Hello from logger")

	app.Listen(":3000")

}

func createBoard(fen string, perspective int) {
}

func applyMove(move string, fen string, perspective int) (string, error) {
	return "", nil
}
