package translator

import (
	"fmt"
	"log"
	"mnb/mnb-lang/lexer"
	"mnb/mnb-lang/parser"
	"reflect"
)

type Translator struct {
	Parser *parser.Parser
}

func New(p *parser.Parser) *Translator {
	return &Translator{
		Parser: p,
	}
}

func (t *Translator) Translate(code string) string {
	input, err := t.Parser.Parse(code)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range input {
		fmt.Println(reflect.TypeOf(v), v.GetValue())
	}

	result := START

	for _, b := range input {
		switch reflect.TypeOf(b) {

		case reflect.TypeOf(lexer.Separator{}):
			result = result + "\n"
		case reflect.TypeOf(lexer.Dot{}):
			result = result + "."
		case reflect.TypeOf(lexer.Comma{}):
			result = result + ","
		case reflect.TypeOf(lexer.DoubleDot{}):
			result = result + ":"

		// Math operations
		case reflect.TypeOf(lexer.Minus{}):
			result = result + "-"
		case reflect.TypeOf(lexer.Plus{}):
			result = result + "+"
		case reflect.TypeOf(lexer.Multiplication{}):
			result = result + "*"
		case reflect.TypeOf(lexer.Division{}):
			result = result + "/"

		// Assignment and Equal
		case reflect.TypeOf(lexer.Assignment{}):
			result = result + ":="
		case reflect.TypeOf(lexer.Equal{}):
			result = result + "="

		// Comprasion
		case reflect.TypeOf(lexer.Bigger{}):
			result = result + ">"
		case reflect.TypeOf(lexer.Less{}):
			result = result + "<"
		case reflect.TypeOf(lexer.IsEqual{}):
			result = result + "=="
		case reflect.TypeOf(lexer.IsNotEqual{}):
			result = result + "!="

		// If
		case reflect.TypeOf(lexer.If{}):
			result = result + "if "
		// Else
		case reflect.TypeOf(lexer.Else{}):
			result = result + "else"
		// Function
		case reflect.TypeOf(lexer.Function{}):
			result = result + "func "
		case reflect.TypeOf(lexer.Return{}):
			result = result + "return "

		// Import
		case reflect.TypeOf(lexer.Import{}):
			result = result + "import "

		// Type
		case reflect.TypeOf(lexer.Type{}):
			result = result + "type "

		// Struct
		case reflect.TypeOf(lexer.Struct{}):
			result = result + "struct"

		// Variables and Values
		case reflect.TypeOf(lexer.Value{}):
			v := b.GetValue()
			if v == "yes" {
				result = result + "true"
			} else if v == "no" {
				result = result + "false"
			} else if string(v[0]) == "'" &&
				string(v[len(v)-1]) == "'" {
				result = result + "\"" + v[1:len(v)-1] + "\""
			} else {
				result = result + v
			}

		case reflect.TypeOf(lexer.Variable{}):
			result = result + b.GetValue() + " "

		// Brackets
		case reflect.TypeOf(lexer.OpenFigureBracket{}):
			result = result + "{"
		case reflect.TypeOf(lexer.CloseFigureBracket{}):
			result = result + "}"
		case reflect.TypeOf(lexer.OpenBracket{}):
			result = result + "("
		case reflect.TypeOf(lexer.CloseBracket{}):
			result = result + ")"

		// Base functions
		case reflect.TypeOf(lexer.Put{}):
			result = result + "println"

		// Types
		case reflect.TypeOf(lexer.String{}):
			result = result + "string"
		case reflect.TypeOf(lexer.Int{}):
			result = result + "int"
		case reflect.TypeOf(lexer.Uint{}):
			result = result + "uint"
		case reflect.TypeOf(lexer.Bool{}):
			result = result + "bool"

		default:
			continue
		}
	}

	return result + END
}
