package tictactoe

import "testing"

func TestTicTacToe(t *testing.T) {
	X := 1
	O := 2
	n := 3
	tt := NewTicTacToe(n)
	_ = tt.Move(0, 0, X)
	_ = tt.Move(0, 2, O)
	_ = tt.Move(2, 2, X)
	_ = tt.Move(1, 1, O)
	_ = tt.Move(2, 0, X)
	_ = tt.Move(1, 0, O)
	winer := tt.Move(2, 1, X)
	if winer != X {
		t.Errorf("expected: %d, got: %d", X, winer)
	}
}

func TestTicTacToeCheckRow(t *testing.T) {
	X := 1
	n := 3
	tt := NewTicTacToe(n)
	_ = tt.Move(0, 0, X)
	_ = tt.Move(0, 1, X)
	winer := tt.Move(0, 2, X)
	if winer != X {
		t.Errorf("expected: %d, got: %d", X, winer)
	}
}

func TestTicTacToeCheckColumn(t *testing.T) {
	X := 1
	n := 3
	tt := NewTicTacToe(n)
	_ = tt.Move(0, 0, X)
	_ = tt.Move(1, 0, X)
	winer := tt.Move(2, 0, X)
	if winer != X {
		t.Errorf("expected: %d, got: %d", X, winer)
	}
}

func TestTicTacToeCheckDialog(t *testing.T) {
	X := 1
	n := 3
	tt := NewTicTacToe(n)
	_ = tt.Move(0, 0, X)
	_ = tt.Move(1, 1, X)
	winer := tt.Move(2, 2, X)
	if winer != X {
		t.Errorf("expected: %d, got: %d", X, winer)
	}
}

func TestTicTacToeCheckAntiDialog(t *testing.T) {
	X := 1
	n := 3
	tt := NewTicTacToe(n)
	_ = tt.Move(0, 2, X)
	_ = tt.Move(1, 1, X)
	winer := tt.Move(2, 0, X)
	if winer != X {
		t.Errorf("expected: %d, got: %d", X, winer)
	}
}
