package blocks

type Value struct {
	BlockValue string
}

func NewValue(block string) Value {
	return Value{
		BlockValue: block,
	}
}

func (p Value) GetBlockValue() string {
	return p.BlockValue
}

func (p Value) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadValueNextBlocks)
}
