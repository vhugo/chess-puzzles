package main

import (
	"testing"

	"github.com/notnil/chess"
)

func TestBoard(t *testing.T) {
	t.Run("print piece", func(t *testing.T) {
		for _, tc := range []struct {
			piece    chess.Piece
			expected string
		}{
			{chess.BlackRook, " *R"},
			{chess.BlackKnight, " *N"},
			{chess.BlackBishop, " *B"},
			{chess.BlackKing, " *K"},
			{chess.BlackQueen, " *Q"},
			{chess.BlackPawn, " *P"},
			{chess.WhiteRook, "  R"},
			{chess.WhiteKnight, "  N"},
			{chess.WhiteBishop, "  B"},
			{chess.WhiteKing, "  K"},
			{chess.WhiteQueen, "  Q"},
			{chess.WhitePawn, "  P"},
			{chess.NoPiece, ""},
		} {
			if printPiece(tc.piece) != tc.expected {
				t.Fatalf("got %q, expected %q", printPiece(tc.piece), tc.expected)
			}
		}
	})
}
