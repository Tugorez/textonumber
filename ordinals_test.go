package textonumber

import (
	"testing"
)

var ordinalstonumber = []struct {
	in       string
	outval   int
	outfound bool
}{
	{"primero", 1, true}, {"quinto", 5, true}, {"décima", 10, true},
	{"primer", 1, true}, {"séptima", 7, true}, {"SépTiMo", 7, true},
	{"pelana", 0, false}, {"noesOrdinal", 0, false}, {"owo", 0, false},
}

func TestOrdinals(t *testing.T) {
	for _, test := range ordinalstonumber {
		if got, found := IsOrdinal(test.in); got != test.outval ||
			found != test.outfound {
			t.Errorf("Ordinal=%s\nExpected: %d %t\nGot: %d %t",
				test.in, test.outval, test.outfound, got, found)
		}
	}
}
