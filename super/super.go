package super

import (
	"fmt"
	"strings"
	"math/rand"
	"tictacgo/utils"
	"tictacgo/regular"
)

var player = 1

var ActiveSectorRow = 0
var ActiveSectorCol = 0

var sectorBlocked = false

var availableMoves = 81
var availableBoards = 9

func Play(playerCount int) {
	InitializeSuperBoard()
	for availableMoves > 0 {
		if player == 1 || (playerCount == 2 && player == 2) {
			PrintSuperBoard(availableMoves, sectorBlocked)
			if availableMoves == 81 || sectorBlocked {
				getSectorMove()
			} else {
				getPlayerMove()
			}
		} else {
			getComputerMove()
		}
		player = utils.SwitchPlayer(playerCount, player)
	}
}

func getSectorMove() {
	fmt.Printf("Player %d, which sector would you like to move in (e.g., A1, B2, <A - C><1 - 3>): ", player)
	var move string
	fmt.Scan(&move)
	sectorRow, sectorCol, sectorValid := parseMove(move)
	if !sectorValid || regular.Board[sectorRow][sectorCol] != " " {
		fmt.Println("Invalid sector, Try again.")
		getSectorMove()
		return
	}
	ActiveSectorRow = sectorRow
	ActiveSectorCol = sectorCol
	getPlayerMove()
}

func getPlayerMove() {
	fmt.Printf("Player %d, enter your move (e.g., A1, B2): ", player)
	var move string
	fmt.Scan(&move)
	row, col, valid := parseMove(move)
	if !valid || GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		getPlayerMove()
		return
	}
	GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = utils.PlayerSymbol(player)
	updateGameBoard(row, col)
	
	ActiveSectorRow = row
	ActiveSectorCol = col
}

func getComputerMove() {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] == " " {
			GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = "O"
			
			updateGameBoard(row, col)

			ActiveSectorRow = row
			ActiveSectorCol = col
			return
		}
	}
}

func parseMove(move string) (int, int, bool) {
	move = strings.ToLower(move)

	if len(move) != 2 {
		return 0, 0, false
	}
	row := int(move[0] - 'a')
	col := int(move[1] - '1')
	return row, col, row >= 0 && row < 3 && col >= 0 && col < 3
}

func updateGameBoard(row int, col int) {
	availableMoves--
	if availableMoves < 76 {
		if utils.CheckWin(player, GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells) {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[i][j] == " " {
						availableMoves--

						GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[i][j] = "-"
					}
				}
			}
			regular.Board[ActiveSectorRow][ActiveSectorCol] = utils.PlayerSymbol(player)
			availableBoards--
		}

		if regular.Board[row][col] == "X" || regular.Board[row][col] == "O" {
			sectorBlocked = true
		} else { sectorBlocked = false }
	}
}
