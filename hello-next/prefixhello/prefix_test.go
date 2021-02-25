package prefixhello

import "testing"

func TestPrefixHello (t *testing.T) {
	cases := []struct {
		input, desired string
	} {
		{"name", "Hello name"},
		{"Hello", "Hello Hello"},
		{"", "Hello "},
	}

	
	for _, c := range cases {
		out := PrefixHello(c.input)
		if(out != c.desired) {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.input, out, c.desired)
		}
	}
	
}