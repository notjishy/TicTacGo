package utils

import "fmt"

var Board [3][3]string

func InitializeBoard() {
	for row := range Board {
		for col := range Board[row] {
			Board[row][col] = " "
		}
	}
}

func PrintBoard() {
	fmt.Println("\n   1   2   3")
	for i, row := range Board {
		fmt.Print(string('a'+i))
		for _, cell := range row {
			fmt.Printf(" | %s", cell)
		}
		fmt.Println(" |")
		if i < 2 {
			fmt.Println("  ---+---+---")
		}
	}
	fmt.Println()
}