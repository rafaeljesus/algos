package class_scheduling

import (
	"reflect"
	"testing"
)

func FindSchedule(schedule [][]float64) [][]float64 {
	res := make([][]float64, 0)
	for i, _ := range res {
		res[i] = make([]float64, 0)
	}
	start := 0
	res = append(res, schedule[start])
	for next := 1; next < len(schedule); next++ {
		if schedule[start][1] <= schedule[next][0] {
			res = append(res, schedule[next])
			start = next
		}
	}
	return res
}

func Test(t *testing.T) {
	tests := []struct {
		in, out [][]float64
	}{
		{
			[][]float64{
				[]float64{9, 10},      // art
				[]float64{9.3, 10.3},  // eng
				[]float64{10, 11},     // math
				[]float64{10.3, 11.3}, // cs
				[]float64{11, 12},     // music
			},
			[][]float64{
				[]float64{9, 10},
				[]float64{10, 11},
				[]float64{11, 12},
			},
		},
	}
	for _, tt := range tests {
		res := FindSchedule(tt.in)
		if !reflect.DeepEqual(res, tt.out) {
			t.Errorf("expected: %v, got %v", tt.out, res)
		}
	}
}
