package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
	wrap "github.com/mitchellh/go-wordwrap"
)

type NotationStr interface {
	String() string
}

type Notation struct {
	*tl.Entity
	x, y int
}

var notation *Notation

func (n *Notation) Update() {
	n.Entity = tl.NewEntityFromCanvas(
		n.x, n.y, tl.CanvasFromString(
			wrap.WrapString(gc.String(), 12),
		),
	)
}

func NewNotation(x, y int, color tl.Attr) *Notation {
	n := new(Notation)
	n.x, n.y = x, y
	n.Entity = tl.NewEntityFromCanvas(n.x, n.y, tl.CanvasFromString(" "))
	return n
}

func printNotation(n NotationStr) string {
	return fmt.Sprintf("   %s", n)
}
