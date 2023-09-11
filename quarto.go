package quarto

import (
	"errors"
)

type QuartoGame struct {
	boards    []*board
	nominated *Piece
	started   bool
	moves     []move
}

func (g *QuartoGame) getCurrentBoard() *board {
	if !g.started {
		return nil
	}

	return g.boards[len(g.boards)-1]
}

func (g *QuartoGame) createNextBoard() *board {
	if !g.started {
		return &board{}
	}

	return g.getCurrentBoard().Clone()
}

func (g *QuartoGame) isPieceInUse(piece *Piece) bool {
	if !g.started {
		return false
	}

	for _, row := range g.getCurrentBoard() {
		for _, p := range row {
			if piece.Is(p) {
				return true
			}
		}
	}

	return false
}

func (g *QuartoGame) isSlotInUse(x, y uint8) bool {
	if !g.started {
		return false
	}

	return g.getCurrentBoard()[x][y] != nil
}

func (g *QuartoGame) Nominate(piece Piece) error {
	if g.isPieceInUse(&piece) {
		return errors.New("piece is already in use on the board")
	}
	if len(g.boards) == 16 {
		return errors.New("game is already over")
	}

	g.nominated = &piece

	return nil
}

func (g *QuartoGame) Place(x, y uint8) error {
	if x > 3 || y > 3 {
		return errors.New("slot is out of bounds")
	}

	if g.isSlotInUse(x, y) {
		return errors.New("slot is already in use on the board")
	}

	if g.nominated == nil {
		return errors.New("no piece has been nominated")
	}
	if len(g.boards) == 16 {
		return errors.New("game is already over")
	}

	g.boards = append(g.boards, g.createNextBoard())
	g.started = true

	currentBoard := g.getCurrentBoard()
	currentBoard[x][y] = g.nominated

	g.nominated = nil

	g.moves = append(g.moves, move{
		Piece: *currentBoard[x][y],
		X:     x,
		Y:     y,
	})

	return nil
}

func (g *QuartoGame) GetAvailablePieces() []Piece {
	available := []Piece{}
	for _, piece := range possiblePieces {
		if !g.isPieceInUse(&piece) {
			available = append(available, piece)
		}
	}

	return available
}

func (g *QuartoGame) IsWon() (bool, int) {
	currentBoard := g.getCurrentBoard()
	if currentBoard == nil {
		return false, 0
	}

	return currentBoard.IsWon(), len(g.boards) % 2
}

func (g *QuartoGame) AsBytes() []byte {
	output := make([]byte, len(g.moves))

	for i, m := range g.moves {
		output[i] = m.AsByte()
	}

	return output
}

func New() *QuartoGame {
	return &QuartoGame{
		boards:    []*board{},
		nominated: nil,
	}
}
