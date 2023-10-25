package regular

import (
	"fmt"
	"strings"
	"math/rand"
	"tictacgo/utils"
)

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
	if !valid || utils.Board[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		getPlayerMove(player)
		return
	}
	utils.Board[row][col] = utils.PlayerSymbol(player)
}

func parseMove(move string) (int, int, bool) {
	move = strings.ToLower(move)

	if len(move) != 2 {
		return 0, 0, false
	}
	row := int(move[0] - 'a')
	col := int(move[1] - '1')
	return row, col, row >= 0 && row < 3 && col >= 0 && col < 3
}

func Play(playerCount int) {
	utils.InitializeBoard()
	player := 1
	availableMoves := 9
	for availableMoves > 0 {
		if (player == 1 || (playerCount == 2 && player == 2)) {
			utils.PrintBoard()
			getPlayerMove(player)
		} else {
			getComputerMove()
		}
		availableMoves--
		if availableMoves < 5 {
			if utils.CheckWin(player) {
				utils.PrintBoard()
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
		if utils.Board[row][col] == " " {
			utils.Board[row][col] = "O"
			return
		}
	}
}
