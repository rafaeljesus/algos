package tictactoe

type TicTacToe struct {
	Size  int
	Board [][]int
}

func NewTicTacToe(n int) *TicTacToe {
	board := make([][]int, n)
	// https://golang.org/doc/effective_go#slices
	for i := range board {
		board[i] = make([]int, n)
	}
	return &TicTacToe{
		Size:  n,
		Board: board,
	}
}

func (t *TicTacToe) Move(row, col, player int) int {
	t.Board[row][col] = player
	if t.isRowWinner(row, player) ||
		t.isColumnWinner(col, player) ||
		t.isDiogonalWinner(player) ||
		t.isAntiDiogonalWinner(player) {
		return player
	}
	return -1
}

// [1], [1], [1]
// [ ], [ ], [ ]
// [ ], [ ], [ ]
func (t *TicTacToe) isRowWinner(row, player int) bool {
	col := t.Board[row]
	for colIdx, _ := range col {
		if t.Board[row][colIdx] != player {
			return false
		}
	}
	return true
}

// [1] [ ] [ ]
// [1] [ ] [ ]
// [1] [ ] [ ]
func (t *TicTacToe) isColumnWinner(col, player int) bool {
	for rowIdx, _ := range t.Board {
		if t.Board[rowIdx][col] != player {
			return false
		}
	}
	return true
}

// [1] [ ] [ ]
// [ ] [1] [ ]
// [ ] [ ] [1]
func (t *TicTacToe) isDiogonalWinner(player int) bool {
	for rowIdx, _ := range t.Board {
		if t.Board[rowIdx][rowIdx] != player {
			return false
		}
	}
	return true
}

// [ ] [ ] [1]
// [ ] [1] [ ]
// [1] [ ] [ ]
func (t *TicTacToe) isAntiDiogonalWinner(player int) bool {
	for rowIdx, row := range t.Board {
		if t.Board[rowIdx][len(row)-1-rowIdx] != player {
			return false
		}
	}
	return true
}
