package utils

import (
	"strings"
)

func SwitchPlayer(playerCount int, player int) int {
	if player == 1 {
		return 2
	}
	return 1
}

func GetPlayerSymbol(player int) string {
	if player == 1 {
		return "X"
	}
	return "O"
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
