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
	if len(player.Input()) < 2 {
		s.SetColor(s.color)
		return
	}

	input := player.Input()
	if input[:2] == s.loc || (len(input) >= 4 && input[2:4] == s.loc) {
		// once puzzle is done no more moves are allowed.
		if puzzler != nil && puzzler.Done() {
			s.SetColor(palette.invalid)
			return
		}

		for _, m := range gc.ValidMoves() {
			if matchMove(m.String(), player.Input(), s.loc) {
				s.SetColor(palette.valid)
				return
			}
		}
		s.SetColor(palette.invalid)
		return
	}

	s.SetColor(s.color)
}

func matchMove(valid, input, loc string) bool {
	return (len(input) >= 4 && input[:4] == valid[:4] && input[2:4] == loc) ||
		(len(input) >= 2 && input[:2] == valid[:2] && input[:2] == loc)
}

func NewSquare(x, y, w, h int, color tl.Attr, loc string) *Square {
	s := &Square{
		loc:   loc,
		color: color,
	}
	s.Rectangle = tl.NewRectangle(x, y, w, h, color)

	return s
}
