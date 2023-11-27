package main

import (
	"fmt"
	"math/rand"
)

type Board []int

const Player = 1
const Computer = 2
const PlayerToken = 1
const ComputerToken = 10
const BoardWidth = 3
const BoardHeight = 3

// getUserMove
// calcComputerMove
// checkWinning

func main() {
	board := Board{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for isGameOver(board) == false {
		printBoard(board)
		c := getUserMove()
		board = placeTokenOnBoard(board, c, Player)
		printBoard(board)
		if isGameOver(board) {
			fmt.Println("Player won")
			return
		}
		c = getComputerMove(board)
		board = placeTokenOnBoard(board, c, Computer)
		printBoard(board)
		if isGameOver(board) {
			fmt.Println("Computer won")
			return
		}
	}
}

// getUserMove gets the move of the user
func getUserMove() int {
	var i int
	fmt.Scanln(&i)
	return i

}

// getComputerMove returns the move done by the computer
func getComputerMove(board Board) int {
	spots := getAvailableSpotsOnBoard(board)
	return spots[rand.Intn(len(spots))]
}

// printBoard prints a tic tac toe board
// 1 | 2 | 3
// 4 | 5 | 6
// 7 | 8 | 9
func printBoard(board Board) {
	for i := 0; i < 9; i += 3 {
		fmt.Println(getBoardChar(board[i]), " | ", getBoardChar(board[i+1]), " | ", getBoardChar(board[i+2]))
	}
	fmt.Println()
}

// getBoardChar prints one field of the tic tac toe board
// at the current print position
func getBoardChar(field int) string {
	switch field {
	case 0:
		return " "
	case PlayerToken:
		return "X"
	case ComputerToken:
		return "O"
	}
	return ""
}

// placeTokenOnBoard gets a coordinate and a board
// Returns the board with the symbol placed on it
// Does NOT check if coordinate on board is available
func placeTokenOnBoard(board Board, cod int, player int) Board {
	switch player {
	case Player:
		board[cod] = PlayerToken
	case Computer:
		board[cod] = ComputerToken
	}
	return board
}

// isGameOver returns true if no more moves are possible
// condition 1: one player has won
// condition 2: the board is filled
func isGameOver(board Board) bool {
	for i := 0; i < len(board); i += 3 {
		sum := board[i] + board[i+1] + board[i+2]
		if sum == 3 || sum == 30 {
			return true
		}
	}
	for i := 0; i < BoardWidth; i++ {
		sum := board[i] + board[i+3] + board[i+6]
		if sum == 3 || sum == 30 {
			return true
		}
	}
	sum := board[0] + board[4] + board[8]
	if sum == 3 || sum == 30 {
		return true
	}
	sum = board[2] + board[4] + board[6]
	if sum == 3 || sum == 30 {
		return true
	}
	return false
}

// getAvailableSpotsOnBoard returns all available spots on the board
func getAvailableSpotsOnBoard(board Board) []int {
	var spots []int
	for i := range board {
		if board[i] == 0 {
			spots = append(spots, i)
		}
	}
	return spots
}
