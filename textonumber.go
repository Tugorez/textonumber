package textonumber

import (
	"fmt"
)

func ToInt(posiblenum string) (num int, found bool) {
	if num, found = IsOrdinal(posiblenum); found {
	} else if num, found = IsRoman(posiblenum); found {
	}
	return
}

func ToString(posiblenum string) (num string, found bool) {
	val, found := ToInt(posiblenum)
	num = fmt.Sprint(val)
	return
}
