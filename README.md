# Quarto

A backend implementation of Quarto's core mechanics as a library.

## Public API

```go
type quarto.Piece struct {
    Tall  bool
    Light bool
    Round bool
    Hole  bool
}

type quarto.Game struct {
    func Nominate(piece Piece)  error
    func Place(x, y uint8)      error           // Bounded from (0, 0) to (3, 3)
    func GetAvailablePieces()   []Piece
    func IsWon()                (bool, int)     // isWon, winningPlayer
    func IsOver()               bool
    func AsBytes()              []byte          // in format defined below
    func ExportBoard()          ([]byte, error) // json representation
}

func New() *quarto.Game
```

## Export format

Games are exported as a series of bytes each describing a move. The format of each move byte as is follows (0 leftmost, 7 rightmost):

| Bits | Value   |
|------|---------|
| 0-1  | X       |
| 2-3  | Y       |
| 4    | IsTall  |
| 5    | IsLight |
| 6    | IsRound |
| 7    | IsHole  |

For example, a move where TLSS (tall, light, square, solid) was placed at 2, 0 would be represented as `0b10001100` (`0x8c`)

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
