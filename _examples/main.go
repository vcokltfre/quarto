package main

import (
	"fmt"

	"github.com/vcokltfre/quarto"
)

func main() {
	game := quarto.New()

	// Player 0 nominates TDRH, Player 1 places at 0, 0
	game.Nominate(quarto.Piece{true, false, true, true})
	game.Place(0, 0)

	// Player 1 nominates SDRH, Player 0 places at 1, 0
	game.Nominate(quarto.Piece{false, false, true, true})
	game.Place(1, 0)

	// Player 0 nominates TDSH, Player 1 places at 2, 0
	game.Nominate(quarto.Piece{true, false, false, true})
	game.Place(2, 0)

	// Player 1 nominates TDRS, Player 0 places at 3, 0
	game.Nominate(quarto.Piece{true, false, true, false})
	game.Place(3, 0)

	// Game is won by player 0
	fmt.Println(game.IsWon())
}
