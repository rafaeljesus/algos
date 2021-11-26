package tictactoe_winner

import (
	"testing"
)

func TestFindWinner(t *testing.T) {
	var tests = []struct {
		moves [][]int
		out   string
	}{
		{
			[][]int{
				[]int{0, 0},
				[]int{2, 0},
				[]int{1, 1},
				[]int{2, 1},
				[]int{2, 2},
			},
			"A",
		},
		{
			[][]int{
				[]int{0, 2},
				[]int{2, 2},
				[]int{1, 1},
				[]int{2, 1},
				[]int{2, 0},
			},
			"A",
		},
		{
			[][]int{
				[]int{0, 0},
				[]int{0, 1},
				[]int{1, 2},
				[]int{1, 1},
				[]int{2, 2},
				[]int{2, 1},
			},
			"B",
		},
		{
			[][]int{
				[]int{0, 0},
				[]int{1, 1},
				[]int{2, 0},
				[]int{1, 0},
				[]int{1, 2},
				[]int{2, 1},
				[]int{0, 1},
				[]int{0, 2},
				[]int{2, 2},
			},
			draw,
		},
		{
			[][]int{
				[]int{0, 0},
				[]int{1, 1},
			},
			pending,
		},
		{
			[][]int{
				[]int{1, 0},
				[]int{2, 2},
				[]int{2, 0},
				[]int{0, 1},
				[]int{1, 1},
			},
			pending,
		},
		{
			[][]int{
				//[[2,0],[0,2],[0,0],[2,2],[1,1],[1,0],[2,1],[0,1]]
				[]int{2, 0},
				[]int{0, 2},
				[]int{0, 0},
				[]int{2, 2},
				[]int{1, 1},
				[]int{1, 0},
				[]int{2, 1},
				[]int{0, 1},
			},
			pending,
		},
	}
	for _, tt := range tests {
		out := FindTicTacToeWinner(tt.moves)
		if out != tt.out {
			t.Errorf("expected: %s, got: %s\n", tt.out, out)
		}
	}
}
