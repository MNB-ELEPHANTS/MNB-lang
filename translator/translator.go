package translator

import (
	"log"
	"mnb/mnb-lang/blocks"
	"mnb/mnb-lang/lexer"
	"reflect"
)

type Translator struct {
	Lexer *lexer.Lexer
}

func New(l *lexer.Lexer) *Translator {
	return &Translator{
		Lexer: l,
	}
}

func (t *Translator) Translate(code string) string {
	input, err := t.Lexer.Lex(code)

	if err != nil {
		log.Fatal(err)
	}

	result := START

	for _, b := range input {
		switch reflect.TypeOf(b) {

		// Math operations
		case reflect.TypeOf(blocks.Separator{}):
			result = result + "\n"

		case reflect.TypeOf(blocks.Minus{}):
			result = result + "-"

		case reflect.TypeOf(blocks.Plus{}):
			result = result + "+"

		case reflect.TypeOf(blocks.Multiplication{}):
			result = result + "*"

		case reflect.TypeOf(blocks.Division{}):
			result = result + "/"

		// Assignment and Equal
		case reflect.TypeOf(blocks.Assignment{}):
			result = result + ":="

		case reflect.TypeOf(blocks.Equal{}):
			result = result + "="

		// Comprasion
		case reflect.TypeOf(blocks.Bigger{}):
			result = result + ">"

		case reflect.TypeOf(blocks.Less{}):
			result = result + "<"

		case reflect.TypeOf(blocks.IsEqual{}):
			result = result + "=="

		// If
		case reflect.TypeOf(blocks.If{}):
			result = result + "if "

		// Variables and Values
		case reflect.TypeOf(blocks.Value{}):
			v := b.GetBlockValue()
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

		case reflect.TypeOf(blocks.Variable{}):
			result = result + b.GetBlockValue()

		// Brackets
		case reflect.TypeOf(blocks.OpenFigureBracket{}):
			result = result + "{"
		case reflect.TypeOf(blocks.CloseFigureBracket{}):
			result = result + "}"
		case reflect.TypeOf(blocks.OpenBracket{}):
			result = result + "("
		case reflect.TypeOf(blocks.CloseBracket{}):
			result = result + ")"

		// Base functions
		case reflect.TypeOf(blocks.Put{}):
			result = result + "println"

		default:
			continue
		}
	}

	return result + END
}
