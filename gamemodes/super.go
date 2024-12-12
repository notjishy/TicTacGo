package gamemodes

import (
	"fmt"
	"tictacgo/utils"
)

// initialize variables
// `player` variable indicates which player's turn it is.
// defaulting to player 1 goes first.
var player int = 1

// keep track of where moves are being made
// used later when gameboard is updated
var row int
var col int

// maximum amount of moves until game ends in a tie
var availableMoves int

// number of boards in the game (it will always print out 9 boards,
// this just is just to track how many are left in play)
var availableBoards int

// game will end when this becomes true
var gameEnd bool = false

func PlaySuper(playerCount int) {
	utils.InitializeSuperBoard()
	// set variables at start of game
	availableMoves = 81
	availableBoards = 9

	// loop through game until until no more moves left and game ties
	for availableMoves > 0 {
		// only ask for player's move if the current turn is for an actual person.
		// i.e. if there is only 1 player, do not ask for user input if it isn't their turn.
		if player == 1 || (playerCount == 2 && player == 2) {
			utils.PrintSuperBoard(availableMoves, utils.SectorBlocked, gameEnd)
			// ask player which board to play in if the current selected board is no longer in play (has been won/tied)
			if availableMoves == 81 || utils.Board[row][col] != " " {
				utils.GetSectorMove(player, availableMoves, availableBoards)
			}
			// acquire move from player
			row, col = utils.GetSuperPlayerMove(player, availableMoves, availableBoards)
		} else {
			// acquire move from computer
			row, col = utils.GetSuperComputerMove(player, availableMoves, availableBoards)
		}

		// update gameboard state (checks for wins in remaining boards, updates the active board, decrements remining moves and boards)
		availableMoves, availableBoards = utils.ProcessMoveAndUpdateGameState(row, col, player, availableMoves, availableBoards)

		// check main board for win condition.
		// if >= 7 boards remaining, no need to check as a win is impossible there
		if availableBoards < 7 {
			if utils.CheckForWin(player, utils.Board) {
				utils.PrintSuperBoard(availableMoves, utils.SectorBlocked, gameEnd)
				fmt.Printf("Player %d wins!\n", player)
				// force game to end if there is winner
				return
			}
		}

		// swap to next player after turn is finished
		player = utils.SwitchPlayer(playerCount, player)
	}
	utils.PrintSuperBoard(availableMoves, utils.SectorBlocked, gameEnd)
	fmt.Println("It's a tie!")
}
