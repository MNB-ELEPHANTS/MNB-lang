package blocks

type Minus struct {
	BlockValue string
}

func NewMinus(block string) Minus {
	return Minus{
		BlockValue: block,
	}
}

func (p Minus) GetBlockValue() string {
	return p.BlockValue
}

func (p Minus) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadMathOpNextBlocks)
}
