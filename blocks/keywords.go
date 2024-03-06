package blocks

var KEYWORDS = []string{
	// Math operators
	"+",
	"-",
	"*",
	"/",
	"=",
	">",
	"<",

	// Variable assignment
	"$$",
	"##",
	"=",

	// Separator
	";",

	// If
	"if",

	// Brackets
	"{",
	"}",
	"(",
	")",

	// Base functions
	"slon",
}

var CODES = map[string]string{
	"@=": "$$",
	"==": "##",
}
