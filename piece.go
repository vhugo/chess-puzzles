package main

import (
	"strings"

	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
)

// Piece represent a single chess piece
type Piece struct {
	*tl.Text
	piece chess.Piece
}

func (p Piece) String() string {
	s := " "
	if p.piece.Color() == chess.Black {
		s = "*"
	}

	switch p.piece.Type() {
	case chess.Pawn:
		return s + "P"
	default:
		return s + strings.ToUpper(p.piece.Type().String())
	}
}

// Color returns the chess piece color. Either black or white
func (p Piece) Color() tl.Attr {
	if p.piece.Color() == chess.Black {
		return palette.black
	}
	return palette.white
}

// NewPiece return a representation of a chess piece
func NewPiece(x, y int, p chess.Piece) *Piece {
	piece := &Piece{piece: p}
	piece.Text = tl.NewText(x+2, y, piece.String(), piece.Color(), tl.ColorDefault)
	return piece
}
