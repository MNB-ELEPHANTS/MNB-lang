package main

import (
	"fmt"
	"mnb/mnb-lang/lexer"
	"mnb/mnb-lang/parser"
	"mnb/mnb-lang/translator"
)

var testStr = `
a = 5
b = 'test string'
if a > b
`

func main() {
	p := parser.New()
	l := lexer.New(p)
	t := translator.New(l)
	fmt.Println(t.Translate(testStr))
}
