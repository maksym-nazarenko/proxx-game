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

func markAllWithBlackholeValue(b Board, v bool, _ BoardArea) {
	for i := 1; i <= int(b.Size()); i++ {
		for j := 1; j <= int(b.Size()); j++ {
			tile, _ := b.TileAt(NewPoint(Coordinate(j), Coordinate(i)))
			tile.Blackhole(v)
		}
	}
}
