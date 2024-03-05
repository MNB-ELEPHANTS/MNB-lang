package main

import (
	"fmt"
	"mnb/mnb-lang/parser"
	"reflect"
)

func main() {
	p := parser.New()
	blocks := p.Parse("fuck 12321 > * + - = yes")
	for _, b := range blocks {
		fmt.Println(reflect.TypeOf(b), b.GetBlockValue())
	}
}
