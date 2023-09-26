package tictactoe

const (
	Empty Piece = iota
	X
	O
)

type Piece int

func (p Piece) String() string {
	return []string{
		"Empty", "X", "O",
	}[p]
}
