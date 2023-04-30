package game

import "math/rand"

var DefaultBlackholesPlaceStrategy BlackholesPlacerFunc = randomBlackholesPlaceStrategy

func randomBlackholesPlaceStrategy(b Board, count BoardArea) {
	totalTiles := int(b.Size() * b.Size())
	activeIndexes := make([]Point, totalTiles)

	for i := 0; i < totalTiles; i++ {
		activeIndexes[i] = NewPoint(Coordinate(i%int(b.Size())+1), Coordinate(i/int(b.Size())+1))
	}

	rand.Shuffle(totalTiles, func(i, j int) {
		activeIndexes[i], activeIndexes[j] = activeIndexes[j], activeIndexes[i]
	})

	for i := range activeIndexes[:count] {
		tile, _ := b.TileAt(activeIndexes[i])
		tile.Blackhole(true)
	}
}
