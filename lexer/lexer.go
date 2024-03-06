package lexer

import (
	"mnb/mnb-lang/parser"
)

type Lexer struct {
	Parser *parser.Parser
}

func New(p *parser.Parser) *Lexer {
	return &Lexer{
		Parser: p,
	}
}

func (l *Lexer) Lex(code string) error {
	input := l.Parser.Parse(code)

	for i := 0; i < len(input)-1; i++ {
		if !input[i].CheckNextBlock(input[i+1]) {
			return BadBlockSequenseErr
		}
	}

	return nil
}
