package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var tests = []struct {
		in  int
		out []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tt := range tests {
		res := Fizzbuzz(tt.in)
		if !reflect.DeepEqual(tt.out, res) {
			t.Errorf("expected: %s, got: %s", tt.out, res)
		}
	}
}
