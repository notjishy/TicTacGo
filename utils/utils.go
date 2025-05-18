package utils

import (
	"fmt"
	"strings"

	"tictacgo/config"
)

var confirmQuit string

func SwitchPlayer(playerCount int, player int) int {
	if player == 1 {
		return 2
	}
	return 1
}

func GetPlayerSymbol(player int) string {
	gameConfig := config.GetConfig()
	if player == 1 {
		return gameConfig.Player1
	}
	return gameConfig.Player2
}

func ParsePlayerMove(move string) (int, int, bool) {
	move = strings.ToLower(move)

	if len(move) != 2 {
		return 0, 0, false
	}
	row := int(move[0] - 'a')
	col := int(move[1] - '1')
	return row, col, row >= 0 && row < 3 && col >= 0 && col < 3
}

func confirmQuitGame() bool {
	fmt.Print("Are you sure you want to quit? (Y/N) ")
	fmt.Scan(&confirmQuit)

	if strings.ToLower(confirmQuit) != "y" {
		return false
	}
	return true
}
