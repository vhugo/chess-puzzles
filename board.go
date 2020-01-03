package main

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
)

// BoardSquare represents a single chess board's square
type BoardSquare struct {
	*tl.Text
	square chess.Square
}

// Board represents the chess board in the interface
type Board struct {
	board [8][8]*BoardSquare
}

var (
	a1 = &BoardSquare{square: chess.A1}
	a2 = &BoardSquare{square: chess.A2}
	a3 = &BoardSquare{square: chess.A3}
	a4 = &BoardSquare{square: chess.A4}
	a5 = &BoardSquare{square: chess.A5}
	a6 = &BoardSquare{square: chess.A6}
	a7 = &BoardSquare{square: chess.A7}
	a8 = &BoardSquare{square: chess.A8}
	b1 = &BoardSquare{square: chess.B1}
	b2 = &BoardSquare{square: chess.B2}
	b3 = &BoardSquare{square: chess.B3}
	b4 = &BoardSquare{square: chess.B4}
	b5 = &BoardSquare{square: chess.B5}
	b6 = &BoardSquare{square: chess.B6}
	b7 = &BoardSquare{square: chess.B7}
	b8 = &BoardSquare{square: chess.B8}
	c1 = &BoardSquare{square: chess.C1}
	c2 = &BoardSquare{square: chess.C2}
	c3 = &BoardSquare{square: chess.C3}
	c4 = &BoardSquare{square: chess.C4}
	c5 = &BoardSquare{square: chess.C5}
	c6 = &BoardSquare{square: chess.C6}
	c7 = &BoardSquare{square: chess.C7}
	c8 = &BoardSquare{square: chess.C8}
	d1 = &BoardSquare{square: chess.D1}
	d2 = &BoardSquare{square: chess.D2}
	d3 = &BoardSquare{square: chess.D3}
	d4 = &BoardSquare{square: chess.D4}
	d5 = &BoardSquare{square: chess.D5}
	d6 = &BoardSquare{square: chess.D6}
	d7 = &BoardSquare{square: chess.D7}
	d8 = &BoardSquare{square: chess.D8}
	e1 = &BoardSquare{square: chess.E1}
	e2 = &BoardSquare{square: chess.E2}
	e3 = &BoardSquare{square: chess.E3}
	e4 = &BoardSquare{square: chess.E4}
	e5 = &BoardSquare{square: chess.E5}
	e6 = &BoardSquare{square: chess.E6}
	e7 = &BoardSquare{square: chess.E7}
	e8 = &BoardSquare{square: chess.E8}
	f1 = &BoardSquare{square: chess.F1}
	f2 = &BoardSquare{square: chess.F2}
	f3 = &BoardSquare{square: chess.F3}
	f4 = &BoardSquare{square: chess.F4}
	f5 = &BoardSquare{square: chess.F5}
	f6 = &BoardSquare{square: chess.F6}
	f7 = &BoardSquare{square: chess.F7}
	f8 = &BoardSquare{square: chess.F8}
	g1 = &BoardSquare{square: chess.G1}
	g2 = &BoardSquare{square: chess.G2}
	g3 = &BoardSquare{square: chess.G3}
	g4 = &BoardSquare{square: chess.G4}
	g5 = &BoardSquare{square: chess.G5}
	g6 = &BoardSquare{square: chess.G6}
	g7 = &BoardSquare{square: chess.G7}
	g8 = &BoardSquare{square: chess.G8}
	h1 = &BoardSquare{square: chess.H1}
	h2 = &BoardSquare{square: chess.H2}
	h3 = &BoardSquare{square: chess.H3}
	h4 = &BoardSquare{square: chess.H4}
	h5 = &BoardSquare{square: chess.H5}
	h6 = &BoardSquare{square: chess.H6}
	h7 = &BoardSquare{square: chess.H7}
	h8 = &BoardSquare{square: chess.H8}

	board = Board{
		board: [8][8]*BoardSquare{
			{a8, b8, c8, d8, e8, f8, g8, h8},
			{a7, b7, c7, d7, e7, f7, g7, h7},
			{a6, b6, c6, d6, e6, f6, g6, h6},
			{a5, b5, c5, d5, e5, f5, g5, h5},
			{a4, b4, c4, d4, e4, f4, g4, h4},
			{a3, b3, c3, d3, e3, f3, g3, h3},
			{a2, b2, c2, d2, e2, f2, g2, h2},
			{a1, b1, c1, d1, e1, f1, g1, h1},
		}}
)

// Squares returns the 8x8 representation of the chess board
func (b Board) Squares() [8][8]*BoardSquare {
	return b.board
}

// Update updates the squares with the correct chess piece
func (b Board) Update() {
	for r := range b.board {
		for _, sq := range b.board[r] {
			piece := Piece{
				piece: gc.Position().Board().Piece(sq.square),
			}
			sq.Text.SetColor(piece.Color(), tl.ColorDefault)
			sq.SetText(piece.String())
		}
	}
}
