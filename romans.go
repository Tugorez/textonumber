package textonumber

var Romans = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

//fromRoman converts a valid roman number into a decimal one.
func fromRoman(roman string) (arabic int) {
	last_digit := 1000
	for _, r := range roman {
		digit := Romans[r]
		if last_digit < digit {
			arabic -= 2 * last_digit
		}
		last_digit = digit
		arabic += digit
	}
	return arabic
}

//IsRoman receives the roman number as string and
//returns the "number" and set "is" to true if it was a valid roman
//set "is" to false otherwise.
func IsRoman(word string) (roman int, is bool) {
	for _, l := range word {
		if _, is = Romans[l]; !is {
			return
		}
	}
	roman = fromRoman(word)
	return
}
