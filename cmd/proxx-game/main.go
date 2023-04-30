package main

import (
	"context"
	"flag"
	"log"

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
	if err := app.ValidateConfig(config); err != nil {
		return err
	}

	// channel with the game result:
	// true - win
	// false - lose
	resultChan := make(chan bool, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	board := game.NewBoard(config.BoardSize, config.BlackholesCount, game.DefaultBlackholesPlaceStrategy)

	totalTilesCount := game.BoardArea(config.BoardSize * config.BoardSize)

	app := tview.NewApplication()

	table := createBoardUI(config.BoardSize)
	helpText := tview.NewTextView().SetText(`Press 'ESC' at any time to exit the game.`)

	stats := tview.NewTextView().SetText("Press 'Enter' to start the game.")
	rules := tview.NewTextView().
		SetText(`Rules

Open tiles one-by-one by clicking enter.
If you open a tile with blackhole behind it, you lose.
When all non-blackholes tiles are opened - you win.
If you avoid blackhole, number on the tile tells you how many blackholes surround this tile (number in range 0..8)
`)

	go func() {
		gameResult := <-resultChan
		defer cancel()
		defer app.Draw()
		defer refreshBoardUI(board, table)
		if gameResult {
			helpText.
				SetText("You won!!!").
				SetTextColor(tcell.ColorGreen).
				SetBorderAttributes(tcell.AttrBold)
			return
		}
		helpText.
			SetText("Game over. You were almost there...").
			SetTextColor(tcell.ColorDarkRed).
			SetBorderAttributes(tcell.AttrBold)
	}()
	table.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyEnter:
			table.SetSelectable(true, true)
			refreshBoardUI(board, table)
			go updateStats(ctx, stats, app)
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
