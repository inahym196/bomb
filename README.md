## How to Play

### install go

[Download and install](https://go.dev/doc/install)

### download inahym196/bomb

```bash
git clone https://github.com/inahym196/bomb
```

### play

```bash
cd bomb
go run cmd/main.go

Enter command> help
Available Commands:
  > start <mode: easy or normal>        Start Game, Select gameMode
  > start custom <width> <bombCount>    Start Game, Set Custom width and bombCount
  > show                                Show board
  > open <row: int> <col: alpha>        Open cell
  > check/uncheck <row> <col>           Check/UnCheck cell
  > help                                Show this help message
  > exit                                Exit the program

Enter command> start easy
gameState: Ready
   A B C D E F G H I
 0 □ □ □ □ □ □ □ □ □
 1 □ □ □ □ □ □ □ □ □
 2 □ □ □ □ □ □ □ □ □
 3 □ □ □ □ □ □ □ □ □
 4 □ □ □ □ □ □ □ □ □
 5 □ □ □ □ □ □ □ □ □
 6 □ □ □ □ □ □ □ □ □
 7 □ □ □ □ □ □ □ □ □
 8 □ □ □ □ □ □ □ □ □

Enter command> open 0 A
gameState: Playing
   A B C D E F G H I
 0 0 0 0 0 0 1 □ 1 0
 1 1 1 0 0 0 1 1 1 0
 2 □ 2 2 2 1 0 0 0 0
 3 □ □ □ □ 1 0 0 1 1
 4 □ □ □ □ 1 0 0 1 □
 5 □ □ □ □ 1 0 0 1 □
 6 □ □ □ □ 1 0 1 1 □
 7 □ □ □ □ 2 1 2 □ □
 8 □ □ □ □ □ □ □ □ □

Enter command> check 2 A
gameState: Playing
   A B C D E F G H I
 0 0 0 0 0 0 1 □ 1 0
 1 1 1 0 0 0 1 1 1 0
 2 x︎ 2 2 2 1 0 0 0 0
 3 □ □ □ □ 1 0 0 1 1
 4 □ □ □ □ 1 0 0 1 □
 5 □ □ □ □ 1 0 0 1 □
 6 □ □ □ □ 1 0 1 1 □
 7 □ □ □ □ 2 1 2 □ □
 8 □ □ □ □ □ □ □ □ □

Enter command> Enjoy!
```
