package main

import (
	"strings"

	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
)

// Notation represets object block responsible for displaying games notation
type Notation struct {
	*tl.Entity
	x, y int
}

// Coordinate represents the area for ranks and files indication outside the
//  board area.
type Coordinate struct {
	*tl.Text
}

var (
	notation *Notation

	files = [8]chess.File{
		chess.FileA, chess.FileB, chess.FileC, chess.FileD,
		chess.FileE, chess.FileF, chess.FileG, chess.FileH,
	}

	ranks = [8]chess.Rank{
		chess.Rank8, chess.Rank7, chess.Rank6, chess.Rank5,
		chess.Rank4, chess.Rank3, chess.Rank2, chess.Rank1,
	}
)

// Update display the new move notation in the reserved area
func (n *Notation) Update() {
	n.Entity = tl.NewEntityFromCanvas(
		n.x, n.y, tl.CanvasFromString(
			wrapNotation(gc.String()),
		),
	)
}

// NewNotation create a new area for notations
func NewNotation(x, y int, color tl.Attr) *Notation {
	n := new(Notation)
	n.x, n.y = x, y
	n.Entity = tl.NewEntityFromCanvas(n.x, n.y, tl.CanvasFromString(" "))
	return n
}

// NewCoordinate create each ranks and/or file displayed on the side of the board
func NewCoordinate(x, y int, c interface{ String() string }) *Coordinate {
	var coordinate Coordinate
	coordinate.Text = tl.NewText(x, y, c.String(), palette.notations, tl.ColorDefault)
	return &coordinate
}

func wrapNotation(s string) string {
	var str string
	for i, n := range strings.Fields(s) {
		if i != 0 && i%2 == 0 {
			str += "\n"
		}
		str += n + " "
	}
	return str
}

// uci returns Universal Chess Interface
func uci(f chess.File, r chess.Rank) string {
	return f.String() + r.String()
}
