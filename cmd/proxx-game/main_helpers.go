package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/maksym-nazarenko/proxx-game/internal/game"
	"github.com/rivo/tview"
)

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

func updateStats(ctx context.Context, statsBlock *tview.TextView, app *tview.Application) {
	startTime := time.Now()
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Second):
			statsBlock.SetText(fmt.Sprintf(`Game time: %v`, time.Since(startTime).Truncate(time.Second)))
			app.Draw()
		}
	}
}
