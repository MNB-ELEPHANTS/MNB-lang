package blocks

type If struct {
	BlockValue string
}

func NewIf(block string) If {
	return If{
		BlockValue: block,
	}
}

func (p If) GetBlockValue() string {
	return p.BlockValue
}

func (p If) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadIfNextBlocks)
}
