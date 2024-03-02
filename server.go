package main

import (
	"log"
	"math/rand"
	"os"
	"slices"
	"strings"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type BoardRows [][]map[string]string

func main() {

	fen_start := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	// move := "e4"

	logger := log.New(os.Stdout, "", 0)
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("game", fiber.Map{})
	})

	app.Get("/board", func(c *fiber.Ctx) error {
		perspective := rand.Intn(1)
		board_input := createBoard(fen_start, perspective)
		board_input = BoardRows{
			{
				{"Piece": "r", "Id": "ra8"},
				{"Piece": "n", "Id": "nb8"},
				{"Piece": "b", "Id": "bc8"},
				{"Piece": "q", "Id": "qd8"},
				{"Piece": "k", "Id": "ke8"},
				{"Piece": "b", "Id": "bf8"},
				{"Piece": "n", "Id": "ng8"},
				{"Piece": "r", "Id": "rh8"},
			}, {
				{"Piece": "p", "Id": "pa7"},
				{"Piece": "p", "Id": "pb7"},
				{"Piece": "p", "Id": "pc7"},
				{"Piece": "p", "Id": "pd7"},
				{"Piece": "p", "Id": "pe7"},
				{"Piece": "p", "Id": "pf7"},
				{"Piece": "p", "Id": "pg7"},
				{"Piece": "p", "Id": "ph7"},
			}, {
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
			}, {
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
			}, {
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
			}, {
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
				{"Piece": "", "Id": ""},
			}, {
				{"Piece": "P", "Id": "pa2"},
				{"Piece": "P", "Id": "pb2"},
				{"Piece": "P", "Id": "pc2"},
				{"Piece": "P", "Id": "pd2"},
				{"Piece": "P", "Id": "pe2"},
				{"Piece": "P", "Id": "pf2"},
				{"Piece": "P", "Id": "pg2"},
				{"Piece": "P", "Id": "ph2"},
			}, {
				{"Piece": "R", "Id": "Ra1"},
				{"Piece": "N", "Id": "Nb1"},
				{"Piece": "B", "Id": "Bc1"},
				{"Piece": "Q", "Id": "Qd1"},
				{"Piece": "K", "Id": "Ke1"},
				{"Piece": "B", "Id": "Bf1"},
				{"Piece": "N", "Id": "Ng1"},
				{"Piece": "R", "Id": "Rh1"},
			},
		}
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

func createBoard(fen string, perspective int) BoardRows {
	var rows BoardRows
	var currRow []map[string]string
	if perspective == 0 {
		for i, r := range fen {
			if unicode.IsDigit(r) {
				repeat := int(r - '0')
				for repeat > 0 {
					currRow = append(currRow, map[string]string{"Piece": "", "Id": ""})
				}
			} else if r == '/' {
				rows = append(rows, currRow)
				clear(currRow)
			} else {
				currRow = append(
					currRow,
					map[string]string{
						"Piece": string([]rune{r}),
						"Id":    strings.Join(string([]rune{r}), string(i)),
					},
				)
			}
		}
	}
}
