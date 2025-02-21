package utils

import (
	"fmt"
	"strings"

	"tictacgo/config"

	"github.com/TwiN/go-color"
	"github.com/inancgumus/screen"
)

// define the board
var Board [3][3]string

// create the board
func InitializeBoard() {
	for row := range Board {
		for col := range Board[row] {
			Board[row][col] = " "
		}
	}
}

// print the board in the terminal
func PrintBoard() {
	config := config.GetConfig()

	// clear the screen
	screen.Clear()

	fmt.Println("\n    1   2   3")
	// print out 1 row at a time
	for i, row := range Board {
		// + i to 'a' so the row letter increments up.
		// row 1 = a, row 2 = b, row 3 = c
		fmt.Print(string(rune('a' + i)))
		for _, cell := range row {
			if strings.HasSuffix(cell, config.Player1) {
				fmt.Printf(" | %s", color.InRed(cell))
			} else {
				fmt.Printf(" | %s", color.InCyan(cell))
			}
		}
		fmt.Println(" |")
		// print horizontal border
		if i < 2 {
			fmt.Println("   ---+---+---")
		}
	}
	fmt.Println()
}
