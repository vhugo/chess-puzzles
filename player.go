package main

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/notnil/chess"
	"github.com/vhugo/chess-puzzles/puzzle"
	"strings"
)

const (
	cmdNew  = "!new"
	cmdHint = "!hint"
)

// Player represents a player in the game
type Player struct {
	*tl.Text
	previousMove string
	move         string
	prefix       string
}

var player *Player

// Tick handle events triggered by the player
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
			p.enterCommand()

		default:
			p.move = p.move + string(e.Ch)
		}

		p.SetText(gc.Position().Turn().Name() + p.prefix + p.move)
	}
}

// NewPlayer returns a new instance of player
func NewPlayer(x, y int, color tl.Attr) *Player {
	p := new(Player)
	p.prefix = " to move: "
	p.Text = tl.NewText(x, y,
		gc.Position().Turn().Name()+p.prefix, color, tl.ColorDefault)
	return p
}

func (p *Player) enterCommand() {
	switch p.move[0] {
	case '!':
		switch strings.ToLower(p.move) {
		case cmdNew:
			p.cmdNewPuzzle()

		case cmdHint:
			nextMove := puzzler.Hint()
			if nextMove == nil {
				return
			}
			p.cmdHint(nextMove.String())
		}

	default:
		p.previousMove = p.move
		move(p.move)
		p.updateScore()
		p.updateStatus()
	}
	p.move = ""
}

func (p *Player) cmdNewPuzzle() {
	pz, err := loadPuzzle()
	if err != nil {
		panic(err)
	}
	gc = chess.NewGame(pz)
	notation.Update()
	board.Update()

	p.previousMove = ""
	p.move = ""

	score.Update("", tl.ColorDefault)
	status.Update(unsolved, tl.ColorDefault)
}

func (p *Player) cmdHint(nextMove string) {
	puzzler.Answer(nil) // fails the puzzle
	p.previousMove = nextMove
	score.Update(failed, tl.RgbTo256Color(100, 0, 0))
}

func (p *Player) updateScore() {
	if puzzler != nil {
		switch {
		case puzzler.Score() == puzzle.SUCCESS:
			score.Update(succeed, tl.RgbTo256Color(0, 100, 0))

		case puzzler.Score() == puzzle.FAILURE:
			score.Update(failed, tl.RgbTo256Color(100, 0, 0))
		}
	}
}

func (p *Player) updateStatus() {
	if puzzler != nil {
		if puzzler.Done() {
			status.Update(solved, tl.RgbTo256Color(120, 100, 0))
		}
	}
}

func move(m string) error {
	move, err := chess.LongAlgebraicNotation{}.Decode(gc.Position(), m)
	if err != nil {
		return err
	}
	if puzzler != nil && !puzzler.Answer(move) {
		player.previousMove = ""
		return nil
	}

	gc.Move(move)
	notation.Update()
	board.Update()

	if puzzler != nil && puzzler.Score() != puzzle.SUCCESS {
		nextMove := puzzler.NextMove()
		if nextMove == nil {
			return nil
		}

		gc.Move(nextMove)
		player.previousMove = nextMove.String()
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
