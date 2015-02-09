package textonumber

import (
	"testing"
)

var outstring = []struct {
	in     string
	outval string
}{
	{"primero", "1"}, {"II", "2"},
	{"décimo", "10"}, {"LXXXIII", "83"},
	{"III", "3"}, {"VII", "7"},
}

var outint = []struct {
	in     string
	outval int
}{
	{"primero", 1}, {"II", 2},
	{"décimo", 10}, {"LXXXIII", 83},
	{"III", 3}, {"I", 1},
}

func TestTextoText(t *testing.T) {
	for _, test := range outstring {
		if got, err := ToString(test.in); got != test.outval ||
			err != nil {
			t.Errorf("TextNumber:%s\nExpected: %s , Got: %s",
				test.in, test.outval, got)
		}
	}
}

func TestTextoInt(t *testing.T) {
	for _, test := range outint {
		if got, err := ToInt(test.in); got != test.outval ||
			err != nil {
			t.Errorf("TextNumber:%s\nExpected: %d , Got: %d",
				test.in, test.outval, got)
		}
	}
}
