package blocks

type Separator struct {
	BlockValue string
}

func NewSeparator(block string) Separator {
	return Separator{
		BlockValue: block,
	}
}

func (p Separator) GetBlockValue() string {
	return p.BlockValue
}

func (p Separator) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadSeparatorNextBlocks)
}
