package game

// Tile interface defines capabilities of a particular cell on the board.
type Tile interface {
	Blackhole(bool)
	IsBlackhole() bool
	IsOpened() bool
	Open()
	SurroundingBlackholesCount() byte
}

type tile struct {
	isOpened                   bool
	isBlackhole                bool
	surroundingBlackholesCount byte
}

var _ Tile = (*tile)(nil)

// NewTile creates new instance of `tile` without modifications - `zero` value.
func NewTile() *tile {
	return &tile{}
}

// NewBlackholeTile creates new tile with blackhole behind it.
func NewBlackholeTile() *tile {
	return &tile{
		isBlackhole: true,
	}
}

func (c *tile) Blackhole(v bool) {
	c.isBlackhole = v
}

func (c *tile) IsBlackhole() bool {
	return c.isBlackhole
}

func (c *tile) IsOpened() bool {
	return c.isOpened
}

func (c *tile) Open() {
	c.isOpened = true
}

func (c *tile) SurroundingBlackholesCount() byte {
	return c.surroundingBlackholesCount
}
