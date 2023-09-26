package tictactoe

import (
	"fmt"
)

type TicTacToe struct {
	board  [3][3]Piece
	winner Piece
	turn   Piece
	moves  []Move
}

func (ttt *TicTacToe) NewGame() {
	ttt.board = [3][3]Piece{}
	ttt.winner = Empty
	ttt.turn = X
	ttt.moves = make([]Move, 0)
}

func (ttt *TicTacToe) GameOver() Piece {
	// Shortcut if the winner has already been decided
	if ttt.winner != Empty {
		return ttt.winner
	}

	// Horizontal three in a row
	if ttt.board[0][0] == ttt.board[0][1] && ttt.board[0][1] == ttt.board[0][2] {
		ttt.winner = ttt.board[0][0]
	}
	if ttt.winner == Empty && ttt.board[1][0] == ttt.board[1][1] && ttt.board[1][1] == ttt.board[1][2] {
		ttt.winner = ttt.board[1][0]
	}
	if ttt.winner == Empty && ttt.board[2][0] == ttt.board[2][1] && ttt.board[2][1] == ttt.board[2][2] {
		ttt.winner = ttt.board[2][0]
	}
	// Vertical three in a row
	if ttt.winner == Empty && ttt.board[0][0] == ttt.board[1][0] && ttt.board[1][0] == ttt.board[2][0] {
		ttt.winner = ttt.board[0][0]
	}
	if ttt.winner == Empty && ttt.board[0][1] == ttt.board[1][1] && ttt.board[1][1] == ttt.board[2][1] {
		ttt.winner = ttt.board[0][1]
	}
	if ttt.winner == Empty && ttt.board[0][2] == ttt.board[1][2] && ttt.board[1][2] == ttt.board[2][2] {
		ttt.winner = ttt.board[0][2]
	}
	// Diagonal
	if ttt.winner == Empty && ttt.board[0][0] == ttt.board[1][1] && ttt.board[1][1] == ttt.board[2][2] {
		ttt.winner = ttt.board[0][0]
	}
	if ttt.winner == Empty && ttt.board[0][2] == ttt.board[1][1] && ttt.board[1][1] == ttt.board[2][0] {
		ttt.winner = ttt.board[0][2]
	}
	return ttt.winner
}

func (ttt *TicTacToe) Move(move Move) Move {
	// Check that the game is not over
	if ttt.winner != Empty {
		move.Err = true
		move.Message = fmt.Sprintf("No new moves can be made. %s already won", ttt.winner)
		return move
	}
	// Check that the move is in bounds
	if move.X < 0 || move.X > 2 || move.Y < 0 || move.Y > 2 {
		move.Err = true
		move.Message = fmt.Sprintf("Cell at (%d, %d) does not exist", move.X, move.Y)
		return move
	}
	// Moves are not allowed to replace a previously placed piece
	if ttt.board[move.X][move.Y] != Empty {
		move.Err = true
		move.Message = fmt.Sprintf("Cell at (%d, %d) not empty", move.X, move.Y)
		return move
	}

	// Place piece
	move.Piece = ttt.turn
	ttt.board[move.X][move.Y] = move.Piece

	// Record move
	ttt.moves = append(ttt.moves, move)

	// Check for winner
	winner := ttt.GameOver()
	if winner != Empty {
		move.Message = fmt.Sprintf("%s is the winner", winner)
		return move
	}

	// Change turn
	if ttt.turn == X {
		ttt.turn = O
	} else {
		ttt.turn = X
	}

	return move
}
