package blocks

type Equal struct {
	BlockValue string
}

func NewEqual(block string) Equal {
	return Equal{
		BlockValue: block,
	}
}

func (p Equal) GetBlockValue() string {
	return p.BlockValue
}

func (p Equal) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadEqualNextBlocks)
}
