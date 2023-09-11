package quarto

type move struct {
	Piece Piece
	X     uint8
	Y     uint8
}

func (m *move) AsByte() byte {
	var b byte

	b |= m.X << 6
	b |= m.Y << 4

	if m.Piece.Tall {
		b |= 1 << 3
	}
	if m.Piece.Light {
		b |= 1 << 2
	}
	if m.Piece.Round {
		b |= 1 << 1
	}
	if m.Piece.Hole {
		b |= 1
	}

	return b
}
