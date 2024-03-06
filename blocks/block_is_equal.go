package blocks

type IsEqual struct {
	BlockValue string
}

func NewIsEqual(block string) IsEqual {
	return IsEqual{
		BlockValue: block,
	}
}

func (p IsEqual) GetBlockValue() string {
	return p.BlockValue
}

func (p IsEqual) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadComparisonNextBlocks)
}
