package search_in_sorted_array

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		in       [][]int
		target   int
		expected []int
	}{
		{
			[][]int{
				[]int{1, 4, 7, 12, 15, 1000},
				[]int{2, 5, 19, 31, 32, 1001},
				[]int{3, 8, 24, 33, 35, 100},
				[]int{40, 41, 42, 44, 45, 100},
				[]int{99, 100, 103, 106, 128, 100},
			},
			44,
			[]int{3, 3},
		},
		{
			[][]int{
				[]int{1, 4, 7, 12, 15, 1000},
				[]int{2, 5, 19, 31, 32, 1001},
				[]int{3, 8, 24, 33, 35, 100},
				[]int{40, 41, 42, 44, 45, 100},
				[]int{99, 100, 103, 106, 128, 100},
			},
			1,
			[]int{0, 0},
		},
	}
	for _, tt := range tests {
		res := SearchInSortedMatrix(tt.in, tt.target)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("expected: %v, got: %v", tt.expected, res)
		}
	}
}

// time O(nm) | space O(1)
func SearchInSortedMatrix(matrix [][]int, target int) []int {
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return []int{row, col}
		}
		if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return []int{-1, -1}
}

// row = 0
// col = 5
// 0 < 5 && 5 > 0? y
// [1, 4, 7, 12, 15, _1000_]
// 1000 == 44? n
// 1000 < 44? n
// else 1000 > 4 y
// skip all col[5] and move pointer to the left
// matrix[row][col-1]
// col = 4
//[1, 4, 7, 12, *15, -1000],
//[2, 5, 19, 31, 32, -1001],
//[3, 8, 24, 33, 35, -1002],
//[40, 41, 42, 44, 45, -1003],
//[99, 100, 103, 106, 128, -1004]

// row = 0
// col = 4
// 0 < 5 && col > 4? y
// 15 == 44? n
// 15 < 44? y
// skip all from the left and move one position down
// matrix[row + 1][col]
// row = 1
//[-1, -4, -7, -12, -15, -1000],
//[2, 5, 19, 31, *32, -1001],
//[3, 8, 24, 33, 35, -1002],
//[40, 41, 42, 44, 45, -1003],
//[99, 100, 103, 106, 128, -1004]

// row = 1
// col = 4
// row < len(matrix) && col > 0
// 1 < % && 4 > 0? y
// 32 == 44? n
// 32 < 44? y
// skip all from the left, move one down
// matrix[row + 1][col]
// row = 2
//[-1, -4, -7, -12, -15, -1000],
//[-2, -5, -19, -31, -32, -1001],
//[3, 8, 24, 33, *35, -1002],
//[40, 41, 42, 44, 45, -1003],
//[99, 100, 103, 106, 128, -1004]

// row = 2
// col = 4
// row < len(matrix) && col > 0
// 2 < 5 && 4 > 0? y
// 35 < 44? y
// discard all from the left, move one down
// matrix[row + 1][col]
// row = 3
//[-1, -4, -7, -12, -15, -1000],
//[-2, -5, -19, -31, -32, -1001],
//[-3, -8, -24, -33, -35, -1002],
//[40, 41, 42, 44, *45, -1003],
//[99, 100, 103, 106, 128, -1004]

// row = 3
// col 4
// row < len(matrix) && col > 0
// 3 < 5 && 4 > 0? y
// 45 == 44? n
// 45 < 44? n
// else 45 > 44
// move one to the left, discard the whole col
// matrix[row][col-1]
// col = 3
//[-1, -4, -7, -12, -15, -1000],
//[-2, -5, -19, -31, -32, -1001],
//[-3, -8, -24, -33, -35, -1002],
//[40, 41, 42, *44, -45, -1003],
//[99, 100, 103, 106, -128, -1004]

// row = 3
// col = 3
// 3 < 5 && 3 > 0? y
// 44 == 44? y
// return [3, 3]
