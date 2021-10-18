package models

// BasicRomanNumber is ...
type BasicRomanNumber struct {
	I int
	V int
	X int
	L int
	C int
	D int
	M int
}

// RomanNumberAlias is ...
type RomanNumberAlias struct {
	Alias map[string]string
}

// Commodity is ...
type Commodity struct {
	Commodity map[string]string
}

// Init is ...
func (myRoman BasicRomanNumber) Init() BasicRomanNumber {
	myRoman.I = 1
	myRoman.V = 5
	myRoman.X = 10
	myRoman.L = 50
	myRoman.C = 100
	myRoman.D = 500
	myRoman.M = 1000
	return myRoman
}
