# textonumber
textonumber convert ordinal or roman numbers into their decimal representation

#How to
```go
package main
import (
	"fmt"
	"github.com/tugorez/textonumber"
	"log"
)

func main() {
	roman := "III"
	number, err := textonumber.ToInt(roman)
	if err != nil {
		log.Fatal(err)
	}
	snumber, err := textonumber.ToString(roman)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The dec representation of %s is %d(as int) \"%s\"(as string)\n",
		roman, number, snumber)
}
```
