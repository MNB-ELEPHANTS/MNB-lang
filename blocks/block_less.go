package blocks

type Less struct {
	BlockValue string
}

func NewLess(block string) Less {
	return Less{
		BlockValue: block,
	}
}

func (p Less) GetBlockValue() string {
	return p.BlockValue
}

func (p Less) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadComparisonNextBlocks)
}
