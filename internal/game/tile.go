package game

// Tile interface defines capabilities of a particular cell on the board.
type Tile interface {
	// Blackhole marks current tile as blackhole or unmarks it if value is 'false'.
	Blackhole(bool)

	// IsBlackhole indicates whether current tile is blackhole.
	IsBlackhole() bool

	// IsOpened indicates whether current tile opened (revealed) or not.
	IsOpened() bool

	// Open marks current tile as opened.
	Open()

	// SurroundingBlackholesCount reports number of blackholes around
	SurroundingBlackholesCount() byte

	// SetSurroundingBlackholesCount updates number of adjacent blackholes.
	SetSurroundingBlackholesCount(count byte)
}

type tile struct {
	isOpened                   bool
	isBlackhole                bool
	surroundingBlackholesCount byte
}

// compile-time check for interface
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

func (c *tile) SetSurroundingBlackholesCount(count byte) {
	c.surroundingBlackholesCount = count
}
