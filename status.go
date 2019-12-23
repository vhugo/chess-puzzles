package main

import tl "github.com/JoelOtter/termloop"

type Status struct {
	*tl.Text
}

const (
	unsolved = "unsolved"
)

var status *Status

func (s *Status) Update(text string, color tl.Attr) {
	s.SetText(" " + text + " ")
	s.SetColor(palette.input, color)
	x, y := s.Position()
	x += len(unsolved) - len(text)
	s.SetPosition(x, y)
}

func NewStatus(x, y int) *Status {
	s := new(Status)
	s.Text = tl.NewText(x-len(unsolved)+2, y, " "+unsolved+" ",
		palette.input, tl.ColorDefault)
	return s
}
