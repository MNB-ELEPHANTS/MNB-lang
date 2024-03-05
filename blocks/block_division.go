package blocks

type Division struct {
	BlockValue string
}

func NewDivision(block string) Division {
	return Division{
		BlockValue: block,
	}
}

func (p Division) GetBlockValue() string {
	return p.BlockValue
}

func (p Division) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadMathOpNextBlocks)
}
