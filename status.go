package main

import tl "github.com/JoelOtter/termloop"

type Status struct {
	*tl.Text
}

const (
	unsolved = "unsolved"
	solved   = "solved"
)

var status *Status

func (s *Status) Update(text string, color tl.Attr) {
	if s == nil {
		return
	}

	x, y := s.Position()
	w, _ := s.Size()
	x += w

	newText := " " + text + " "

	s.SetText(newText)
	s.SetPosition(x-len(newText), y)
	s.SetColor(palette.input, color)
}

func NewStatus(x, y int) *Status {
	s := new(Status)
	s.Text = tl.NewText(x, y, "", palette.input, tl.ColorDefault)
	s.Update(unsolved, tl.ColorDefault)
	return s
}
