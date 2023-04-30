package game

import "fmt"

// BoardArea is an alias type for the max count of tiles on the board.
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

		// OpenedTilesCount reports total number of tiles opened on the board, excluding balckholes
		OpenedTilesCount() BoardArea

		// Size returns a length of the board side.
		Size() Coordinate

		// TileAt returns Tile object at specific position.
		// This method returns error in case of invalid coordinates, e.g. 'out of bounds'.
		TileAt(Point) (Tile, error)
	}
)

type board struct {
	tiles            []Tile
	size             Coordinate
	openedTilesCount BoardArea
}

// compile-time interface checking for Board implementation
var _ Board = (*board)(nil)

// NewBoard creates new instance of a board.
//
// Arguments:
//
//	size - a size of the board side.
//	blackholesCount - number of blackholes to place on the board.
//	blackholesPlacer - function-strategy to place blackholes on the board.
func NewBoard(size Coordinate, blackholesCount BoardArea, blackholesPlacer BlackholesPlacerFunc) *board {
	b := &board{
		size: size,
	}
	b.tiles = make([]Tile, 0, size*size)
	for i := 0; i < int(size*size); i++ {
		b.tiles = append(b.tiles, NewTile())
	}
	blackholesPlacer(b, blackholesCount)

	b.calculateAdjacentBlackholes()

	return b
}

func (b *board) OpenTile(p Point) (bool, error) {
	tile, err := b.TileAt(p)
	if err != nil {
		return false, err
	}
	if tile.IsOpened() {
		return false, nil
	}

	tile.Open()
	if tile.IsBlackhole() {
		return true, nil
	}
	b.openedTilesCount++

	if tile.SurroundingBlackholesCount() == 0 {
		for _, neighborPoint := range b.neighborPoints(p) {
			// ignore blackhole flag and error, as the current tile has no surrounding blackholes
			// and all tiles must be valid, since they are collected from the set inside board
			_, _ = b.OpenTile(neighborPoint)
		}
	}

	return false, nil
}

func (b board) OpenedTilesCount() BoardArea {
	return b.openedTilesCount
}

func (b *board) Size() Coordinate {
	return b.size
}

// TileAt implements finding a tile by its coordinates on the board.
//
// Current implementation may return an error if coordinates are out of bounds.
func (b *board) TileAt(p Point) (Tile, error) {
	// todo(maksym): consider different strategy for error:
	// - panic
	// - lose the game
	if p.X() > b.size || p.Y() > b.size {
		return nil, NewOutOfBoundsError(fmt.Sprintf("point (%d, %d) points to tile outside the board", p.X(), p.Y()))
	}

	return b.tiles[(p.Y()-1)*b.size+(p.X()-1)], nil
}

// neighborPoints collects all neighbor tiles coordinates, including diagonals.
func (b board) neighborPoints(p Point) []Point {
	// current implementation is sub-optimal with a bunch of "if's" to check if we have a neighbor tile at that point.
	// without necessity or technical requirement, it's fine for now.
	neighbors := []Point{}

	// upper row
	if p.X() > 1 && p.Y() > 1 {
		neighbors = append(neighbors, NewPoint(p.X()-1, p.Y()-1))
	}
	if p.Y() > 1 {
		neighbors = append(neighbors, NewPoint(p.X(), p.Y()-1))
	}
	if p.X() < b.size && p.Y() > 1 {
		neighbors = append(neighbors, NewPoint(p.X()+1, p.Y()-1))
	}

	// left/right tiles
	if p.X() > 1 {
		neighbors = append(neighbors, NewPoint(p.X()-1, p.Y()))
	}
	if p.X() < b.size {
		neighbors = append(neighbors, NewPoint(p.X()+1, p.Y()))
	}

	// lower row
	if p.X() > 1 && p.Y() < b.size {
		neighbors = append(neighbors, NewPoint(p.X()-1, p.Y()+1))
	}
	if p.Y() < b.size {
		neighbors = append(neighbors, NewPoint(p.X(), p.Y()+1))
	}
	if p.X() < b.size && p.Y() < b.size {
		neighbors = append(neighbors, NewPoint(p.X()+1, p.Y()+1))
	}

	return neighbors
}

// calculateAdjacentBlackholes calculates the number of adjucent blackholes and sets the value for each tile
func (b board) calculateAdjacentBlackholes() {
	// brute-force implementation, can be improved with proper algorithm.
	// Despite it looks a bit ugly, it is an initialisation phase, so it's okay.
	// It must be revised if board has really huge size of this action is done at runtime (for example, blackholes a dynamically added)
	for row := 1; row <= int(b.size); row++ {
		for col := 1; col <= int(b.size); col++ {
			var blackholes byte
			currentPoint := NewPoint(Coordinate(col), Coordinate(row))
			currentTile, _ := b.TileAt(currentPoint)
			if currentTile.IsBlackhole() {
				continue
			}
			for _, p := range b.neighborPoints(currentPoint) {
				tile, _ := b.TileAt(p)
				if tile.IsBlackhole() {
					blackholes++
				}
			}
			if blackholes > 0 {
				currentTile.SetSurroundingBlackholesCount(blackholes)
			}
		}
	}
}
