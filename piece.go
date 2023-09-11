package quarto

type Piece struct {
	Tall  bool `json:"t"`
	Light bool `json:"l"`
	Round bool `json:"r"`
	Hole  bool `json:"h"`
}

func (p *Piece) is(other *Piece) bool {
	if other == nil {
		return false
	}

	return p.Tall == other.Tall &&
		p.Light == other.Light &&
		p.Round == other.Round &&
		p.Hole == other.Hole
}

func (p *Piece) clone() *Piece {
	return &Piece{
		Tall:  p.Tall,
		Light: p.Light,
		Round: p.Round,
		Hole:  p.Hole,
	}
}

var possiblePieces = []Piece{
	{true, true, true, true},
	{true, true, true, false},
	{true, true, false, true},
	{true, true, false, false},
	{true, false, true, true},
	{true, false, true, false},
	{true, false, false, true},
	{true, false, false, false},
	{false, true, true, true},
	{false, true, true, false},
	{false, true, false, true},
	{false, true, false, false},
	{false, false, true, true},
	{false, false, true, false},
	{false, false, false, true},
	{false, false, false, false},
}

func piecesHaveCommonFeature(p1, p2, p3, p4 *Piece) bool {
	if p1 == nil || p2 == nil || p3 == nil || p4 == nil {
		return false
	}

	return p1.Tall == p2.Tall && p2.Tall == p3.Tall && p3.Tall == p4.Tall ||
		p1.Light == p2.Light && p2.Light == p3.Light && p3.Light == p4.Light ||
		p1.Round == p2.Round && p2.Round == p3.Round && p3.Round == p4.Round ||
		p1.Hole == p2.Hole && p2.Hole == p3.Hole && p3.Hole == p4.Hole
}
