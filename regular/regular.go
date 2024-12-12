package regular

import (
	"fmt"
	"tictacgo/utils"
)

// initialize variables
// `player` variable indicates which player's turn it is,
// starting with player 1 going first
var player int

// maximum number of turns in the game before it force ends
var availableMoves int

func Play(playerCount int) {
	player = 1
	availableMoves = 9
	utils.InitializeBoard()

	// loop game until no more moves left
	// if availableMoves runs out, the game is a tie
	for availableMoves > 0 {
		// only ask for player's move if the current turn is for an actual person.
		// i.e. if there is only 1 player, do not ask for user input if it isn't their turn.
		if player == 1 || (playerCount == 2 && player == 2) {
			utils.PrintBoard()
			// acquire move from player.
			utils.GetRegPlayerMove(player, utils.Board)
		} else {
			utils.GetRegComputerMove()
		}
		// decrement moves remaining
		availableMoves--

		// check board for win conditions.
		// if moves is >= 5, no need to check as it would be impossible
		if availableMoves < 5 {
			if utils.CheckWin(player, utils.Board) {
				utils.PrintBoard()
				fmt.Printf("Player %d wins!\n", player)
				// force end game if someone won
				return
			}
		}
		// swap to next player after turn is finished
		player = utils.SwitchPlayer(playerCount, player)
	}
	utils.PrintBoard()
	fmt.Println("It's a tie!")
}
