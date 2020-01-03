package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Square represents a chess board's square
type Square struct {
	*tl.Rectangle
	loc   string
	color tl.Attr
}

// Tick handles events and update squares accordingly
func (s *Square) Tick(e tl.Event) {
	move := player.move
	if possibleMove(move, s.loc) {

		// once puzzle is done no more moves are allowed.
		if puzzleIsDone() {
			s.SetColor(palette.invalid)
			return
		}

		for _, m := range gc.ValidMoves() {
			if matchMove(m.String(), move, s.loc) {
				s.SetColor(palette.valid)
				return
			}
		}
		s.SetColor(palette.invalid)
		return
	}

	previousMove := player.previousMove
	if matchPreviousMove(previousMove, s.loc) {
		s.SetColor(palette.moved)
		return
	}

	s.SetColor(s.color)
}

func possibleMove(move, loc string) bool {
	return (len(move) >= 2 && move[:2] == loc) ||
		(len(move) >= 4 && move[2:4] == loc)
}

func matchMove(valid, move, loc string) bool {
	return (len(move) >= 4 && move[:4] == valid[:4] && move[2:4] == loc) ||
		(len(move) >= 2 && move[:2] == valid[:2] && move[:2] == loc)
}

func matchPreviousMove(previousMove, loc string) bool {
	return len(previousMove) >= 4 &&
		(previousMove[:2] == loc || previousMove[2:4] == loc)
}

func puzzleIsDone() bool {
	return puzzler != nil && puzzler.Done()
}

// NewSquare returns a termloop new rectangle used to draw chess board's squares
func NewSquare(x, y, w, h int, color tl.Attr, loc string) *Square {
	s := &Square{
		loc:   loc,
		color: color,
	}
	s.Rectangle = tl.NewRectangle(x, y, w, h, color)

	return s
}
