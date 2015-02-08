package textonumber

import (
	"testing"
)

var outstring = []struct {
	in       string
	outval   string
	outfound bool
}{
	{"primero", "1", true}, {"II", "2", true},
	{"décimo", "10", true}, {"LXXXIII", "83", true},
	{"iii", "0", false}, {"i", "0", false},
}

var outint = []struct {
	in       string
	outval   int
	outfound bool
}{
	{"primero", 1, true}, {"II", 2, true},
	{"décimo", 10, true}, {"LXXXIII", 83, true},
	{"iii", 0, false}, {"i", 0, false},
}

func TestTextoText(t *testing.T) {
	for _, test := range outstring {
		if got, found := ToString(test.in); got != test.outval ||
			found != test.outfound {
			t.Errorf("TextNumber:%s\nExpected: %s %t\nGot: %s %t",
				test.in, test.outval, test.outfound, got, found)
		}
	}
}

func TestTextoInt(t *testing.T) {
	for _, test := range outint {
		if got, found := ToInt(test.in); got != test.outval ||
			found != test.outfound {
			t.Errorf("TextNumber:%s\nExpected: %d %t\nGot: %d %t",
				test.in, test.outval, test.outfound, got, found)
		}
	}
}
