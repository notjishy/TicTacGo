package utils

import (
	"fmt"
	"strings"

	"tictacgo/config"
)

var confirmQuit string

func SwitchPlayer(player int) int {
	if player == 1 {
		return 2
	}
	return 1
}

func GetPlayerSymbol(player int) string {
	if player == 1 {
		return config.Settings.Player1.Symbol
	}
	return config.Settings.Player2.Symbol
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

func confirmQuitGame() (bool, error) {
	fmt.Print("Are you sure you want to quit? (Y/N) ")
	_, err := fmt.Scan(&confirmQuit)
	if err != nil {
		return false, err
	}

	if strings.ToLower(confirmQuit) != "y" {
		return false, nil
	}
	return true, nil
}
