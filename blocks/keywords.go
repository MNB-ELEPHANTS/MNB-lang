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
	"put",
}

var CODES = map[string]string{
	"@=": "$$",
	"==": "##",
}
