package game

var DefaultBlackholesPlaceStrategy BlackholesPlacerFunc = defaultBlackholesPlaceStrategy

func defaultBlackholesPlaceStrategy(b Board, count BoardArea) {
	tile, _ := b.TileAt(NewPoint(1, 2))
	tile.Blackhole(true)
}
