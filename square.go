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

	if player.Input()[:2] == s.loc || player.Input()[2:] == s.loc {
		// once puzzle is done no more moves are allowed.
		if puzzler != nil && puzzler.Done() {
			s.SetColor(palette.invalid)
			return
		}

		for _, m := range gc.ValidMoves() {
			if (m.String() == player.Input() && len(player.Input()) == 4) ||
				(player.Input()[:2] == m.String()[:2] && player.Input()[:2] == s.loc) ||
				(player.Input()[2:] == m.String()[2:] && player.Input()[2:] == s.loc) {

				s.SetColor(palette.valid)
				return
			}
		}
		s.SetColor(palette.invalid)
		return
	}

	s.SetColor(s.color)
}

func NewSquare(x, y, w, h int, color tl.Attr, loc string) *Square {
	s := &Square{
		loc:   loc,
		color: color,
	}
	s.Rectangle = tl.NewRectangle(x, y, w, h, color)

	return s
}
