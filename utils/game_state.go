package utils

// THIS FUNCTION IS ONLY USED IN THE SUPER MODE
func ProcessMoveAndUpdateGameState(row int, col int, player int, availableMoves int, availableBoards int) (int, int) {
	// decrement remaining moves
	availableMoves--

	openSubBoards := GetEmptySubBoards()
	if availableMoves <= 76 {
		if CheckForWin(player, GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells) {
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
			Board[ActiveSectorRow][ActiveSectorCol] = GetPlayerSymbol(player)
			availableBoards--
		} else {
			// if no win, check if that sub-board can still be played
			// if there are no spaces left, then it cannot be played and its spot is
			// filled in the super-board to block it out
			openSubSpaces := GetEmptySpaces()
			if openSubSpaces == 0 {
				Board[ActiveSectorRow][ActiveSectorCol] = "-"
			}
		}

		// set the variable to indicate whether or not the next sector/subboard cannot be in play
		// the game will use this variable to determine if it needs to ask the player/computer to choose a different sector/subboard
		if Board[row][col] == " " && openSubBoards >= 1 {
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

// check for a win in a board. this function is utilized by both modes
func CheckForWin(player int, Board [3][3]string) bool {
	symbol := GetPlayerSymbol(player)
	// Check rows
	for i := 0; i < 3; i++ {
		if Board[i][0] == symbol && Board[i][1] == symbol && Board[i][2] == symbol {
			return true
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if Board[0][i] == symbol && Board[1][i] == symbol && Board[2][i] == symbol {
			return true
		}
	}
	// Check diagonals
	if (Board[0][0] == symbol && Board[1][1] == symbol && Board[2][2] == symbol) ||
		(Board[0][2] == symbol && Board[1][1] == symbol && Board[2][0] == symbol) {
		return true
	}
	return false
}
