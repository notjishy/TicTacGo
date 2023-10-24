package regular

import (
	"fmt"
	"math/rand"
	"tictacgo/logic"
)

var board [3][3]string
var boxes = []string{"a1", "a2", "a3", "b1", "b2", "b3", "c1", "c2", "c3"}

func initializeBoard() {
	for row := range board {
		for col := range board[row] {
			board[row][col] = " "
		}
	}
}

func printBoard() {
	fmt.Println("   1   2   3")
	for i, row := range board {
		fmt.Print(string('a'+i))
		for _, cell := range row {
			fmt.Printf(" | %s", cell)
		}
		fmt.Println(" |")
		if i < 2 {
			fmt.Println("  ---+---+---")
		}
	}
	fmt.Println()
}

func switchPlayer(playerCount int, player int) int {
	if player == 1 {
		return 2
	}
	return 1
}

func getPlayerMove(player int) {
	fmt.Printf("Player %d, enter your move (e.g., A1, B2): ", player)
	var move string
	fmt.Scanln(&move)
	row, col, valid := parseMove(move)
	if !valid || board[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		getPlayerMove(player)
		return
	}
	board[row][col] = logic.PlayerSymbol(player)
}

func parseMove(move string) (int, int, bool) {
	if len(move) != 2 {
		return 0, 0, false
	}
	row := int(move[0] - 'a')
	col := int(move[1] - '1')
	return row, col, row >= 0 && row < 3 && col >= 0 && col < 3
}

func Play(playerCount int) {
	initializeBoard()
	player := 1
	availableMoves := 9
	for availableMoves > 0 {
		if (player == 1 || (playerCount == 2 && player == 2)) {
			printBoard()
			getPlayerMove(player)
		} else {
			getComputerMove()
		}
		availableMoves--
		if availableMoves < 5 {
			if logic.CheckWin(player, board) {
				printBoard()
				fmt.Printf("Player %d wins!\n", player)
				return
			}
		}
		player = switchPlayer(playerCount, player)
	}
	fmt.Println("It's a tie!")
}

func getComputerMove() {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if board[row][col] == " " {
			board[row][col] = "O"
			return
		}
	}
}
