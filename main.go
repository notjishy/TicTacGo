package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"strings"
	"tictacgo/gamemodes"

	"github.com/inancgumus/screen"
)

func main() {
	// print main menu here and again at end of loop
	// that way we dont immediately clear out the error messages if there is one, because we dont close the program
	// from those errors theres no need for that
	mainMenu()

	for {
		// get input using askMode function
		selectedMode, err := askMode()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue // retry if there was an error
		}

		if selectedMode == "q" {
			break // quits the game
		}

		// get amount of players
		playerCount, err := askPlayerCount()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue // retry if there was an error
		}

		if selectedMode == "r" {
			err = gamemodes.PlayRegular(playerCount)
		} else if selectedMode == "s" {
			err = gamemodes.PlaySuper(playerCount)
		}

		mainMenu() // see above
		if err != nil {
			fmt.Println("Error during game:", err)
			continue
		}
	}
}

// main menu screen
func mainMenu() {
	// clear the screen
	screen.Clear()

	// print menu
	fmt.Print("\n\n _______ _   _______          _____\n",
		"|__   __(_) |__   __|        / ____|\n",
		"   | |   _  ___| | __ _  ___| |  __  ___\n",
		"   | |  | |/ __| |/ _` |/ __| | |_ |/ _ \\\n",
		"   | |  | | (__| | (_| | (__| |__| | (_) |\n",
		"   |_|  |_|\\___|_|\\__,_|\\___|\\_____|\\___/\n\n\n\n")

	fmt.Print("               Select a Mode:\n",
		(color.InGreen("  (R)egular  ")+color.With(color.Reset, " ||  "))+color.InCyan(
			" (S)uper  ")+color.With(color.Reset, " ||  ")+color.InRed(" (Q)uit\n\n"))
}

func askMode() (string, error) {
	var modeInput string

	// retrieve input from user
	_, err := fmt.Scan(&modeInput)
	if err != nil {
		return "", err
	}
	modeInput = strings.ToLower(modeInput)

	// check input with mode selection
	if modeInput != "r" && modeInput != "s" && modeInput != "q" {
		fmt.Println("Invalid option, please choose one of the listed modes.")
		return askMode() // recursively ask for input if no valid input
	}

	// send selected mode forward
	return modeInput, nil
}

// ask the user how many players (1 or 2)
func askPlayerCount() (int, error) {
	var inputNum int
	fmt.Print("\n(1) Player   ||    (2) Players\n")
	_, err := fmt.Scan(&inputNum)
	if err != nil {
		return 0, err
	}

	// verify valid number of players
	if inputNum != 1 && inputNum != 2 {
		fmt.Println("Invalid option, please choose either 1 or 2 players.")
		return askPlayerCount() // recursively ask for player count if no valid input
	}

	// send selected player count forward
	return inputNum, nil
}
