package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"math/rand"
	"os"
)

type Board []int

const Player = 1
const Computer = 2
const Draw = 3
const PlayerToken = 50
const ComputerToken = 10
const BoardWidth = 3
const BoardHeight = 3

// getUserMove
// calcComputerMove
// checkWinning

func main() {
	board := Board{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for {
		printGameResultAndExitIfFinished(board)
		printBoard(board)
		c := getUserMove(board)
		board = placeTokenOnBoard(board, c, Player)
		printBoard(board)
		printGameResultAndExitIfFinished(board)
		c = getComputerMove(board)
		board = placeTokenOnBoard(board, c, Computer)
		printBoard(board)
		printGameResultAndExitIfFinished(board)
	}
}

func printGameResultAndExitIfFinished(board Board) {
	gameOver, winner := isGameOver(board)
	if gameOver {
		switch winner {
		case Player:
			fmt.Println("Player won")
		case Computer:
			fmt.Println("Computer won")
		default:
			fmt.Println("It is a draw")
		}
		os.Exit(0)
	}
	return
}

// getUserMove gets the move of the user
func getUserMove(board Board) int {
	var i int
	possibleMoves := getAvailableSpotsOnBoard(board)
	tm.Print("Your move: ", possibleMoves, " ")
	tm.Flush()
	fmt.Scanln(&i)
	return i

}

// getComputerMove returns the move done by the computer
func getComputerMove(board Board) int {
	// check if computer could win
	// check if player could win
	// try all available spots, for each check
	//   could player win after next move => avoid
	//   could computer win after next move => play
	// choose random move
	trueOrFalseCom, cCom := boardsForComputer(board)
	if trueOrFalseCom == true {
		return cCom
	}
	trueOrFlalsePl, cPl := noWinForPlayerBoards(board)
	if trueOrFlalsePl == true {
		return cPl
	} else {
		return cPl
	}

}

// printBoard prints a tic tac toe board
// 1 | 2 | 3
// 4 | 5 | 6
// 7 | 8 | 9
func printBoard(board Board) {
	tm.Clear()
	tm.MoveCursor(1, 1)
	for i := 0; i < 9; i += 3 {
		tm.Print(getBoardChar(board[i]), " | ", getBoardChar(board[i+1]), " | ", getBoardChar(board[i+2]))
		tm.Printf("     %d %d %d\n", i, i+1, i+2)
	}
	tm.Println()
	tm.Flush()
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
// returns the winning player or draw
func isGameOver(board Board) (bool, int) {
	for i := 0; i < len(board); i += 3 {
		sum := board[i] + board[i+1] + board[i+2]
		if sum == 150 {
			return true, Player
		}
		if sum == 30 {
			return true, Computer
		}
	}
	for i := 0; i < BoardWidth; i++ {
		sum := board[i] + board[i+3] + board[i+6]
		if sum == 150 {
			return true, Player
		}
		if sum == 30 {
			return true, Computer
		}
	}
	sum := board[0] + board[4] + board[8]
	if sum == 150 {
		return true, Player
	}
	if sum == 30 {
		return true, Computer
	}
	sum = board[2] + board[4] + board[6]
	if sum == 150 {
		return true, Player
	}
	if sum == 30 {
		return true, Computer
	}
	if len(getAvailableSpotsOnBoard(board)) == 0 {
		return true, Draw
	}
	return false, 0
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

// boardsForComputer generates all possible boards from the next move to determine the best one
func boardsForComputer(board Board) (bool, int) {
	spots := getAvailableSpotsOnBoard(board)
	var testBoard = make([]int, len(board))
	for _, i := range spots {
		copy(testBoard, board)
		testBoard = placeTokenOnBoard(testBoard, i, Computer)
		gameOver, winner := isGameOver(testBoard)
		if gameOver == true && winner == Computer {
			return true, i
		}
	}
	return false, spots[rand.Intn(len(spots))]
}

// if boardsForComputer doesn't return true noWinForPlayerBoards will determine a move to block the win of the player
func noWinForPlayerBoards(board Board) (bool, int) {
	spots := getAvailableSpotsOnBoard(board)
	var testBoard = make([]int, len(board))
	for _, i := range spots {
		copy(testBoard, board)
		testBoard = placeTokenOnBoard(testBoard, i, Player)
		gameOver, winner := isGameOver(testBoard)
		if gameOver == true && winner == Player {
			return true, i
		}
	}
	return false, spots[rand.Intn(len(spots))]
}
