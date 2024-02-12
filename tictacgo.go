package main

import (
	"fmt"
	"strings"
	"tictacgo/regular"
	"tictacgo/super"
	"github.com/TwiN/go-color"
)

func askMode() string {
    var inputChar string
	fmt.Print("\n\n _______ _   _______          _____\n",
	"|__   __(_) |__   __|        / ____|\n",
	"   | |   _  ___| | __ _  ___| |  __  ___\n",
	"   | |  | |/ __| |/ _` |/ __| | |_ |/ _ \\\n",
	"   | |  | | (__| | (_| | (__| |__| | (_) |\n",
	"   |_|  |_|\\___|_|\\__,_|\\___|\\_____|\\___/\n\n\n\n")
    fmt.Print("               Select a Mode:\n",
			(color.InGreen("  (R)egular  ") + color.With(color.Reset, " ||  ")) + color.InCyan(" (S)uper  ") + color.With(color.Reset, " ||  ") + color.InRed(" (Q)uit\n\n"))
    fmt.Scan(&inputChar)
    inputChar = strings.ToLower(inputChar)

    if inputChar != "r" && inputChar != "s" && inputChar != "q" {
        fmt.Println("Invalid option, please choose one of the listed modes.")
        return askMode() // recursively ask for input
    }

    return inputChar
}

func askPlayerCount() int {
	var inputNum int
	fmt.Print("\n(1) Player   ||    (2) Players\n")
	fmt.Scan(&inputNum)

	if inputNum != 1 && inputNum != 2 {
		fmt.Println("Invalid option, please choose either 1 or 2 players.")
		return askPlayerCount()
	}

	return inputNum
}

func main() {
	for {
		selectedMode := askMode() // get input using askMode function
		if selectedMode == "q" { break } // quit game

		playerCount := askPlayerCount() // get amount of players using askPlayerCount function

		// run chosen game mode
		if selectedMode == "r" {
			regular.Play(playerCount)
		} else if (selectedMode == "s") {
			super.Play(playerCount)
		}
	}
}
