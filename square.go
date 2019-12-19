package main

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
)

type Square struct {
	*tl.Rectangle
	loc   string
	color tl.Attr
}

type SquarePalette struct {
	dark, light, pieces, notations, current tl.Attr
	valid, invalid                          tl.Attr
}

func (s SquarePalette) alternate() tl.Attr {
	if palette.current == palette.light {
		palette.current = palette.dark
		return palette.current
	}
	palette.current = palette.light
	return palette.current
}

var (
	palette = SquarePalette{
		dark:      tl.ColorCyan,
		light:     tl.ColorWhite,
		pieces:    tl.ColorBlack,
		notations: tl.ColorCyan,
		current:   tl.ColorDefault,
		valid:     tl.ColorYellow,
		invalid:   tl.ColorRed,
	}

	square = struct {
		w, h int
	}{
		w: 7,
		h: 3,
	}
)

func (s *Square) Tick(e tl.Event) {
	if len(player.Input()) < 2 {
		s.SetColor(s.color)
		return
	}

	if player.Input()[:2] == s.loc || player.Input()[2:] == s.loc {

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

func getSquare(f chess.File, r chess.Rank) chess.Square {
	return chess.Square((int(r) * 8) + int(f))
}
