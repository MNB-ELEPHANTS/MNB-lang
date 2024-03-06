package parser

import "unicode"

func (p *Parser) IsValue(v string) bool {
	return v == "yes" || v == "no" ||
		IsDigit(v) ||
		string(v[0]) == "'" ||
		string(v[len(v)-1]) == "'"
}

func IsDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
