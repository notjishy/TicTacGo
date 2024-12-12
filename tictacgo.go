package main

import (
	"fmt"
	"strings"
	"tictacgo/regular"
	"tictacgo/super"

	"github.com/TwiN/go-color"
)

func main() {
	for {
		selectedMode := askMode() // get input using askMode function
		if selectedMode == "q" {
			break
		} // quits the game

		// get amount of players
		playerCount := askPlayerCount()

		// run chosen game mode
		if selectedMode == "r" {
			regular.Play(playerCount)
		} else if selectedMode == "s" {
			super.Play(playerCount)
		}
	}
}

// main menu screen
func askMode() string {
	var modeInput string
	// print menu
	fmt.Print("\n\n _______ _   _______          _____\n",
		"|__   __(_) |__   __|        / ____|\n",
		"   | |   _  ___| | __ _  ___| |  __  ___\n",
		"   | |  | |/ __| |/ _` |/ __| | |_ |/ _ \\\n",
		"   | |  | | (__| | (_| | (__| |__| | (_) |\n",
		"   |_|  |_|\\___|_|\\__,_|\\___|\\_____|\\___/\n\n\n\n")
	fmt.Print("               Select a Mode:\n",
		(color.InGreen("  (R)egular  ")+color.With(color.Reset, " ||  "))+color.InCyan(" (S)uper  ")+color.With(color.Reset, " ||  ")+color.InRed(" (Q)uit\n\n"))

	// retreive input from user
	fmt.Scan(&modeInput)
	modeInput = strings.ToLower(modeInput)

	// check input with mode selection
	if modeInput != "r" && modeInput != "s" && modeInput != "q" {
		fmt.Println("Invalid option, please choose one of the listed modes.")
		return askMode() // recursively ask for input if no valid input
	}

	// send selected mode forward
	return modeInput
}

// ask the user how many players (1 or 2)
func askPlayerCount() int {
	var inputNum int
	fmt.Print("\n(1) Player   ||    (2) Players\n")
	fmt.Scan(&inputNum)

	// verify valid number of players
	if inputNum != 1 && inputNum != 2 {
		fmt.Println("Invalid option, please choose either 1 or 2 players.")
		return askPlayerCount() // recursively ask for player count if no valid input
	}

	// send selected player count forward
	return inputNum
}
