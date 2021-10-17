package robot

import (
	"reflect"
	"testing"
)

func TestRobotHandleMove(t *testing.T) {
	var input = []struct {
		move string
		out  []int
	}{
		{"UDDLRL", []int{-1, -1}},
		{"UDDLLRUUUDUURUDDUULLDRRRR", []int{2, 3}},
	}
	for _, tt := range input {
		pos := RobotHandleMove(tt.move)
		if !reflect.DeepEqual(pos, tt.out) {
			t.Errorf("unexpected %v, got: %v", pos, tt.out)
		}
	}
}
