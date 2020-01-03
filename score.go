package main

import tl "github.com/JoelOtter/termloop"

// Score represents the score of the game
type Score struct {
	*tl.Text
}

const (
	succeed = "succeed"
	failed  = "failed"
)

var score *Score

// Update updates the score on the user interface
func (s *Score) Update(text string, color tl.Attr) {
	s.SetText("  " + text + "  ")
	s.SetColor(palette.input, color)
}

// NewScore returns a new instance of score
func NewScore(x, y int) *Score {
	s := new(Score)
	s.Text = tl.NewText(x, y, "", palette.input, tl.ColorDefault)
	return s
}
