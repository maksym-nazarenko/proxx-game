package app

import (
	"errors"
	"fmt"

	"github.com/maksym-nazarenko/proxx-game/internal/game"
)

const (
	// Minimum and maximum supported board size
	MinBoardSize = 3
	MaxBoardSize = 99
)

type Configuration struct {
	BoardSize       game.Coordinate
	BlackholesCount game.BoardArea
}

func ValidateConfig(config Configuration) error {
	if config.BoardSize < MinBoardSize || config.BoardSize > MaxBoardSize {
		return fmt.Errorf("invalid board size: supported values are in range %d..%d", MinBoardSize, MaxBoardSize)
	}
	if config.BlackholesCount >= game.BoardArea(config.BoardSize*config.BoardSize) {
		return errors.New("number of blackholes cannot be more than TotalTilesCount-1")
	}

	return nil
}
