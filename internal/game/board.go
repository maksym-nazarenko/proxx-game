package game

import "fmt"

type BoardArea = uint16

type (
	// BlackholesPlacerFunc defines a functions that is capable of filling the board with blackholes
	BlackholesPlacerFunc func(b Board, count BoardArea)

	// Board defines capabilities of the game `Board`.
	Board interface {
		// OpenTile opens tile by its coordinates.
		// The first return value indicates whether tile has a bloackhole behind it.
		// The second return value inidcates error, if any.
		OpenTile(Point) (bool, error)

		// Size returns a length of the board side.
		Size() Coordinate

		// TileAt returns Tile object at specific position.
		// This method returns error in case of invalid coordinates, e.g. 'out of bounds'.
		TileAt(Point) (Tile, error)
	}
)

type board struct {
	tiles []Tile
	size  Coordinate
}

var _ Board = (*board)(nil)

func NewBoard(size Coordinate, blackholesCount BoardArea, blackholesPlacer BlackholesPlacerFunc) *board {
	b := &board{
		size: size,
	}
	b.tiles = make([]Tile, 0, size*size)
	for i := 0; i < int(size*size); i++ {
		b.tiles = append(b.tiles, NewTile())
	}
	blackholesPlacer(b, blackholesCount)

	return b
}

func (b *board) OpenTile(p Point) (bool, error) {
	tile, err := b.TileAt(p)
	if err != nil {
		return false, err
	}
	tile.Open()

	return tile.IsBlackhole(), nil
}

func (b *board) Size() Coordinate {
	return b.size
}

func (b *board) TileAt(p Point) (Tile, error) {
	if p.X() > b.size || p.Y() > b.size {
		return nil, NewOutOfBoundsError(fmt.Sprintf("point (%d, %d) points to tile outside the board", p.X(), p.Y()))
	}

	return b.tiles[(p.Y()-1)*b.size+(p.X()-1)], nil
}
