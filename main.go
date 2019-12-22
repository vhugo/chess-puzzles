package main

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
	"github.com/vhugo/chess-puzzles/puzzle"
)

var gc *chess.Game
var game *tl.Game
var puzzler puzzle.Puzzler

func main() {
	pz, err := loadPuzzle()
	if err != nil {
		panic(err)
	}
	gc = chess.NewGame(pz)
	game = tl.NewGame()
	game.Screen().SetFps(30)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorDefault,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	// board
	for r := range ranks {
		for f := range make([]bool, 8) {
			x := layout.board.marginLeft + (f * layout.square.w)
			y := layout.board.marginTop + (r * layout.square.h)
			level.AddEntity(
				NewSquare(x, y, layout.square.w, layout.square.h,
					palette.alternate(), uci(files[f], ranks[r])))
		}
		palette.alternate()
	}

	// pieces on the board
	bsquares := board.Squares()
	for r := range bsquares {
		for f, sq := range bsquares[r] {
			piece := gc.Position().Board().Piece(sq.square)
			x := layout.piece.marginLeft + (f * layout.square.w)
			y := layout.piece.marginTop + (r * layout.square.h)
			sq.Text = NewPiece(x, y, piece).Text
			level.AddEntity(sq)
		}
	}

	// coordinates
	for r, rank := range ranks {
		x := layout.ranks.marginLeft
		y := layout.ranks.marginTop + (r * layout.square.h)
		level.AddEntity(NewCoordinate(x, y, rank))
	}

	for f, file := range files {
		x := layout.files.marginLeft + layout.board.marginLeft + (f * layout.square.w)
		y := layout.files.marginTop + (layout.board.h * layout.square.h)
		level.AddEntity(NewCoordinate(x, y, file))
	}

	// notations
	notation = NewNotation(
		layout.notation.marginLeft+(layout.board.w*layout.square.w),
		layout.notation.marginTop, tl.ColorWhite)
	level.AddEntity(notation)

	// user interface
	player = NewPlayer(
		layout.input.marginLeft+layout.board.marginLeft,
		layout.input.marginTop+(layout.board.h*layout.square.h),
		palette.input)
	level.AddEntity(player)

	status = NewStatus(
		layout.input.marginLeft+layout.board.marginLeft,
		layout.input.marginTop+(layout.board.h*layout.square.h)+1)
	level.AddEntity(status)

	score = NewScore(
		layout.input.marginLeft+layout.board.marginLeft+(layout.board.w*layout.square.w),
		layout.input.marginTop+(layout.board.h*layout.square.h)+1)
	level.AddEntity(score)

	game.Screen().SetLevel(level)
	game.Start()
}
