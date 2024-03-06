package blocks

type OpenBracket struct {
	BlockValue string
}

func NewOpenBracket(block string) OpenBracket {
	return OpenBracket{
		BlockValue: block,
	}
}

func (p OpenBracket) GetBlockValue() string {
	return p.BlockValue
}

func (p OpenBracket) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadMathOpNextBlocks)
}
