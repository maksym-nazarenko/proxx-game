package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTile(t *testing.T) {
	testCases := []struct {
		name     string
		new      func() Tile
		testFunc func(*testing.T, Tile)
	}{
		{
			name: "regular tile",
			new:  func() Tile { return NewTile() },
			testFunc: func(t *testing.T, c Tile) {
				assert.False(t, c.IsBlackhole())
			},
		},
		{
			name: "blackhole tile",
			new:  func() Tile { return NewBlackholeTile() },
			testFunc: func(t *testing.T, c Tile) {
				assert.True(t, c.IsBlackhole())
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testFunc(t, tc.new())
		})
	}
}

func TestOpenTile(t *testing.T) {
	c := NewTile()
	assert.False(t, c.IsOpened())
	c.Open()
	assert.True(t, c.IsOpened())
	c.Open()
	assert.True(t, c.IsOpened(), "subsequent call to Open() should not affect the state")
}

func TestBlackholeTile(t *testing.T) {
	c := NewTile()
	assert.False(t, c.IsBlackhole())
	c.Blackhole(false)
	assert.False(t, c.IsBlackhole(), "unsetting blackhole flag of non-blackhole tile should have no effect")

	c.Blackhole(true)
	assert.True(t, c.IsBlackhole(), "after setting blackhole flag, IsBlackhole() should report the changed state")
	c.Blackhole(true)
	assert.True(t, c.IsBlackhole(), "subsequent call to Blackhole() with the same value should have no effect")
}
