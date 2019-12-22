package main

import tl "github.com/JoelOtter/termloop"

type Score struct {
	*tl.Text
}

var score *Score

func (s *Score) Update(text string, color tl.Attr) {
	s.SetText(text)
	s.SetColor(palette.input, color)
}

func NewScore(x, y int) *Score {
	s := new(Score)
	s.Text = tl.NewText(x-len("unsolved"), y, "unsolved",
		palette.input, tl.ColorDefault)
	return s
}
