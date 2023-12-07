package super

import (
	"fmt"
	"strings"
	"github.com/TwiN/go-color"
)

const (
	subRows = 3
	subCols = 3
	rows = 3
	cols = 3
)

var GameBoard SuperBoard

type SuperBoard struct {
	Cells [rows][cols]SubBoard
}

type SubBoard struct {
	Cells [subRows][subCols]string
}

// initializes the SuperBoard with empty spaces
func InitializeSuperBoard() {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for x := 0; x < subRows; x++ {
				for y := 0; y < subCols; y++ {
					GameBoard.Cells[i][j].Cells[x][y] = " "
				}
			}
		}
	}
}

// prints the SuperBoard with sub-boards
func PrintSuperBoard() {
    // print the top row with column numbers for each super board
    fmt.Println("\n   1   2   3   |   1   2   3   |   1   2   3")
    fmt.Println("  ---+---+---  |  ---+---+---  |  ---+---+---")

    // iterate over each row of the SuperBoard
    for i := 0; i < rows; i++ {
        // iterate over each sub-row within a SuperBoard row
        for subRow := 0; subRow < subRows; subRow++ {
            // calculate the current row label based on the row index and sub-row index
            rowLabel := string('a' + byte(subRow))
            
            // print the row label at the beginning of each sub-row
            fmt.Print(rowLabel, " ")

            // iterate over each column of the SuperBoard
            for j := 0; j < cols; j++ {
                // print the sub-board for the current row and column
                printSubBoard(GameBoard.Cells[i][j].Cells[subRow])
                fmt.Print(" |")
            }
            fmt.Println() // move to the next line after printing a sub-row

            // print the horizontal dividers between sub-rows
            if subRow < subRows-1 {
                fmt.Println("  ---+---+---  |  ---+---+---  |  ---+---+---")
            }
        }

        // print the horizontal dividers between rows of the SuperBoard, except for the last row
        if i < rows-1 {
            fmt.Println("  ===+===+===  |  ===+===+===  |  ===+===+===")
        }
    }
    fmt.Println() // move to the next line after printing the SuperBoard
}

// prints a sub-board
func printSubBoard(subBoardRow [3]string) {
    for _, cell := range subBoardRow {
        if strings.HasSuffix(cell, "X") {
            fmt.Printf(" | %s", color.InRed(cell))
        } else {
            fmt.Printf(" | %s", color.InBlue(cell))
        }
    }
}
