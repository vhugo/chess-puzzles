package main

import tl "github.com/JoelOtter/termloop"

type Score struct {
	*tl.Text
}

const (
	succeed = "succeed"
	failed  = "failed"
)

var score *Score

func (s *Score) Update(text string, color tl.Attr) {
	s.SetText("  " + text + "  ")
	s.SetColor(palette.input, color)
}

func NewScore(x, y int) *Score {
	s := new(Score)
	s.Text = tl.NewText(x, y, "", palette.input, tl.ColorDefault)
	return s
}
