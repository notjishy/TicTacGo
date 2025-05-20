package utils

import (
	"fmt"
	"strings"

	"tictacgo/config"

	"github.com/TwiN/go-color"
	"github.com/inancgumus/screen"
)

// Board - define the board
var Board [3][3]string

// InitializeBoard - create the board
func InitializeBoard() {
	for row := range Board {
		for col := range Board[row] {
			Board[row][col] = " "
		}
	}
}

// PrintBoard - print the board in the terminal
func PrintBoard() {
	// clear the screen
	screen.Clear()

	fmt.Println("\n    1   2   3")
	// print out 1 row at a time
	for i, row := range Board {
		// + i to 'a' so the row letter increments up.
		// row 1 = a, row 2 = b, row 3 = c
		fmt.Print(string(rune('a' + i)))
		for _, cell := range row {
			if strings.HasSuffix(cell, config.Settings.Player1.Symbol) {
				playerColor := config.Settings.Player1.GetColor()
				fmt.Printf(" | %s", color.With(playerColor, cell))
			} else {
				playerColor := config.Settings.Player2.GetColor()
				fmt.Printf(" | %s", color.With(playerColor, cell))
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
