// WARNING
// THIS FILE IS FUCKING AWFUL AND WILL BE REFACTORED

package super

import (
	"fmt"
	"math/rand"
	"strings"
	"tictacgo/config"
	"tictacgo/modes/regular"
	"tictacgo/utils"

	"github.com/TwiN/go-color"
	"github.com/inancgumus/screen"
)

const (
	subRows = 3
	subCols = 3
	rows    = 3
	cols    = 3
)

var (
	player1 = func() string {
		return config.Settings.Player1.GetColor()
	}
	player2 = func() string {
		return config.Settings.Player2.GetColor()
	}
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

// InitializeSuperBoard - initializes the SuperBoard with empty spaces
func InitializeSuperBoard() {
	// Iterate over each cell in a flattened index
	totalCells := rows * cols * subRows * subCols
	for idx := 0; idx < totalCells; idx++ {
		// Calculate the respective indices for i, j, x, y
		i := idx / (cols * subRows * subCols)
		j := (idx / (subRows * subCols)) % cols
		x := (idx / subCols) % subRows
		y := idx % subCols

		// Set the cell to empty space
		GameBoard.Cells[i][j].Cells[x][y] = " "
	}

	regular.InitializeBoard()
}

// GetEmptySpaces - count how many empty spaces remaining in a subboard
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

// GetEmptySubBoards - count how many subboards remain in play
func GetEmptySubBoards() int {
	openSubBoards := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if regular.Board[i][j] == " " {
				openSubBoards++
			}
		}
	}
	return openSubBoards
}

// PrintSuperBoard - prints the SuperBoard with sub-boards
func PrintSuperBoard(gameEnd bool) {
	// clear the screen
	screen.Clear()

	var regColsTaken []int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if regular.Board[i][j] != " " {
				regColsTaken = append(regColsTaken, j)
			}
		}
	}

	// print the top row with column numbers for each super board
	fmt.Printf("\n     1   2   3 | 1   2   3 | 1   2   3\n")
	printSubHorizontalDivider(0, regColsTaken, gameEnd)

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
				cellColor := standard
				if i == ActiveSectorRow && j == ActiveSectorCol && !SectorBlocked && !gameEnd {
					cellColor = active
				} else if regular.Board[i][j] == config.Settings.Player1.Symbol {
					cellColor = player1()
				} else if regular.Board[i][j] == config.Settings.Player2.Symbol {
					cellColor = player2()
				} else if regular.Board[i][j] == "-" {
					cellColor = tie
				}
				printSubBoardRow(GameBoard.Cells[i][j].Cells[subRow], i, j, cellColor)
			}
			fmt.Println() // move to the next line after printing a sub-row

			// print the horizontal dividers between sub-rows
			if subRow < subRows-1 {
				printSubHorizontalDivider(i, regColsTaken, gameEnd)
			}
		}

		// print the horizontal dividers between rows of the SuperBoard, except for the last row
		if i < rows-1 {
			fmt.Println(color.InBold("  ===+===+==== | ====+==== | ====+===+==="))
		}
	}
	printSubHorizontalDivider(2, regColsTaken, gameEnd)
	fmt.Println("") // move to the next line after printing the SuperBoard
}

func printSubHorizontalDivider(i int, regColsTaken []int, gameEnd bool) {
	first := standard
	second := standard
	third := standard

	if !SectorBlocked && !gameEnd {
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
			if regular.Board[i][regColsTaken[j]] != " " {
				var blockedSectorColor string
				if regular.Board[i][regColsTaken[j]] == config.Settings.Player1.Symbol {
					blockedSectorColor = player1()
				} else if regular.Board[i][regColsTaken[j]] == config.Settings.Player2.Symbol {
					blockedSectorColor = player2()
				} else if regular.Board[i][regColsTaken[j]] == "-" {
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
	fmt.Println(color.With(first, "   +---+---+---") + color.With(accent, "|") + color.With(second,
		"---+---+---") + color.With(accent, "|") + color.With(third, "---+---+---+"))
}

// prints respective row of a subboard
// i dont remember what the fuck is going on here, but it does the thing said above
func printSubBoardRow(subBoardRow [3]string, i int, j int, sectorColor string) {
	for x, cell := range subBoardRow {
		cellPart := "" // what the hell is a cell part
		if x < 1 && j > 0 {
			cellPart = color.With(accent, " | %-3s")
		} else {
			cellPart = color.With(sectorColor, " | %-3s")
		}

		if regular.Board[i][j] == config.Settings.Player1.Symbol {
			fmt.Printf(cellPart, color.With(player1(), cell))
			if j > 1 && x > 1 {
				fmt.Print(color.With(player1(), " |"))
			}
		} else if regular.Board[i][j] == config.Settings.Player2.Symbol {
			fmt.Printf(cellPart, color.With(player2(), cell))
			if j > 1 && x > 1 {
				fmt.Print(color.With(player2(), " |"))
			}
		} else if regular.Board[i][j] == "-" {
			fmt.Printf(cellPart, color.With(tie, cell))
			if j > 1 && x > 1 {
				fmt.Print(color.With(tie, " |"))
			}
		} else {
			if strings.HasSuffix(cell, config.Settings.Player1.Symbol) {
				fmt.Printf(cellPart, color.With(player1(), cell))
			} else {
				fmt.Printf(cellPart, color.With(player2(), cell))
			}
			if j > 1 && x > 1 {
				fmt.Print(color.With(sectorColor, " |"))
			}
		}
	}
}

// GetSuperPlayerMove - get the move input from player in the Super mode
func GetSuperPlayerMove(player int, availableMoves int, availableBoards int) (int, int, bool, error) {
getSuperPlayerMoveFunctionStart:
	fmt.Printf("Player %d, enter your move (e.g., A1, B2) or (Q) to quit the game: ", player)
	var move string
	_, err := fmt.Scan(&move)
	if err != nil {
		return 0, 0, false, err
	}

	// if player chose to quit, confirm if they are sure
	// if they change their mind, restart this function from the beginning
	if strings.ToLower(move) == "q" {
		confirmQuit, err := utils.ConfirmQuit()
		if err != nil {
			return 0, 0, false, err
		}
		if confirmQuit {
			return 0, 0, true, nil
		} else {
			goto getSuperPlayerMoveFunctionStart
		}
	}

	row, col, valid := regular.ParseMoveInput(move)
	if !valid || GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		_, _, _, err := GetSuperPlayerMove(player, availableMoves, availableBoards)
		if err != nil {
			return 0, 0, false, err
		}
		return row, col, false, nil
	}
	GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = config.GetPlayerSymbol(player)
	return row, col, false, nil
}

// GetSuperComputerMove - decide move selection for the computer in Super mode
// will also decide next sector/subboard if necessary
func GetSuperComputerMove(player int) (int, int) {
	// loop until a move is made
	for {
		// decide move if sector/subboard is playable
		if !SectorBlocked {
			for {
				row := rand.Intn(3)
				col := rand.Intn(3)
				if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] == " " {
					GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = config.GetPlayerSymbol(player)

					return row, col
				}
			}
			// select a new sector/subboard
		} else {
			for {
				row := rand.Intn(3)
				col := rand.Intn(3)
				if regular.Board[row][col] == " " {
					ActiveSectorRow = row
					ActiveSectorCol = col
					SectorBlocked = false

					return row, col
				}
			}
		}
	}
}

// GetSectorMove - get player input for the next sector/subboard to make a move in
func GetSectorMove(player int, availableMoves int, availableBoards int) (bool, error) {
getSectorMoveFunctionStart:
	fmt.Printf("Player %d, which sector would you like to move in (e.g., A1, B2, <A - C><1 - 3>) or (Q) to quit the game: ", player)
	var move string
	_, err := fmt.Scan(&move)
	if err != nil {
		return false, err
	}

	// if player chose to quit, confirm if they are sure
	// if they change their mind, restart this function from the beginning
	if strings.ToLower(move) == "q" {
		confirmQuit, err := utils.ConfirmQuit()
		if err != nil {
			return false, err
		}
		if confirmQuit {
			return true, nil
		} else {
			goto getSectorMoveFunctionStart
		}
	}

	sectorRow, sectorCol, sectorValid := regular.ParseMoveInput(move)
	// loop this function a valid input is received
	if !sectorValid || regular.Board[sectorRow][sectorCol] != " " {
		fmt.Println("Invalid sector, Try again.")
		_, err := GetSectorMove(player, availableMoves, availableBoards)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	// set the chosen sector/subboard
	ActiveSectorRow = sectorRow
	ActiveSectorCol = sectorCol
	SectorBlocked = false

	return false, nil
}

// ProcessMoveAndUpdateGameState - THIS FUNCTION IS ONLY USED IN THE SUPER MODE
func ProcessMoveAndUpdateGameState(row int, col int, player int, availableMoves int, availableBoards int) (int, int) {
	// decrement remaining moves
	availableMoves--

	openSubBoards := GetEmptySubBoards()
	if availableMoves <= 76 {
		if regular.CheckForWin(player, GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells) {
			// if a win is found, loop through remaning spaces that are still open in that board and remove that many turns from the remaining moves counter.
			// those empty spaces are also filled in with a "-".
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[i][j] == " " {
						availableMoves--

						GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[i][j] = "-"
					}
				}
			}
			// set the subboard that was won to the respective player on the super-board and decrement 1 from the remaiming boards count
			regular.Board[ActiveSectorRow][ActiveSectorCol] = config.GetPlayerSymbol(player)
			availableBoards--
		} else {
			// if no win, check if that sub-board can still be played
			// if there are no spaces left, then it cannot be played and its spot is
			// filled in the super-board to block it out
			openSubSpaces := GetEmptySpaces()
			if openSubSpaces == 0 {
				regular.Board[ActiveSectorRow][ActiveSectorCol] = "-"
			}
		}

		// set the variable to indicate whether or not the next sector/subboard cannot be in play
		// the game will use this variable to determine if it needs to ask the player/computer to choose a different sector/subboard
		if regular.Board[row][col] == " " && openSubBoards >= 1 {
			SectorBlocked = false
		} else {
			SectorBlocked = true
		}
	}
	// set the next active sector/subboard
	// it does not matter if it can be in play or not, the game logic will check that
	if openSubBoards >= 1 {
		ActiveSectorRow = row
		ActiveSectorCol = col
	}
	return availableMoves, availableBoards
}
