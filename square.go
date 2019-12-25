package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Square struct {
	*tl.Rectangle
	loc   string
	color tl.Attr
}

func (s *Square) Tick(e tl.Event) {
	move := player.move
	if (len(move) >= 2 && move[:2] == s.loc) ||
		(len(move) >= 4 && move[2:4] == s.loc) {

		// once puzzle is done no more moves are allowed.
		if puzzler != nil && puzzler.Done() {
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
	if len(previousMove) >= 4 &&
		(previousMove[:2] == s.loc || previousMove[2:4] == s.loc) {
		s.SetColor(palette.moved)
		return
	}

	s.SetColor(s.color)
}

func matchMove(valid, move, loc string) bool {
	return (len(move) >= 4 && move[:4] == valid[:4] && move[2:4] == loc) ||
		(len(move) >= 2 && move[:2] == valid[:2] && move[:2] == loc)
}

func NewSquare(x, y, w, h int, color tl.Attr, loc string) *Square {
	s := &Square{
		loc:   loc,
		color: color,
	}
	s.Rectangle = tl.NewRectangle(x, y, w, h, color)

	return s
}
