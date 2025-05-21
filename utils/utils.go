package utils

import (
	"fmt"
	"strings"
)

var confirmQuit string

func SwitchPlayer(player int) int {
	if player == 1 {
		return 2
	}
	return 1
}

func ConfirmQuit() (bool, error) {
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
