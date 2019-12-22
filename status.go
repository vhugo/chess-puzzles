package main

import tl "github.com/JoelOtter/termloop"

type Status struct {
	*tl.Text
}

var status *Status

func (s *Status) Update(text string, color tl.Attr) {
	s.SetText(text)
	s.SetColor(palette.input, color)
}

func NewStatus(x, y int) *Status {
	s := new(Status)
	s.Text = tl.NewText(x, y, " ",
		palette.input, tl.ColorDefault)
	return s
}
