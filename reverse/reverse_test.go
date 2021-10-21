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
		res2 := Reverse2(tt.in)
		if res2 != tt.out {
			t.Errorf("reverse2 failed: got %s", res2)
		}
	}
}
