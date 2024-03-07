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

		// Math operations
		case reflect.TypeOf(lexer.Separator{}):
			result = result + "\n"
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
			result = result + b.GetValue()

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

		default:
			continue
		}
	}

	return result + END
}
