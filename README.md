A Proxx game
============

# Rules

> :information_source:
>
> The game works best with [monospaced font](https://en.wikipedia.org/wiki/Monospaced_font) in your terminal.

The game is happening on a board of NxN tiles.

At the beginning, all tiles are unrevealed and displayed with caption `'?'`.
Tile might have a black hole behind it, it might not. 

When player opens a tile without blackhole behind it, the caption of the tile is updated with a number of blackholes around, including diagonal tiles.

Player wins when all the tiles without blackhole behind them are opened.

Player loses when clicks on tile with blackhole.

# Building the project

Requirements:
- [Go](https://go.dev/) with of version `'>= 1.20'`
- (optional) `make` utility

Build with `go` compiler:
```sh
$ go build ./cmd/...
```

or using `make` target:
```sh
$ make build
```

To see all available targets, run:
```sh
$ make help
```

Before submitting new PR, ensure that unit-tests and linter report no errors by running:
```sh
$ make lint test
```
