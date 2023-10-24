package main

import (
	"fmt"
	"strings"
	"tictacgo/modes/regular"
	"tictacgo/modes/super"
)

func askMode() string {
    var inputChar string
    fmt.Print("Which mode would you like to play? (R/Regular | S/Super) ")
    fmt.Scan(&inputChar)
    inputChar = strings.ToLower(inputChar)

    if inputChar != "r" && inputChar != "s" {
        fmt.Println("Invalid option, please choose one of the listed modes.")
        return askMode() // recursively ask for input
    }

    return inputChar
}

func askPlayerCount() int {
	var inputNum int
	fmt.Print("How many players will there be? (1/2) ")
	fmt.Scan(&inputNum)

	if inputNum != 1 && inputNum != 2 {
		fmt.Println("Invalid option, please choose either 1 or 2 players.")
		return askPlayerCount()
	}

	return inputNum
}

func main() {
    selectedMode := askMode() // get input using askMode function
	playerCount := askPlayerCount() // get amount of players using askPlayerCount function
	
	if selectedMode == "r" {
		regular.Play(playerCount)
	} else if (selectedMode == "s") {
		super.Play()
	}
}
