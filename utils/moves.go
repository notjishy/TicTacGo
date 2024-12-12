package utils

import (
	"fmt"
	"math/rand"
)

// get the move input from a player in the Regular mode
func GetRegularPlayerMove(player int, board [3][3]string) {
	fmt.Printf("Player %d, enter your move (e.g., A1, B2): ", player)
	var move string
	fmt.Scan(&move)
	row, col, valid := ParsePlayerMove(move)
	// loop this function until a valid input is received
	if !valid || board[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		GetRegularPlayerMove(player, board)
		return
	}
	// set the players move to the board
	Board[row][col] = GetPlayerSymbol(player)
}

// decide the move selection for the computer in the Regular mode
func GetRegularComputerMove() {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if Board[row][col] == " " {
			// set the move to the board
			Board[row][col] = GetPlayerSymbol
			return
		}
	}
}

// get the move input from player in the Super mode
func GetSuperPlayerMove(player int, availableMoves int, availableBoards int) (int, int) {
	fmt.Printf("Player %d, enter your move (e.g., A1, B2): ", player)
	var move string
	fmt.Scan(&move)
	row, col, valid := ParsePlayerMove(move)
	if !valid || GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		GetSuperPlayerMove(player, availableMoves, availableBoards)
		return row, col
	}
	GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = GetPlayerSymbol(player)
	return row, col
}

// decide move selection for the computer in Super mode
// will also decide next sector/subboard if necessary
func GetSuperComputerMove(player int, availableMoves int, availableBoards int) (int, int) {
	// loop until a move is made
	for {
		// decide move if sector/subboard is playable
		if !SectorBlocked {
			for {
				row := rand.Intn(3)
				col := rand.Intn(3)
				if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] == " " {
					GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = GetPlayerSymbol(player)

					return row, col
				}
			}
		// select a new sector/subboard
		} else {
			for {
				row := rand.Intn(3)
				col := rand.Intn(3)
				if Board[row][col] == " " {
					ActiveSectorRow = row
					ActiveSectorCol = col
					SectorBlocked = false

					return row, col
				}
			}
		}
	}
}

// get player input for the next sector/subboard to make a move in
func GetSectorMove(player int, availableMoves int, availableBoards int) {
	fmt.Printf("Player %d, which sector would you like to move in (e.g., A1, B2, <A - C><1 - 3>): ", player)
	var move string
	fmt.Scan(&move)
	sectorRow, sectorCol, sectorValid := ParsePlayerMove(move)
	// loop this function a valid input is received
	if !sectorValid || Board[sectorRow][sectorCol] != " " {
		fmt.Println("Invalid sector, Try again.")
		GetSectorMove(player, availableMoves, availableBoards)
		return
	}
	// set the chosen sector/subboard
	ActiveSectorRow = sectorRow
	ActiveSectorCol = sectorCol
}
