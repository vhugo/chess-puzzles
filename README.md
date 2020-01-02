# Chess Puzzles

A terminal-based chess puzzles interface. The inspiration behind it comes from chess puzzles at
[chess.com](https://www.chess.com/puzzles). This project was made possible because of two awesome 
projects, [termloop](https://github.com/JoelOtter/termloop) as the game engine and
[chess](https://github.com/notnil/chess) as the chess mechanics/engine.

![alt puzzle solved][chess_puzzle]

## Getting started

A binary release is still on the making so until then you will need [Golang](https://golang.org/doc/install) 
to compile and run the source yourself. You can either install [Golang](https://golang.org/doc/install) 
or use [Docker](https://www.docker.com/) with the following command `docker run -ti golang /bin/bash` 
to have a proper environment to run this project.

These are the commands to get the source, compile and run:

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


## License

Licensed under the MIT License. 

[chess_puzzle]: https://i.imgur.com/ByQRCmb.gif
