package blocks

type CloseBracket struct {
	BlockValue string
}

func NewCloseBracket(block string) CloseBracket {
	return CloseBracket{
		BlockValue: block,
	}
}

func (p CloseBracket) GetBlockValue() string {
	return p.BlockValue
}

func (p CloseBracket) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadMathOpNextBlocks)
}
