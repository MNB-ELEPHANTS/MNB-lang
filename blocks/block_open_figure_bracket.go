package blocks

type OpenFigureBracket struct {
	BlockValue string
}

func NewOpenFigureBracket(block string) OpenFigureBracket {
	return OpenFigureBracket{
		BlockValue: block,
	}
}

func (p OpenFigureBracket) GetBlockValue() string {
	return p.BlockValue
}

func (p OpenFigureBracket) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadMathOpNextBlocks)
}
