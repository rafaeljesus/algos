package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		capacity, out int
	}{
		{7, 22},
		{6, 17},
	}
	profits := []int{1, 6, 10, 16}
	weights := []int{1, 2, 3, 5}
	for _, tt := range tests {
		out := Knapsack(profits, weights, tt.capacity)
		if out != tt.out {
			t.Errorf("expected: %d, got: %d", tt.out, out)
		}
	}
}
