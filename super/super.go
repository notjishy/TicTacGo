package super

import (
	"fmt"
	"strings"
	"math/rand"
	"tictacgo/utils"
	"tictacgo/regular"
)

var player int
var ActiveSectorRow int
var ActiveSectorCol int
var availableMoves int
var availableBoards int

var sectorBlocked bool
var gameEnd bool = false

func Play(playerCount int) {
	InitializeSuperBoard()
	availableMoves = 81
	availableBoards = 9
	player = 1
	sectorBlocked = false
	for availableMoves > 0 {
		if player == 1 || (playerCount == 2 && player == 2) {
			PrintSuperBoard(availableMoves, sectorBlocked, gameEnd)
			if availableMoves == 81 || sectorBlocked {
				getSectorMove()
			} else {
				getPlayerMove()
			}
		} else {
			getComputerMove()
		}

		if availableBoards < 7 {
			gameEnd = utils.CheckWin(player, regular.Board)
			if gameEnd {
				PrintSuperBoard(availableMoves, sectorBlocked, gameEnd)
				fmt.Printf("Player %d wins!\n", player)
				return
			}
		}

		player = utils.SwitchPlayer(playerCount, player)
	}
	PrintSuperBoard(availableMoves, sectorBlocked, gameEnd)
	fmt.Println("It's a tie!")
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
	updateGameState(row, col)
}

func getComputerMove() {
	for {
		if !sectorBlocked {
			for {
				row := rand.Intn(3)
				col := rand.Intn(3)
				if GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] == " " {
					GameBoard.Cells[ActiveSectorRow][ActiveSectorCol].Cells[row][col] = "O"
					
					updateGameState(row, col)
					return
				}
			}
			return
		} else {
			for {
				row := rand.Intn(3)
				col := rand.Intn(3)
				if regular.Board[row][col] == " " {
					ActiveSectorRow = row
					ActiveSectorCol = col
					sectorBlocked = false
					break
				}
			}
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

func updateGameState(row int, col int) {
	availableMoves--

	openSubBoards := GetEmptySubBoards()
	if availableMoves <= 76 {
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
		} else {
			openSubSpaces := GetEmptySpaces()
			if openSubSpaces == 0 {
				regular.Board[ActiveSectorRow][ActiveSectorCol] = "-"
			}
		}
		
		if ((regular.Board[row][col] == "X" || regular.Board[row][col] == "O" || regular.Board[row][col] == "-") && openSubBoards >= 1) {
			sectorBlocked = true
		} else { sectorBlocked = false }
	}

	if openSubBoards >= 1 {
		ActiveSectorRow = row
		ActiveSectorCol = col
	}
}
