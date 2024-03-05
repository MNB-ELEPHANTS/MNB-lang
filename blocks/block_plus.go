package blocks

type Plus struct {
	BlockValue string
}

func NewPlus(block string) Plus {
	return Plus{
		BlockValue: block,
	}
}

func (p Plus) GetBlockValue() string {
	return p.BlockValue
}

func (p Plus) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadMathOpNextBlocks)
}
