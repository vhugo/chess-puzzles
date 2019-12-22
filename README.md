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
go get github.com/vhugo/chess-puzzles.git
cd $GOPATH/src/github.com/vhugo/chess-puzzles
go run .
```

## To-do

- [ ] Load puzzles from chess.com API
- [ ] Load puzzles locally from JSON files
- [ ] Setup release of binaries in this repo
