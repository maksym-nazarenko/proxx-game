package game

type (
	Coordinate = byte

	// Point represents coordinates on a two-dimension board.
	Point struct {
		x Coordinate
		y Coordinate
	}
)

func NewPoint(x, y Coordinate) Point {
	return Point{x, y}
}

func (p Point) X() Coordinate {
	return p.x
}

func (p Point) Y() Coordinate {
	return p.y
}
