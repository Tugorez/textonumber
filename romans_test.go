package textonumber

import (
	"testing"
)

var romanstonumbers = []struct {
	in       string
	outval   int
	outfound bool
}{
	{"III", 3, true}, {"VI", 6, true}, {"XIX", 19, true},
	{"MCMXCIII", 1993, true}, {"CL", 150, true}, {"VII", 7, true},
	{"ABC", 0, false}, {"iii", 0, false}, {"Cl", 0, false},
}

func TestRomans(t *testing.T) {
	for _, test := range romanstonumbers {
		if num, found := IsRoman(test.in); num != test.outval ||
			found != test.outfound {
			t.Errorf("\nExpected: num=%d found=%t\nGot: num=%d found=%t", test.outval, test.outfound, num, found)
		}

	}
}
