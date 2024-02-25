package main

import (
	"html/template"
	"log"
	"os"
	"unicode"

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
		board2_input := [8][8]string{
			{
				"r", "n", "b", "q", "k", "b", "n", "r",
			}, {
				"p", "p", "p", "p", "p", "p", "p", "p",
			}, {
				"", "", "", "", "", "", "", "",
			}, {
				"", "", "", "", "", "", "", "",
			}, {
				"", "", "", "", "", "", "", "",
			}, {
				"", "", "", "", "", "", "", "",
			}, {
				"P", "P", "P", "P", "P", "P", "P", "P",
            }, {
				"R", "N", "B", "Q", "K", "B", "N", "R",
            },
		}
        return c.Render("board2", fiber.Map{"Rows": board2_input})
		return c.Render("pieces/K", fiber.Map{}, "square", "board")
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

func createBoard(fen string, perspective int) *template.Template {
	tpl := template.New("board")
	skip := 0
	if perspective == 0 {
		for _, c := range fen {
			if skip > 0 {
				skip--
				// TODO: add empty square
			}
			if unicode.IsDigit(c) {
				skip = int(c-'0') - 1
				// TODO: add empty square
			} else {
				// TODO: add square with piece
			}
		}
		return tpl
	}
	for i := len(fen); i >= 0; i-- {
	}
	return tpl
}

func applyMove(move string, fen string, perspective int) (string, error) {
	return "", nil
}
