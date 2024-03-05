package blocks

type Multiplication struct {
	BlockValue string
}

func NewMultiplication(block string) Multiplication {
	return Multiplication{
		BlockValue: block,
	}
}

func (p Multiplication) GetBlockValue() string {
	return p.BlockValue
}

func (p Multiplication) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadMathOpNextBlocks)
}
