package super

import (
	"fmt"
	"strings"
	"tictacgo/utils"
)

var activeSectorRow = 0
var activeSectorCol = 0

func Play(playerCount int) {
	InitializeSuperBoard()
	player := 1
	availableMoves := 81
	for availableMoves > 0 {
		if player == 1 || (playerCount == 2 && player == 2) {
			PrintSuperBoard()
			if availableMoves == 81 {
				getInitialMove()
			} else {
				getPlayerMove(player, activeSectorRow, activeSectorCol)
			}
		}
		availableMoves--
		player = utils.SwitchPlayer(playerCount, player)
	}

	// fmt.Println("This mode is not currently functional,\nbut here is a preview of what the game looks like so far!\n\nAnyways...")
}

func getInitialMove() {
	fmt.Printf("Player 1, which sector would you like to start in (e.g., A1, B2, <A - C><1 - 3>)")
	var move string
	fmt.Scan(&move)
	sectorRow, sectorCol, sectorValid := parseMove(move)
	if !sectorValid {
		fmt.Println("Invalid sector, Try again.")
		getInitialMove()
		return
	}
	getPlayerMove(1, sectorRow, sectorCol)
}

func getPlayerMove(player, sectorRow, sectorCol int) {
	fmt.Printf("Player %d, enter your move (e.g., A1, B2): ", player)
	var move string
	fmt.Scan(&move)
	row, col, valid := parseMove(move)
	if !valid || GameBoard.Cells[sectorRow][sectorCol].Cells[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		getPlayerMove(player, sectorRow, sectorCol)
		return
	}
	GameBoard.Cells[sectorRow][sectorCol].Cells[row][col] = utils.PlayerSymbol(player)
	activeSectorRow = row
	activeSectorCol = col
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
