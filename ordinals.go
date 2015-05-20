package textonumber

import (
	"strings"
)

//If you want to add a ordinal(in other language) number just put the
//name as a key , and the val as the map-val
var Ordinals = map[string]int{
	"único":           0,
	"unico":           0,
	"única":           0,
	"unica":           0,
	"primero":         1,
	"primer":          1,
	"primera":         1,
	"segundo":         2,
	"segunda":         2,
	"tercero":         3,
	"tercer":          3,
	"tercera":         3,
	"cuarto":          4,
	"cuarta":          4,
	"quinto":          5,
	"quinta":          5,
	"sexto":           6,
	"sexta":           6,
	"séptimo":         7,
	"septimo":         7,
	"séptima":         7,
	"septima":         7,
	"octavo":          8,
	"octava":          8,
	"noveno":          9,
	"novena":          9,
	"nono":            9,
	"nona":            9,
	"décimo":          10,
	"decimo":          10,
	"décima":          10,
	"decima":          10,
	"undécimo":        11,
	"undecimo":        11,
	"décimoprimero":   11,
	"decimoprimero":   11,
	"duodécimo":       12,
	"duodecimo":       12,
	"décimosegundo":   12,
	"decimosegundo":   12,
	"vigésimo":        20,
	"vigesimo":        20,
	"vigésima":        20,
	"vigesima":        20,
	"trigésimo":       30,
	"trigesimo":       30,
	"trigésima":       30,
	"trigesima":       30,
	"cuadragésimo":    40,
	"cuadragesimo":    40,
	"cuadragésima":    40,
	"cuadragesima":    40,
	"quincuagésimo":   50,
	"quincuagésima":   50,
	"sexagésimo":      60,
	"sexagésima":      60,
	"septuagésimo":    70,
	"septuagésima":    70,
	"octogésimo":      80,
	"octogésima":      80,
	"nonagésimo":      90,
	"nonagésima":      90,
	"centésimo":       100,
	"centésima":       100,
	"ducentésimo":     200,
	"ducentésima":     200,
	"tricentésimo":    300,
	"tricentésima":    300,
	"cuadrigentésimo": 400,
	"cuadrigentésima": 400,
	"quingentésimo":   500,
	"quingentésima":   500,
	"sexcentésimo":    600,
	"sexcentésima":    600,
	"septingentésimo": 700,
	"septingentésima": 700,
	"octingentésimo":  800,
	"octingentésima":  800,
	"noningentésimo":  900,
	"noningentésima":  900,
	"milésimo":        1000,
	"milésima":        1000,
}

//IsOrdinal receives the ordinal number as string and
//returns the "number" and set "is" to true if it was a valid ordinal
//set "is" to false otherwise.
func IsOrdinal(word string) (ordinal int, is bool) {
	word = strings.ToLower(word)
	ordinal, is = Ordinals[word]
	return
}
