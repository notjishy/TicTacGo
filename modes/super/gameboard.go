package super

import (
	"fmt"
	"strings"
	"github.com/TwiN/go-color"
)

const (
	subRows = 3
	subCols = 3
	rows    = 3
	cols    = 3
)

var (
	player1 = color.Red
	player2 = color.Cyan
	standard = color.Blue
	accent = color.White
	active = color.Green
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
func PrintSuperBoard(availableMoves int) {
	// print the top row with column numbers for each super board
	fmt.Printf("\n     1   2   3 | 1   2   3 | 1   2   3\n")
	printSubHorDivider(0, availableMoves)

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
				color := standard
				if (i == ActiveSectorRow && j == ActiveSectorCol && availableMoves != 81) { color = active }
				printSubBoard(GameBoard.Cells[i][j].Cells[subRow], j, color)
			}
			fmt.Println() // move to the next line after printing a sub-row

			// print the horizontal dividers between sub-rows
			if subRow < subRows-1 {
				printSubHorDivider(i, availableMoves)
			}
		}

		// print the horizontal dividers between rows of the SuperBoard, except for the last row
		if i < rows-1 {
			fmt.Println(color.InBold("  ===+===+==== | ====+==== | ====+===+==="))
		}
	}
	printSubHorDivider(3, availableMoves)
	fmt.Println("") // move to the next line after printing the SuperBoard
}

func printSubHorDivider(num int, availableMoves int) {
	first := standard
	second := standard
	third := standard

	if (availableMoves != 81) {
		if (ActiveSectorRow == 0 && num < 1) {
			if (ActiveSectorCol == 0) {
				first = active
			} else if (ActiveSectorCol == 1) {
				second = active
			} else {
				third = active
			}
		} else if (ActiveSectorRow == 1 && num == 1) {
			if (ActiveSectorCol == 0) {
				first = active
			} else if (ActiveSectorCol == 1) {
				second = active
			} else {
				third = active
			}
		} else if (ActiveSectorRow == 2 && num > 1) {
			if (ActiveSectorCol == 0) {
				first = active
			} else if (ActiveSectorCol == 1) {
				second = active
			} else {
				third = active
			}
		}
	}
	fmt.Println(color.With(first, "   +---+---+---") + color.With(accent, "|") + color.With(second, "---+---+---") + color.With(accent, "|") + color.With(third, "---+---+---+"))
}

// prints a sub-board
func printSubBoard(subBoardRow [3]string, j int, sectorColor string) {

	for i, cell := range subBoardRow {
		cellPart := ""
		if i < 1 && j > 0 {
			cellPart = color.With(accent, " | %-3s")
		} else {
			cellPart = color.With(sectorColor, " | %-3s")
		}

		if strings.HasSuffix(cell, "X") {
			fmt.Printf(cellPart, color.With(player1, cell))
		} else {
			fmt.Printf(cellPart, color.With(player2, cell))
		}

		if j > 1 && i > 1 { fmt.Print(color.With(sectorColor, " |")) }
	}
}
