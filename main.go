package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
)

var gc *chess.Game
var game *tl.Game

func main() {
	gc = chess.NewGame()
	game = tl.NewGame()

	// game.SetDebugOn(true)
	game.Screen().SetFps(30)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorDefault,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	// game.Log("chess board:", gc.Position().Board().SquareMap())

	// board
	makeSquare := func(x, y int, loc string) *Square {
		return NewSquare(x, y, square.w, square.h, palette.alternate(), loc)
	}
	for r := range ranks {
		for f := range board[r] {
			level.AddEntity(
				makeSquare((f+1)*square.w, r*square.h, fmt.Sprintf("%s%s", files[f], ranks[r])))
		}
		palette.alternate()
	}

	// pieces on the board
	for r := range board {
		for f, sq := range board[r] {
			piece := gc.Position().Board().Piece(sq.square)

			sq.Text = tl.NewText(((f+1)*square.w)+1, (r*square.h)+1, printPiece(piece),
				palette.pieces, tl.ColorDefault)
			level.AddEntity(sq)
		}
		palette.alternate()
	}

	// notations
	for r, rank := range ranks {
		level.AddEntity(
			tl.NewText(1, (r*square.h + 1), printNotation(rank),
				palette.notations, tl.ColorDefault))
	}

	for f, file := range files {
		level.AddEntity(
			tl.NewText((f+1)*square.w, 8*square.h, printNotation(file),
				palette.notations, tl.ColorDefault))
	}
	notation = NewNotation((9*square.w)+1, 0, tl.ColorWhite)
	level.AddEntity(notation)

	// user interface
	prefix := "  your move: "
	level.AddEntity(tl.NewText(0, (9*square.h)-1, prefix, tl.ColorWhite, tl.ColorDefault))
	player = NewPlayer(len(prefix), (9*square.h)-1, tl.ColorWhite)
	level.AddEntity(player)

	game.Screen().SetLevel(level)
	game.Start()
}
