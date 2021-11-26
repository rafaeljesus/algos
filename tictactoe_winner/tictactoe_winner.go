package tictactoe_winner

const (
	boardSize = 3
	pending   = "Pending"
	draw      = "Draw"
)

func FindTicTacToeWinner(movesArr [][]int) string {
	matrix := make([][]string, boardSize)
	for i, _ := range matrix {
		matrix[i] = make([]string, boardSize)
	}
	matrixLen := pow(len(matrix))
	var count int
	for i, moves := range movesArr {
		count++
		player := toPlayerStr(i % 2)
		row, col := moves[0], moves[1]
		matrix[row][col] = player
		if isRowWinner(matrix, player, row) ||
			isColWinner(matrix, player, col) ||
			isDiagonalWinner(matrix, player) ||
			isAntiDiagonalWinner(matrix, player) {
			return player
		}
	}
	if count <= matrixLen-1 {
		return pending
	}
	return draw
}

func isRowWinner(matrix [][]string, player string, row int) bool {
	for _, r := range matrix {
		for colIdx, _ := range r {
			if matrix[row][colIdx] != player {
				return false
			}
		}
	}
	return true
}

func isColWinner(matrix [][]string, player string, col int) bool {
	for rowIdx, _ := range matrix {
		if matrix[rowIdx][col] != player {
			return false
		}
	}
	return true
}

func isDiagonalWinner(matrix [][]string, player string) bool {
	for rowIdx, _ := range matrix {
		if matrix[rowIdx][rowIdx] != player {
			return false
		}
	}
	return true
}

func isAntiDiagonalWinner(matrix [][]string, player string) bool {
	for rowIdx, row := range matrix {
		if matrix[rowIdx][len(row)-1-rowIdx] != player {
			return false
		}
	}
	return true
}

func toPlayerStr(player int) string {
	if player == 0 {
		return "A"
	}
	return "B"
}

func pow(value int) int {
	return value * value
}
