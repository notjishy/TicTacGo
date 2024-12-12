package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

func SwitchPlayer(playerCount int, player int) int {
	if player == 1 {
		return 2
	}
	return 1
}

func playerSymbol(player int) string {
	if player == 1 {
		return "X"
	}
	return "O"
}

func GetRegPlayerMove(player int, board [3][3]string) {
	fmt.Printf("Player %d, enter your move (e.g., A1, B2): ", player)
	var move string
	fmt.Scan(&move)
	row, col, valid := ParseMove(move)
	if !valid || board[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		GetRegPlayerMove(player, board)
		return
	}
	Board[row][col] = playerSymbol(player)
}

func GetRegComputerMove() {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if Board[row][col] == " " {
			Board[row][col] = "O"
			return
		}
	}
}

func GetSuperPlayerMove(player int, availableMoves int, availableBoards int) (int, int) {
	fmt.Printf("Player %d, enter your move (e.g., A1, B2): ", player)
	var move string
	fmt.Scan(&move)
	row, col, valid := ParseMove(move)
	if !valid || GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		GetSuperPlayerMove(player, availableMoves, availableBoards)
		return row, col
	}
	GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = playerSymbol(player)
	return row, col
}

func GetSuperComputerMove(player int, availableMoves int, availableBoards int) (int, int) {
	for {
		if !SectorBlocked {
			for {
				row := rand.Intn(3)
				col := rand.Intn(3)
				if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] == " " {
					GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = playerSymbol(player)

					return row, col
				}
			}
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

func GetSectorMove(player int, availableMoves int, availableBoards int) {
	fmt.Printf("Player %d, which sector would you like to move in (e.g., A1, B2, <A - C><1 - 3>): ", player)
	var move string
	fmt.Scan(&move)
	sectorRow, sectorCol, sectorValid := ParseMove(move)
	if !sectorValid || Board[sectorRow][sectorCol] != " " {
		fmt.Println("Invalid sector, Try again.")
		GetSectorMove(player, availableMoves, availableBoards)
		return
	}
	ActiveSectorRow = sectorRow
	ActiveSectorCol = sectorCol
}

func ParseMove(move string) (int, int, bool) {
	move = strings.ToLower(move)

	if len(move) != 2 {
		return 0, 0, false
	}
	row := int(move[0] - 'a')
	col := int(move[1] - '1')
	return row, col, row >= 0 && row < 3 && col >= 0 && col < 3
}

func UpdateGameState(row int, col int, player int, availableMoves int, availableBoards int) (int, int) {
	availableMoves--

	openSubBoards := GetEmptySubBoards()
	if availableMoves <= 76 {
		if CheckWin(player, GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells) {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[i][j] == " " {
						availableMoves--

						GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[i][j] = "-"
					}
				}
			}
			Board[ActiveSectorRow][ActiveSectorCol] = playerSymbol(player)
			availableBoards--
		} else {
			openSubSpaces := GetEmptySpaces()
			if openSubSpaces == 0 {
				Board[ActiveSectorRow][ActiveSectorCol] = "-"
			}
		}

		if Board[row][col] == " " && openSubBoards >= 1 {
			SectorBlocked = false
		} else {
			SectorBlocked = true
		}
	}
	if openSubBoards >= 1 {
		ActiveSectorRow = row
		ActiveSectorCol = col
	}
	return availableMoves, availableBoards
}

func CheckWin(player int, Board [3][3]string) bool {
	symbol := playerSymbol(player)
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
