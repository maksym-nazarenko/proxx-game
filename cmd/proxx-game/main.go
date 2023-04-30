package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/maksym-nazarenko/proxx-game/internal/app"
	"github.com/maksym-nazarenko/proxx-game/internal/game"
	"github.com/rivo/tview"
)

func main() {
	size := flag.Int("size", 5, "Board size")
	blackholes := flag.Int("blackholes", 10, "Number of blackholes to place on the board.")
	flag.Parse()

	config := app.Configuration{
		BoardSize:       game.Coordinate(*size),
		BlackholesCount: game.BoardArea(*blackholes),
	}
	if err := realMain(config); err != nil {
		log.Fatal(err)
	}
}

func realMain(config app.Configuration) error {
	// channel with the game result:
	// true - win
	// false - lose
	resultChan := make(chan bool, 1)

	if err := app.ValidateConfig(config); err != nil {
		return err
	}
	board := game.NewBoard(config.BoardSize, config.BlackholesCount, game.DefaultBlackholesPlaceStrategy)
	totalTilesCount := game.BoardArea(config.BoardSize * config.BoardSize)

	app := tview.NewApplication()

	table := createBoardUI(config.BoardSize)
	helpText := tview.NewTextView().SetText(`Press 'ESC' at any time to exit the game. Press 'Enter' to start the game.`)

	stats := tview.NewTextView().SetText("Stats:")
	rules := tview.NewTextView().SetText("Rules")

	go func() {
		gameResult := <-resultChan
		defer refreshBoardUI(board, table)
		if gameResult {
			helpText.SetText("You won!!!")
			return
		}
		helpText.SetText("Aah, you lose :(")
	}()
	table.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyEnter:
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row, column int) {
		blackhole, err := board.OpenTile(game.NewPoint(game.Coordinate(column+1), game.Coordinate(row+1)))
		if err != nil {
			panic("wtf")
		}
		if blackhole {
			resultChan <- false
			return
		}
		if (game.BoardArea(totalTilesCount) - board.OpenedTilesCount()) == config.BlackholesCount {
			resultChan <- true
			return
		}

		refreshBoardUI(board, table)
	})

	grid := createGridUI(table, helpText, stats, rules)
	app.SetRoot(grid, true)

	return app.
		EnableMouse(true).
		Run()
}

func createBoardUI(size game.Coordinate) *tview.Table {
	table := tview.NewTable().SetBorders(true)
	var rows, cols int = int(size), int(size)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			table.SetCell(
				row, col,
				tview.NewTableCell("?").
					SetAlign(tview.AlignCenter),
			)
		}
	}
	return table
}

func createGridUI(table, helpText, stats, rules tview.Primitive) *tview.Grid {
	return tview.NewGrid().
		SetRows(2, 4, 0).
		SetColumns(-2, -5).
		SetBorders(true).
		AddItem(helpText, 0, 0, 1, 2, 0, 0, false).
		AddItem(stats, 1, 0, 1, 1, 0, 0, false).
		AddItem(rules, 2, 0, 1, 1, 0, 0, false).
		AddItem(table, 1, 1, 2, 1, 0, 0, true)
}

func refreshBoardUI(board game.Board, table *tview.Table) {
	for row := 1; row <= int(board.Size()); row++ {
		for col := 1; col <= int(board.Size()); col++ {
			tile, _ := board.TileAt(game.NewPoint(game.Coordinate(col), game.Coordinate(row)))
			if !tile.IsOpened() {
				continue
			}
			tableCellUI := table.GetCell(row-1, col-1)
			if tile.IsBlackhole() {
				tableCellUI.SetText("H")
				continue
			}
			if tile.SurroundingBlackholesCount() > 0 {
				tableCellUI.Text = strconv.Itoa(int(tile.SurroundingBlackholesCount()))
				continue
			}

			tableCellUI.Text = " "
		}
	}
}
