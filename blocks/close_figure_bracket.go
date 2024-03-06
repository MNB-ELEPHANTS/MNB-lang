package blocks

type CloseFigureBracket struct {
	BlockValue string
}

func NewCloseFigureBracket(block string) CloseFigureBracket {
	return CloseFigureBracket{
		BlockValue: block,
	}
}

func (p CloseFigureBracket) GetBlockValue() string {
	return p.BlockValue
}

func (p CloseFigureBracket) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadMathOpNextBlocks)
}
