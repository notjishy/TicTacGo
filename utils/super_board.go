package utils

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
	player1  = color.Red
	player2  = color.Cyan
	standard = color.Blue
	accent   = color.Reset
	active   = color.Green
	tie      = color.Yellow
)

var ActiveSectorRow int
var ActiveSectorCol int
var SectorBlocked bool

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

	InitializeBoard()
}

func GetEmptySpaces() int {
	openSubSpaces := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[i][j] == " " {
				openSubSpaces++
			}
		}
	}
	return openSubSpaces
}

func GetEmptySubBoards() int {
	openSubBoards := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if Board[i][j] == " " {
				openSubBoards++
			}
		}
	}
	return openSubBoards
}

// prints the SuperBoard with sub-boards
func PrintSuperBoard(availableMoves int, sectorBlocked bool, gameEnd bool) {
	var regColsTaken []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if Board[i][j] != " " {
				regColsTaken = append(regColsTaken, j)
			}
		}
	}

	// print the top row with column numbers for each super board
	fmt.Printf("\n     1   2   3 | 1   2   3 | 1   2   3\n")
	printSubHorDivider(0, regColsTaken, availableMoves, sectorBlocked, gameEnd)

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
				if i == ActiveSectorRow && j == ActiveSectorCol && availableMoves != 81 && !sectorBlocked && !gameEnd {
					color = active
				} else if Board[i][j] == "X" {
					color = player1
				} else if Board[i][j] == "O" {
					color = player2
				} else if Board[i][j] == "-" {
					color = tie
				}
				printSubBoard(GameBoard.Cells[i][j].Cells[subRow], i, j, color)
			}
			fmt.Println() // move to the next line after printing a sub-row

			// print the horizontal dividers between sub-rows
			if subRow < subRows-1 {
				printSubHorDivider(i, regColsTaken, availableMoves, sectorBlocked, gameEnd)
			}
		}

		// print the horizontal dividers between rows of the SuperBoard, except for the last row
		if i < rows-1 {
			fmt.Println(color.InBold("  ===+===+==== | ====+==== | ====+===+==="))
		}
	}
	printSubHorDivider(2, regColsTaken, availableMoves, sectorBlocked, gameEnd)
	fmt.Println("") // move to the next line after printing the SuperBoard
}

func printSubHorDivider(i int, regColsTaken []int, availableMoves int, sectorBlocked bool, gameEnd bool) {
	first := standard
	second := standard
	third := standard

	if availableMoves < 81 && !sectorBlocked && !gameEnd {
		if (ActiveSectorRow == 0 && i < 1) || (ActiveSectorRow == 1 && i == 1) || (ActiveSectorRow == 2 && i > 1) {
			if ActiveSectorCol == 0 {
				first = active
			} else if ActiveSectorCol == 1 {
				second = active
			} else {
				third = active
			}
		}
	}
	if len(regColsTaken) > 0 {
		for j := 0; j < len(regColsTaken); j++ {
			if Board[i][regColsTaken[j]] != " " {
				var blockedSectorColor string
				if Board[i][regColsTaken[j]] == "X" {
					blockedSectorColor = player1
				} else if Board[i][regColsTaken[j]] == "O" {
					blockedSectorColor = player2
				} else if Board[i][regColsTaken[j]] == "-" {
					blockedSectorColor = tie
				}

				switch regColsTaken[j] {
				case 0:
					first = blockedSectorColor
				case 1:
					second = blockedSectorColor
				case 2:
					third = blockedSectorColor
				}
			}
		}
	}
	fmt.Println(color.With(first, "   +---+---+---") + color.With(accent, "|") + color.With(second, "---+---+---") + color.With(accent, "|") + color.With(third, "---+---+---+"))
}

// prints a sub-board
func printSubBoard(subBoardRow [3]string, i int, j int, sectorColor string) {
	for x, cell := range subBoardRow {
		cellPart := ""
		if x < 1 && j > 0 {
			cellPart = color.With(accent, " | %-3s")
		} else {
			cellPart = color.With(sectorColor, " | %-3s")
		}

		if Board[i][j] == "X" {
			fmt.Printf(cellPart, color.With(player1, cell))
			if j > 1 && x > 1 {
				fmt.Print(color.With(player1, " |"))
			}
		} else if Board[i][j] == "O" {
			fmt.Printf(cellPart, color.With(player2, cell))
			if j > 1 && x > 1 {
				fmt.Print(color.With(player2, " |"))
			}
		} else if Board[i][j] == "-" {
			fmt.Printf(cellPart, color.With(tie, cell))
			if j > 1 && x > 1 {
				fmt.Print(color.With(tie, " |"))
			}
		} else {
			if strings.HasSuffix(cell, "X") {
				fmt.Printf(cellPart, color.With(player1, cell))
			} else {
				fmt.Printf(cellPart, color.With(player2, cell))
			}
			if j > 1 && x > 1 {
				fmt.Print(color.With(sectorColor, " |"))
			}
		}
	}
}
