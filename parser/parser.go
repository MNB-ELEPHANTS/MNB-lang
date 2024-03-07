package parser

import (
	"mnb/mnb-lang/lexer"
)

type Parser struct {
	Lexer *lexer.Lexer
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{
		Lexer: l,
	}
}

func (p *Parser) Parse(code string) ([]lexer.Token, error) {
	input := p.Lexer.Lex(code)

	return input, nil
}
