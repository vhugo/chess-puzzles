package main

import (
	"strings"

	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
)

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

func NewPiece(x, y int, p chess.Piece) *Piece {
	piece := &Piece{piece: p}
	piece.Text = tl.NewText(x+2, y, piece.String(), palette.pieces, tl.ColorDefault)
	return piece
}
