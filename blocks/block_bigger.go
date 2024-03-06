package blocks

type Bigger struct {
	BlockValue string
}

func NewBigger(block string) Bigger {
	return Bigger{
		BlockValue: block,
	}
}

func (p Bigger) GetBlockValue() string {
	return p.BlockValue
}

func (p Bigger) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadComparisonNextBlocks)
}
