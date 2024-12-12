package utils

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
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
	fmt.Println("\n    1   2   3")
	// print out 1 row at a time
	for i, row := range Board {
		// + i to 'a' so the row letter increments up.
		// row 1 = a, row 2 = b, row 3 = c
		fmt.Print(string('a' + i))
		for _, cell := range row {
			if strings.HasSuffix(cell, "X") {
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
