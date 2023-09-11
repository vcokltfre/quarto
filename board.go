package quarto

import "encoding/json"

type (
	pieceLocation [2]uint8
	winLocation   [4]pieceLocation
	board         [4][4]*Piece
)

var winLocations = []winLocation{
	// Horizontal
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, // Top
	{{0, 1}, {1, 1}, {2, 1}, {3, 1}}, // Middle Top
	{{0, 2}, {1, 2}, {2, 2}, {3, 2}}, // Middle Bottom
	{{0, 3}, {1, 3}, {2, 3}, {3, 3}}, // Bottom

	// Vertical
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, // Left
	{{1, 0}, {1, 1}, {1, 2}, {1, 3}}, // Middle Left
	{{2, 0}, {2, 1}, {2, 2}, {2, 3}}, // Middle Right
	{{3, 0}, {3, 1}, {3, 2}, {3, 3}}, // Right

	// Diagonal
	{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, // Start Top Left
	{{3, 0}, {2, 1}, {1, 2}, {0, 3}}, // Start Top Right

	// Square
	{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, // Top Left
	{{1, 0}, {2, 0}, {1, 1}, {2, 1}}, // Top Middle
	{{2, 0}, {3, 0}, {2, 1}, {3, 1}}, // Top Right
	{{0, 1}, {1, 1}, {0, 2}, {1, 2}}, // Middle Left
	{{1, 1}, {2, 1}, {1, 2}, {2, 2}}, // Middle Middle
	{{2, 1}, {3, 1}, {2, 2}, {3, 2}}, // Middle Right
	{{0, 2}, {1, 2}, {0, 3}, {1, 3}}, // Bottom Left
	{{1, 2}, {2, 2}, {1, 3}, {2, 3}}, // Bottom Middle
	{{2, 2}, {3, 2}, {2, 3}, {3, 3}}, // Bottom Right
}

func (b *board) AsJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *board) clone() *board {
	clone := &board{}

	for x, row := range b {
		for y, piece := range row {
			if piece == nil {
				continue
			}

			clone[x][y] = piece.clone()
		}
	}

	return clone
}

func (b *board) IsWon() bool {
	for _, location := range winLocations {
		p1 := b[location[0][0]][location[0][1]]
		p2 := b[location[1][0]][location[1][1]]
		p3 := b[location[2][0]][location[2][1]]
		p4 := b[location[3][0]][location[3][1]]

		if piecesHaveCommonFeature(p1, p2, p3, p4) {
			return true
		}
	}

	return false
}
