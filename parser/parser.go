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
	preCode := strings.Split(code, "\"")

	var parts []string

	for i, part := range preCode {
		if i%2 != 0 {
			parts = append(parts, "\""+part+"\"")
		} else {
			parts = append(parts, part)
		}
	}

	var parts2 []string

	replaceList := map[string]string{
		"\n": " ",
	}

	for _, kv := range p.KeyWords {
		replaceList[kv] = " " + kv + " "
	}

	for i, part := range parts {
		if i%2 == 0 {
			for k, v := range replaceList {
				part = strings.ReplaceAll(part, k, v)
			}
			parts2 = append(parts2, part)
		}
	}

	var end []string

	for _, part := range parts2 {
		part2 := strings.Split(part, " ")
		for _, part3 := range part2 {
			if part3 != "" {
				end = append(end, part3)
			}
		}
	}

	return end
}

func (p *Parser) Parse(code string) []blocks.Block {
	preBlocks := p.PreProcess(code)
	var endBlocks []blocks.Block

	for _, b := range preBlocks {
		if b == "+" {
			endBlocks = append(endBlocks, blocks.NewPlus("+"))
		} else if b == "-" {
			endBlocks = append(endBlocks, blocks.NewMinus("-"))
		} else if b == "*" {
			endBlocks = append(endBlocks, blocks.NewMultiplication("*"))
		} else if b == "/" {
			endBlocks = append(endBlocks, blocks.NewDivision("/"))
		} else if p.IsValue(b) {
			endBlocks = append(endBlocks, blocks.NewValue(b))
		} else if p.IsVariable(b) {
			endBlocks = append(endBlocks, blocks.NewVariable(b))
		}
	}

	return endBlocks
}
