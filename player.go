package main

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
	"github.com/vhugo/chess-puzzles/puzzle"
	"strings"
)

const (
	cmdNew = "!new"
)

type Player struct {
	*tl.Text
	previousMove string
	move         string
	prefix       string
}

var player *Player

func (p *Player) Tick(e tl.Event) {

	if e.Type == tl.EventKey {
		switch e.Key {
		case tl.KeyBackspace2, tl.KeyBackspace:
			if len(p.move) == 0 {
				return
			}
			p.move = p.move[:len(p.move)-1]

		case tl.KeyCtrlU:
			p.move = ""

		case tl.KeyEnter:
			if len(p.move) == 0 {
				return
			}

			switch p.move[0] {
			case '!':
				switch strings.ToLower(p.move) {
				case cmdNew:
					pz, err := loadPuzzle()
					if err != nil {
						panic(err)
					}
					gc = chess.NewGame(pz)
					notation.Update()
					board.Update()

					score.Update("", tl.ColorDefault)
					status.Update(unsolved, tl.ColorDefault)
				}

			default:
				move(p.move)

				p.previousMove = p.move
				if puzzler != nil {
					switch {
					case puzzler.Score() == puzzle.SUCCESS:
						score.Update(succeed, tl.RgbTo256Color(0, 100, 0))

					case puzzler.Score() == puzzle.FAILURE:
						score.Update(failed, tl.RgbTo256Color(100, 0, 0))
						p.previousMove = ""
					}

					if puzzler.Done() {
						status.Update(solved, tl.RgbTo256Color(120, 100, 0))
					}
				}
			}
			p.move = ""

		default:
			p.move = p.move + string(e.Ch)
		}

		p.SetText(gc.Position().Turn().Name() + p.prefix + p.move)
	}
}

func NewPlayer(x, y int, color tl.Attr) *Player {
	p := new(Player)
	p.prefix = " to move: "
	p.Text = tl.NewText(x, y,
		gc.Position().Turn().Name()+p.prefix, color, tl.ColorDefault)
	return p
}

func move(m string) error {
	move, err := chess.LongAlgebraicNotation{}.Decode(gc.Position(), m)
	if err != nil {
		return err
	}
	if puzzler != nil && !puzzler.Answer(move) {
		return nil
	}

	gc.Move(move)
	notation.Update()
	board.Update()

	if puzzler != nil && puzzler.Score() != puzzle.SUCCESS {
		gc.Move(puzzler.NextMove())
		notation.Update()
		board.Update()
	}

	return nil
}

func loadPuzzle() (func(*chess.Game), error) {
	var err error

	puzzler, err = puzzle.New(puzzle.CHESSDOTCOM)
	if err != nil {
		return nil, err
	}

	return puzzler.NewGame()
}
