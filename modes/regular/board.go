package regular

import (
	"fmt"
	"math/rand"
	"strings"
	"tictacgo/utils"

	"tictacgo/config"

	"github.com/TwiN/go-color"
	"github.com/inancgumus/screen"
)

// Board - define the board
var Board [3][3]string

// InitializeBoard - create the board
func InitializeBoard() {
	for row := range Board {
		for col := range Board[row] {
			Board[row][col] = " "
		}
	}
}

// PrintBoard - print the board in the terminal
func PrintBoard(player1Color string, player2Color string) {
	// clear the screen
	screen.Clear()

	fmt.Println("\n    1   2   3")
	// print out 1 row at a time
	for i, row := range Board {
		// + i to 'a' so the row letter increments up.
		// row 1 = a, row 2 = b, row 3 = c
		fmt.Print(string(rune('a' + i)))
		for _, cell := range row {
			if strings.HasSuffix(cell, config.Settings.Player1.Symbol) {
				fmt.Printf(" | %s", color.With(player1Color, cell))
			} else {
				fmt.Printf(" | %s", color.With(player2Color, cell))
			}
		}
		fmt.Println(" |")
		// print horizontal border
		if i < 2 {
			fmt.Println("   ---+---+---")
		}
	}
	fmt.Println()
}

// CheckForWin - check for a win in a board. this function is utilized by both modes
func CheckForWin(player int, Board [3][3]string) bool {
	symbol := config.GetPlayerSymbol(player)
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

// GetPlayerMove - get the move input from a player in the Regular mode
func GetPlayerMove(player int, board [3][3]string) (bool, error) {
getRegularPlayerMoveFunctionStart:
	fmt.Printf("Player %d, enter your move (e.g., A1, B2) or (Q) to quit the game: ", player)
	var move string
	_, err := fmt.Scan(&move)
	if err != nil {
		return false, err
	}

	// if player chose to quit, confirm if they are sure
	// if they change their mind, restart this function from the beginning
	if strings.ToLower(move) == "q" {
		confirmQuit, err := utils.ConfirmQuit()
		if err != nil {
			return false, err
		}
		if confirmQuit {
			return true, nil
		} else {
			goto getRegularPlayerMoveFunctionStart
		}
	}

	row, col, valid := ParseMoveInput(move)
	// loop this function until a valid input is received
	if !valid || board[row][col] != " " {
		fmt.Println("Invalid move. Try again.")
		_, err := GetPlayerMove(player, board)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	// set the players move to the board
	Board[row][col] = config.GetPlayerSymbol(player)
	return false, nil
}

// GetComputerMove - decide the move selection for the computer in the Regular mode
func GetComputerMove(player int) {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if Board[row][col] == " " {
			// set the move to the board
			Board[row][col] = config.GetPlayerSymbol(player)
			return
		}
	}
}

// ParseMoveInput - parse the move input from the player
func ParseMoveInput(move string) (int, int, bool) {
	move = strings.ToLower(move)

	if len(move) != 2 {
		return 0, 0, false
	}
	row := int(move[0] - 'a')
	col := int(move[1] - '1')
	return row, col, row >= 0 && row < 3 && col >= 0 && col < 3
}
