package tictactoe

type Move struct {
	Piece   Piece
	X       int
	Y       int
	Message string
	Err     bool
}
