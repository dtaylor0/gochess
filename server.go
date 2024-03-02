package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type BoardRows [][]map[string]string

func main() {

	// move := "e4"

	logger := log.New(os.Stdout, "", 0)
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("game", fiber.Map{})
	})

	app.Get("/board", func(c *fiber.Ctx) error {
		perspective := rand.Intn(2)
		fen_start := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
        board_input := createBoard(fen_start, perspective)
		return c.Render("board", fiber.Map{"Rows": board_input})
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

func applyMove(move string, fen string, perspective int) (string, error) {
	return "", nil
}

func reverseString(s string) string {
	res := ""
	for _, r := range s {
		res = strings.Join([]string{string(r), res}, "")
	}
	return res
}

func createBoard(fen string, perspective int) BoardRows {
	var rows BoardRows
	var currRow []map[string]string
	if perspective == 1 {
		fen = reverseString(fen)
	}
	for i, r := range fen {
		if unicode.IsDigit(r) {
			repeat := int(r - '0')
			for repeat > 0 {
				currRow = append(currRow, map[string]string{"Piece": "", "Id": ""})
				repeat--
			}
		} else if r == '/' {
			rows = append(rows, currRow)
			currRow = nil
		} else {
			currRow = append(
				currRow,
				map[string]string{
					"Piece": string([]rune{r}),
					"Id":    strings.Join([]string{string(r), fmt.Sprint(i)}, ""),
				},
			)
		}
	}
    rows = append(rows, currRow)
	return rows
}
