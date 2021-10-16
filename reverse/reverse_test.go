package reverse

import "testing"

func TestReverse(t *testing.T) {
	var input = []struct {
		in, out string
	}{
		{"golang", "gnalog"},
		{"observability", "ytilibavresbo"},
	}
	for _, tt := range input {
		res := Reverse(tt.in)
		if res != tt.out {
			t.Errorf("reverse failed: got %s", res)
		}
	}
}
