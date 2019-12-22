package main

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
	"strings"
)

const (
	cmdNew = "!new"
)

type Player struct {
	*tl.Text
	buffer string
	prefix string
}

var player *Player

func (p *Player) Input() string {
	return p.buffer
}

func (p *Player) Tick(e tl.Event) {
	if e.Type == tl.EventKey {
		switch e.Key {
		case tl.KeyBackspace2, tl.KeyBackspace:
			if len(p.buffer) == 0 {
				return
			}
			p.buffer = p.buffer[:len(p.buffer)-1]

		case tl.KeyEnter:
			if len(p.buffer) == 0 {
				return
			}

			switch p.buffer[0] {
			case '!':
				switch strings.ToLower(p.buffer) {
				case cmdNew:
					gc = chess.NewGame()
					notation.Update()
					board.Update()
				}
				// game.Log("command caller", p.buffer)

			default:
				move(p.buffer)
				// if err := move(p.buffer); err != nil {
				// 	game.Log("move error: ", err)
				// }
			}

			p.buffer = ""

		default:
			p.buffer = p.buffer + string(e.Ch)
		}

		p.SetText(gc.Position().Turn().Name() + p.prefix + p.buffer)
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
	gc.Move(move)
	notation.Update()
	board.Update()
	// game.Log("chess board: ", gc.Position().Board().Draw())

	return nil
}
