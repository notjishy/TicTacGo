package game

import (
	"fmt"
	"tictacgo/config"
	"tictacgo/game/board"
	"tictacgo/utils"

	"github.com/eiannone/keyboard"
)

func PlayRegular(playerCount int) error {
	board.InitializeBoard()
	// set variables at start of game
	player = 1
	availableMoves = 9

	player1Color := config.Settings.Player1.GetColor()
	player2Color := config.Settings.Player2.GetColor()

	// loop game until no more moves left
	// if availableMoves runs out, the game is a tie
	for availableMoves > 0 {
		// only ask for player's move if the current turn is for an actual person.
		// i.e. if there is only 1 player, do not ask for user input if it isn't their turn.
		if player == 1 || (playerCount == 2 && player == 2) {
			board.PrintBoard(player1Color, player2Color)
			// acquire move from player. returns boolean value indicating if player quit the game or not
			didPlayerQuit, err := board.GetPlayerMove(player, board.Grid)
			if err != nil {
				return err
			}
			// if player chose to quit the game, force end this gameloop
			if didPlayerQuit {
				break
			} // force end game and go back to main
		} else {
			board.GetComputerMove(player)
		}
		// decrement moves remaining
		availableMoves--

		// check board for win conditions.
		// if moves is >= 5, no need to check as it would be impossible
		if availableMoves < 5 {
			if board.CheckForWin(player, board.Grid) {
				board.PrintBoard(player1Color, player2Color)
				fmt.Printf("Player %d wins!\n", player)
				// wait for user to press a key before returning to main menu
				fmt.Print("Press any key to go back to main menu...")

				err := keyboard.Open() // begin keyboard listening
				if err != nil {
					return err
				}

				// get key that is pressed, not storing it, no need to know what it is exactly
				_, _, err = keyboard.GetKey()
				if err != nil {
					return err
				}

				err = keyboard.Close() // end keyboard listening
				if err != nil {
					return err
				}
				// force end game
				return nil
			}
		}
		// swap to next player after turn is finished
		player = utils.SwitchPlayer(player)
	}
	// only print message if player did not quit the game
	if !didPlayerQuit {
		board.PrintBoard(player1Color, player2Color)
		fmt.Println("It's a tie!")
	}
	return nil
}
