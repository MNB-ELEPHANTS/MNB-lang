package lexer

import (
	"unicode"
)

type Lexer struct{}

func New() *Lexer {
	return &Lexer{}
}

func (l *Lexer) Lex(code string) []Token {
	var tokens []Token

	cursor := 0

	for cursor < len(code) {
		// Code separator is next-line sign
		if string(code[cursor]) == "\n" {
			tokens = append(tokens, Separator{Value: "\n"})
		} else if string(code[cursor]) == "." {
			tokens = append(tokens, Dot{Value: "."})
		} else if string(code[cursor]) == "," {
			tokens = append(tokens, Comma{Value: ","})

			// Math operations
		} else if string(code[cursor]) == "+" {
			tokens = append(tokens, Plus{Value: "+"})
		} else if string(code[cursor]) == "-" {
			tokens = append(tokens, Minus{Value: "-"})
		} else if string(code[cursor]) == "*" {
			tokens = append(tokens, Multiplication{Value: "*"})
		} else if string(code[cursor]) == "/" {
			tokens = append(tokens, Division{Value: "/"})
		} else if string(code[cursor]) == "=" {
			tokens = append(tokens, Equal{Value: "="})
		} else if string(code[cursor]) == "@" &&
			(cursor < len(code)-1 && string(code[cursor+1]) == "=") {
			tokens = append(tokens, Assignment{Value: "@="})
			cursor++

			// Brackets
		} else if string(code[cursor]) == "(" {
			tokens = append(tokens, OpenBracket{Value: "("})
		} else if string(code[cursor]) == ")" {
			tokens = append(tokens, CloseBracket{Value: ")"})
		} else if string(code[cursor]) == "{" {
			tokens = append(tokens, OpenFigureBracket{Value: "{"})
		} else if string(code[cursor]) == "}" {
			tokens = append(tokens, CloseFigureBracket{Value: "}"})

			// Comprasion tokens
		} else if string(code[cursor]) == ">" {
			tokens = append(tokens, Bigger{Value: ">"})
		} else if string(code[cursor]) == "<" {
			tokens = append(tokens, Less{Value: "<"})
		} else if string(code[cursor]) == "=" &&
			string(code[cursor+1]) == "=" {
			tokens = append(tokens, IsEqual{Value: "=="})
			cursor++
		} else if string(code[cursor]) == "!" &&
			string(code[cursor+1]) == "=" {
			tokens = append(tokens, IsNotEqual{Value: "!="})
			cursor++
		} else if unicode.IsLetter(rune(code[cursor])) {
			word := ""
			for unicode.IsLetter(rune(code[cursor])) {
				word = word + string(code[cursor])
				cursor++
			}
			cursor--

			switch word {
			case "if":
				tokens = append(tokens, If{"if"})
			case "else":
				tokens = append(tokens, Else{"else"})
			case "slon":
				tokens = append(tokens, Put{Value: "slon"})
			case "fn":
				tokens = append(tokens, Function{Value: "fn"})
			case "return":
				tokens = append(tokens, Return{Value: "fn"})
			default:
				tokens = append(tokens, Variable{Value: word})
			}

			// Variables
		} else if string(code[cursor]) == "'" {
			value := ""
			cursor++
			for string(code[cursor]) != "'" {
				value = value + string(code[cursor])
				cursor++
			}
			tokens = append(tokens, Value{Value: "\"" + value + "\""})
		} else if unicode.IsDigit(rune(code[cursor])) {
			num := ""
			for unicode.IsDigit(rune(code[cursor])) {
				num = num + string(code[cursor])
				cursor++
			}
			cursor--
			tokens = append(tokens, Value{Value: num})
		}

		cursor++
	}

	return tokens
}
