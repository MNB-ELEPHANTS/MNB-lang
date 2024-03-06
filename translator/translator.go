package translator

import (
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
	input := t.Lexer.Parser.Parse(code)

	result := START

	for _, b := range input {
		switch reflect.TypeOf(b) {

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

		case reflect.TypeOf(blocks.Equal{}):
			result = result + "="

		case reflect.TypeOf(blocks.Bigger{}):
			result = result + ">"

		case reflect.TypeOf(blocks.Less{}):
			result = result + "<"

		case reflect.TypeOf(blocks.If{}):
			result = result + "if "

		case reflect.TypeOf(blocks.Value{}):
			result = result + b.GetBlockValue()

		case reflect.TypeOf(blocks.Variable{}):
			result = result + b.GetBlockValue()

		default:
			continue
		}
	}

	return result + END
}
