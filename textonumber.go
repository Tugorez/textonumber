package textonumber

import (
	"fmt"
)

func ToInt(posiblenum string) (num int, err error) {
	var found bool
	if num, found = IsOrdinal(posiblenum); found == true {
	} else if num, found = IsRoman(posiblenum); found == true {

	} else {
		err = fmt.Errorf("%s is not a number", posiblenum)
	}
	return
}

func ToString(posiblenum string) (num string, err error) {
	val, err := ToInt(posiblenum)
	num = fmt.Sprint(val)
	return
}
