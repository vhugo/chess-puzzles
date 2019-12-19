package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
)

type BoardSquare struct {
	*tl.Text
	square chess.Square
}

var (
	A1 = &BoardSquare{square: chess.A1}
	A2 = &BoardSquare{square: chess.A2}
	A3 = &BoardSquare{square: chess.A3}
	A4 = &BoardSquare{square: chess.A4}
	A5 = &BoardSquare{square: chess.A5}
	A6 = &BoardSquare{square: chess.A6}
	A7 = &BoardSquare{square: chess.A7}
	A8 = &BoardSquare{square: chess.A8}
	B1 = &BoardSquare{square: chess.B1}
	B2 = &BoardSquare{square: chess.B2}
	B3 = &BoardSquare{square: chess.B3}
	B4 = &BoardSquare{square: chess.B4}
	B5 = &BoardSquare{square: chess.B5}
	B6 = &BoardSquare{square: chess.B6}
	B7 = &BoardSquare{square: chess.B7}
	B8 = &BoardSquare{square: chess.B8}
	C1 = &BoardSquare{square: chess.C1}
	C2 = &BoardSquare{square: chess.C2}
	C3 = &BoardSquare{square: chess.C3}
	C4 = &BoardSquare{square: chess.C4}
	C5 = &BoardSquare{square: chess.C5}
	C6 = &BoardSquare{square: chess.C6}
	C7 = &BoardSquare{square: chess.C7}
	C8 = &BoardSquare{square: chess.C8}
	D1 = &BoardSquare{square: chess.D1}
	D2 = &BoardSquare{square: chess.D2}
	D3 = &BoardSquare{square: chess.D3}
	D4 = &BoardSquare{square: chess.D4}
	D5 = &BoardSquare{square: chess.D5}
	D6 = &BoardSquare{square: chess.D6}
	D7 = &BoardSquare{square: chess.D7}
	D8 = &BoardSquare{square: chess.D8}
	E1 = &BoardSquare{square: chess.E1}
	E2 = &BoardSquare{square: chess.E2}
	E3 = &BoardSquare{square: chess.E3}
	E4 = &BoardSquare{square: chess.E4}
	E5 = &BoardSquare{square: chess.E5}
	E6 = &BoardSquare{square: chess.E6}
	E7 = &BoardSquare{square: chess.E7}
	E8 = &BoardSquare{square: chess.E8}
	F1 = &BoardSquare{square: chess.F1}
	F2 = &BoardSquare{square: chess.F2}
	F3 = &BoardSquare{square: chess.F3}
	F4 = &BoardSquare{square: chess.F4}
	F5 = &BoardSquare{square: chess.F5}
	F6 = &BoardSquare{square: chess.F6}
	F7 = &BoardSquare{square: chess.F7}
	F8 = &BoardSquare{square: chess.F8}
	G1 = &BoardSquare{square: chess.G1}
	G2 = &BoardSquare{square: chess.G2}
	G3 = &BoardSquare{square: chess.G3}
	G4 = &BoardSquare{square: chess.G4}
	G5 = &BoardSquare{square: chess.G5}
	G6 = &BoardSquare{square: chess.G6}
	G7 = &BoardSquare{square: chess.G7}
	G8 = &BoardSquare{square: chess.G8}
	H1 = &BoardSquare{square: chess.H1}
	H2 = &BoardSquare{square: chess.H2}
	H3 = &BoardSquare{square: chess.H3}
	H4 = &BoardSquare{square: chess.H4}
	H5 = &BoardSquare{square: chess.H5}
	H6 = &BoardSquare{square: chess.H6}
	H7 = &BoardSquare{square: chess.H7}
	H8 = &BoardSquare{square: chess.H8}
)

var files = [8]chess.File{
	chess.FileA, chess.FileB, chess.FileC, chess.FileD,
	chess.FileE, chess.FileF, chess.FileG, chess.FileH,
}

var ranks = [8]chess.Rank{
	chess.Rank8, chess.Rank7, chess.Rank6, chess.Rank5,
	chess.Rank4, chess.Rank3, chess.Rank2, chess.Rank1,
}

var board = [8][8]*BoardSquare{
	[8]*BoardSquare{A8, B8, C8, D8, E8, F8, G8, H8},
	[8]*BoardSquare{A7, B7, C7, D7, E7, F7, G7, H7},
	[8]*BoardSquare{A6, B6, C6, D6, E6, F6, G6, H6},
	[8]*BoardSquare{A5, B5, C5, D5, E5, F5, G5, H5},
	[8]*BoardSquare{A4, B4, C4, D4, E4, F4, G4, H4},
	[8]*BoardSquare{A3, B3, C3, D3, E3, F3, G3, H3},
	[8]*BoardSquare{A2, B2, C2, D2, E2, F2, G2, H2},
	[8]*BoardSquare{A1, B1, C1, D1, E1, F1, G1, H1},
}

func printPiece(r chess.Piece) string {
	var p string
	switch r.Type() {
	case chess.King:
		p = "K"
	case chess.Queen:
		p = "Q"
	case chess.Rook:
		p = "R"
	case chess.Bishop:
		p = "B"
	case chess.Knight:
		p = "N"
	case chess.Pawn:
		p = "P"
	}

	if len(p) == 0 {
		return ""
	}

	if r.Color() == chess.Black {
		return fmt.Sprintf(" *%s", p)
	}
	return fmt.Sprintf("  %s", p)
}

func updateBoard() {
	for r := range board {
		for _, sq := range board[r] {
			piece := gc.Position().Board().Piece(sq.square)
			sq.SetText(printPiece(piece))
		}
	}
}
