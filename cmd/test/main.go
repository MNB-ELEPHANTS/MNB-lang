package main

import (
	"fmt"
	"mnb/mnb-lang/parser"
	"reflect"
)

var testStr = `a = 5
b = 'test string'
a2 = a * 10
a < a2`

func main() {
	p := parser.New()
	blocks := p.Parse(testStr)
	for _, b := range blocks {
		fmt.Println(reflect.TypeOf(b), b.GetBlockValue())
	}
}
