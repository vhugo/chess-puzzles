package puzzle

import (
	"fmt"

	"github.com/notnil/chess"
)

// Puzzler represents an adapter that can return puzzle from different sources.
type Puzzler interface {
	NewGame() (func(*chess.Game), error)
	Answer(*chess.Move) bool
	NextMove() *chess.Move
	Hint() *chess.Move
	Score() Score
	Done() bool
}

type Score int8

// Source represents the origin of requested puzzles.
type Source int

const (
	_           Source = iota
	CHESSDOTCOM        // chess.com
)

const (
	NOSCORE Score = iota
	SUCCESS
	FAILURE
)

// New return an instance for puzzle of a particular source.
func New(source Source) (Puzzler, error) {
	switch source {
	case CHESSDOTCOM:
		return NewChessDotCom(ChessDotComURL, ChessDotComTimeout, ChessDotComExpires)

	default:
		return nil, fmt.Errorf("unknown source")
	}
}
