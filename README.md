# Chess Puzzles

A terminal-based chess puzzles interface. The inspiration behind it comes from
the recent enjoyment of solving chess puzzles at
[chess.com](https://www.chess.com/puzzles).

This was built in an afternoon as an experiment to see if it was possible and it
was only possible because of two awesome projects
[termloop](https://github.com/JoelOtter/termloop) as the game engine and
[chess](https://github.com/notnil/chess) with all the chess mechanics.

Needless to say, this is a work in progress, the game starts and you can enter
[long algebraic
notation](https://en.wikipedia.org/wiki/Algebraic_notation_(chess)#Long_algebraic_notation)
(e.g. e2e4) which makes the pieces to move on the board, however, puzzles are
not being loaded yet. It is next on my list.

## Getting started

A binary version is still on the making so until then you will need
[Golang](https://golang.org/doc/install) to compile and run the source yourself.
After installing Golang following the sequence of commands:

```bash
go get -u github.com/vhugo/chess-puzzles.git
go install github.com/vhugo/chess-puzzles.git
$GOPATH/bin/chess-puzzles
```

## Commands and key bindings

### Commands

commands are keywords you can type that are liked to an action: 

| commands | description                                                                                |
|----------|--------------------------------------------------------------------------------------------|
| !new     | get a new puzzle, if available                                                             |
| !hint    | highlight the next move of the puzzle, but marks it as failed.                             |
| e2e3     | move a piece from square `e2` to `e3` if the move is legal                                 |
| e7e8q    | move a piece from square `e7` to `e8` and promote `pawn` to `queen` if the move is allowed |

***To move you pieces on the board use [long algebraic notation](https://en.wikipedia.org/wiki/Algebraic_notation_(chess)#Long_algebraic_notation) in [Universal Chess Interface (UCI)](https://en.wikipedia.org/wiki/Universal_Chess_Interface) format.***

### key bindings

key bindings are keystrokes (shortcuts) linked to an action: 

| keys   | description          |
|--------|----------------------|
| Ctrl+C | exits the game       |
| Ctrl+U | clean the input area |
| Enter  | submit your command  |

## To-do

- [X] Load puzzles from chess.com API
- [X] Add score to show success or failure
- [X] Add command to get new puzzle 
- [X] Add command to reveal next move (this should mark the puzzle as failed) 
- [ ] Add a clock that resets when a new puzzle starts
- [ ] Add help command to display list of commands
- [ ] Load puzzles locally from JSON files
- [ ] Setup release of binaries in this repo
