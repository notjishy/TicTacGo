package utils

func PlayerSymbol(player int) string {
	if player == 1 {
		return "X"
	}
	return "O"
}

func CheckWin(player int, Board [3][3]string) bool {
	symbol := PlayerSymbol(player)
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
