package parser

import "unicode"

func (p *Parser) IsVariable(v string) bool {
	for _, r := range v {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
