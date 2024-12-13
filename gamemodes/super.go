package gamemodes

import (
	"fmt"
	"log"
	"tictacgo/utils"
	
	"github.com/eiannone/keyboard"
)

// initialize variables
// keep track of where moves are being made
// used later when gameboard is updated
var row int
var col int

// number of boards in the game (it will always print out 9 boards,
// this just is just to track how many are left in play)
var availableBoards int

// game will end when this becomes true
var gameEnd bool = false

func PlaySuper(playerCount int) {
	utils.InitializeSuperBoard()
	// set variables at start of game
	player = 1
	availableMoves = 81
	availableBoards = 9
	utils.SectorBlocked = true

	// loop through game until until no more moves left and game ties
	for availableMoves > 0 {
		// only ask for player's move if the current turn is for an actual person.
		// i.e. if there is only 1 player, do not ask for user input if it isn't their turn.
		if player == 1 || (playerCount == 2 && player == 2) {
			utils.PrintSuperBoard(availableMoves, gameEnd)
			// ask player which board to play in if the current selected board is no longer in play (has been won/tied)
			if availableMoves == 81 || utils.Board[row][col] != " " {
				// returns boolean value indicating if the player has quit the game
				didPlayerQuit = utils.GetSectorMove(player, availableMoves, availableBoards)
				// if player has quit the game, force end this gameloop
				if didPlayerQuit { break }

				// print the board again (so the player can see the highlighted sector/subboard)
				utils.PrintSuperBoard(availableMoves, gameEnd)
			}
			// acquire move from player
			row, col, didPlayerQuit = utils.GetSuperPlayerMove(player, availableMoves, availableBoards)

			if didPlayerQuit { break }
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
				utils.PrintSuperBoard(availableMoves, gameEnd)
				fmt.Printf("Player %d wins!\n", player)
				// wait for user to press a key before returning to main menu
				fmt.Print("Press any key to go back to main menu...")
				err := keyboard.Open() // begin keyboard listening
				if err != nil { log.Fatal(err) }
				_, _, err = keyboard.GetKey() // get the key that is pressed (we arent storing it we dont need to know what it is)
				if err != nil { log.Fatal(err) }
				defer keyboard.Close() // end keyboard listening
				// force end game
				break
			}
		}

		// swap to next player after turn is finished
		player = utils.SwitchPlayer(playerCount, player)
	}
	// only display tie message if player did not quit the game
	if !didPlayerQuit {
		utils.PrintSuperBoard(availableMoves, gameEnd)
		fmt.Println("It's a tie!")
	}
}
