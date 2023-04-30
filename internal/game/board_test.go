package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoardOpenTile(t *testing.T) {
	testCases := []struct {
		name            string
		board           Board
		tileCoordinates Point
		expected        bool
	}{
		{
			name: "test OpenTile without blackhole",
			board: func() Board {
				b := NewBoard(2, 0, func(b Board, count BoardArea) { markAllWithBlackholeValue(b, false, 0) })
				tile, _ := b.TileAt(NewPoint(2, 1))
				tile.Blackhole(false)

				return b
			}(),
			tileCoordinates: NewPoint(2, 1),
		},
		{
			name: "test OpenTile on blackhole",
			board: func() Board {
				b := NewBoard(2, 0, func(b Board, count BoardArea) { markAllWithBlackholeValue(b, true, 0) })
				return b
			}(),
			tileCoordinates: NewPoint(2, 1),
			expected:        true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.board.OpenTile(tc.tileCoordinates)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestBoardOpenTileAutoOpen(t *testing.T) {
	testCases := []struct {
		name            string
		board           Board
		tileCoordinates Point
		expectOpened    []Point
	}{
		{
			// H - blackhole, 0-8 - regular tile, X - tile to open
			//
			// 1 1 1 H
			// H 2 X 2
			// 1 2 H 1
			// 0 1 1 1
			name: "test OpenTile with blackholes around",
			board: func() Board {
				b := NewBoard(4, 0, func(b Board, count BoardArea) {
					for _, p := range []Point{NewPoint(4, 1), NewPoint(1, 2), NewPoint(3, 3)} {
						tile, _ := b.TileAt(p)
						tile.Blackhole(true)
					}
				})
				return b
			}(),
			tileCoordinates: NewPoint(3, 2),
			expectOpened: []Point{
				NewPoint(3, 2),
			},
		},
		{
			// H - blackhole, 0-8 - regular tile, X - tile to open, + - will be auto-opened
			//
			// H + X +
			// 1 + + +
			// 0 + + +
			// 0 1 H 1
			name: "test OpenTile with further auto-open",
			board: func() Board {
				b := NewBoard(4, 0, func(b Board, count BoardArea) {
					for _, p := range []Point{NewPoint(1, 1), NewPoint(3, 4)} {
						tile, _ := b.TileAt(p)
						tile.Blackhole(true)
					}
				})

				return b
			}(),
			tileCoordinates: NewPoint(3, 1),
			expectOpened: []Point{
				NewPoint(2, 1),
				NewPoint(3, 1),
				NewPoint(4, 1),
				NewPoint(2, 2),
				NewPoint(3, 2),
				NewPoint(4, 2),
				NewPoint(2, 3),
				NewPoint(3, 3),
				NewPoint(4, 3),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.board.OpenTile(tc.tileCoordinates)
			require.NoError(t, err)
			assert.False(t, actual)
			assert.ElementsMatch(t, tc.expectOpened, getAllOpenTileCoordinates(tc.board))
		})
	}
}

func TestNeighbor(t *testing.T) {
	testCases := []struct {
		name            string
		board           *board
		tileCoordinates Point
		expectedPoints  []Point
	}{
		{
			name: "top left corner tile",
			board: func() *board {
				return NewBoard(3, 0, func(b Board, count BoardArea) {})
			}(),
			tileCoordinates: NewPoint(1, 1),
			expectedPoints: []Point{
				NewPoint(2, 1),
				NewPoint(1, 2),
				NewPoint(2, 2),
			},
		},
		{
			name: "top right corner tile",
			board: func() *board {
				return NewBoard(3, 0, func(b Board, count BoardArea) {})
			}(),
			tileCoordinates: NewPoint(3, 1),
			expectedPoints: []Point{
				NewPoint(2, 1),
				NewPoint(2, 2),
				NewPoint(3, 2),
			},
		},
		{
			name: "bottom left tile",
			board: func() *board {
				return NewBoard(3, 0, func(b Board, count BoardArea) {})
			}(),
			tileCoordinates: NewPoint(1, 3),
			expectedPoints: []Point{
				NewPoint(1, 2),
				NewPoint(2, 2),
				NewPoint(2, 3),
			},
		},
		{
			name: "bottom right tile",
			board: func() *board {
				return NewBoard(3, 0, func(b Board, count BoardArea) {})
			}(),
			tileCoordinates: NewPoint(3, 3),
			expectedPoints: []Point{
				NewPoint(2, 2),
				NewPoint(3, 2),
				NewPoint(2, 3),
			},
		},
		{
			name: "surrounded",
			board: func() *board {
				return NewBoard(5, 0, func(b Board, count BoardArea) {})
			}(),
			tileCoordinates: NewPoint(2, 3),
			expectedPoints: []Point{
				NewPoint(1, 2),
				NewPoint(2, 2),
				NewPoint(3, 2),
				NewPoint(1, 3),
				NewPoint(3, 3),
				NewPoint(1, 4),
				NewPoint(2, 4),
				NewPoint(3, 4),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.board.neighborPoints(tc.tileCoordinates)
			assert.ElementsMatch(t, tc.expectedPoints, actual)
		})
	}
}

// getAllOpenTileCoordinates is a helper function to collect coordinates of all open tiles on the board.
func getAllOpenTileCoordinates(b Board) []Point {
	points := []Point{}
	for row := 0; row < int(b.Size()); row++ {
		for col := 0; col < int(b.Size()); col++ {
			p := NewPoint(Coordinate(col+1), Coordinate(row+1))
			tile, _ := b.TileAt(p)
			if tile.IsOpened() {
				points = append(points, p)
			}
		}
	}

	return points
}

// markAllWithBlackholeValue is a helper function to mark all tiles on the board with pre-defined Blackhole indicator.
// Effectively, create board with all tiles marked as regular or blackholes.
func markAllWithBlackholeValue(b Board, v bool, _ BoardArea) {
	for i := 1; i <= int(b.Size()); i++ {
		for j := 1; j <= int(b.Size()); j++ {
			tile, _ := b.TileAt(NewPoint(Coordinate(j), Coordinate(i)))
			tile.Blackhole(v)
		}
	}
}
