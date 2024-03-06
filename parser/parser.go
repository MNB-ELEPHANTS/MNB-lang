package parser

import (
	"mnb/mnb-lang/blocks"
	"strings"
)

type Parser struct {
	KeyWords []string
}

func New() *Parser {
	return &Parser{
		KeyWords: blocks.KEYWORDS,
	}
}

func (p *Parser) PreProcess(code string) []string {
	preCode := strings.Split(code, "'")

	var parts []string

	// Split by quotes
	for i, part := range preCode {
		if i%2 != 0 {
			parts = append(parts, "'"+part+"'")
		} else {
			parts = append(parts, part)
		}
	}

	var parts2 []string

	replaceList := map[string]string{
		"\n": " ; ",
	}

	for _, kv := range p.KeyWords {
		replaceList[kv] = " " + kv + " "
	}

	for i, part := range parts {
		if i%2 == 0 {
			for k, v := range blocks.CODES {
				part = strings.ReplaceAll(part, k, v)
			}

			for k, v := range replaceList {
				part = strings.ReplaceAll(part, k, v)
			}
		}
		parts2 = append(parts2, part)
	}

	var end []string

	for i, part := range parts2 {
		if i%2 == 0 {
			part2 := strings.Split(part, " ")
			for _, part3 := range part2 {
				if part3 != "" {
					end = append(end, part3)
				}
			}
		}
		end = append(end, part)
	}

	return end
}

func (p *Parser) Parse(code string) []blocks.Block {
	preBlocks := p.PreProcess(code)
	var endBlocks []blocks.Block

	for _, b := range preBlocks {
		if b == "" {
			continue
		}

		// Base functions
		if b == "put" {
			endBlocks = append(endBlocks, blocks.NewPut("put"))

			// Math operations
		} else if b == "+" {
			endBlocks = append(endBlocks, blocks.NewPlus("+"))
		} else if b == "-" {
			endBlocks = append(endBlocks, blocks.NewMinus("-"))
		} else if b == "*" {
			endBlocks = append(endBlocks, blocks.NewMultiplication("*"))
		} else if b == "/" {
			endBlocks = append(endBlocks, blocks.NewDivision("/"))

			// Comprasion
		} else if b == ">" {
			endBlocks = append(endBlocks, blocks.NewBigger(">"))
		} else if b == "<" {
			endBlocks = append(endBlocks, blocks.NewLess("<"))
		} else if b == blocks.CODES["=="] {
			endBlocks = append(endBlocks, blocks.NewIsEqual("=="))

			// Assignment and Equal
		} else if b == blocks.CODES["@="] {
			endBlocks = append(endBlocks, blocks.NewAssignment("@="))
		} else if b == "=" {
			endBlocks = append(endBlocks, blocks.NewEqual("="))

			// Separator
		} else if b == ";" {
			endBlocks = append(endBlocks, blocks.NewSeparator(b))

			// If
		} else if b == "if" {
			endBlocks = append(endBlocks, blocks.NewIf(b))

			// Variables and Values
		} else if p.IsValue(b) {
			endBlocks = append(endBlocks, blocks.NewValue(b))
		} else if p.IsVariable(b) {
			endBlocks = append(endBlocks, blocks.NewVariable(b))

			// Brackets
		} else if b == "{" {
			endBlocks = append(endBlocks, blocks.NewOpenFigureBracket(b))
		} else if b == "}" {
			endBlocks = append(endBlocks, blocks.NewCloseFigureBracket(b))
		} else if b == "(" {
			endBlocks = append(endBlocks, blocks.NewOpenBracket(b))
		} else if b == ")" {
			endBlocks = append(endBlocks, blocks.NewCloseBracket(b))
		}

	}

	return endBlocks
}
