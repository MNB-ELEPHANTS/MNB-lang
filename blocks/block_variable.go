package blocks

type Variable struct {
	BlockValue string
}

func NewVariable(block string) Variable {
	return Variable{
		BlockValue: block,
	}
}

func (p Variable) GetBlockValue() string {
	return p.BlockValue
}

func (p Variable) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadVariableNextBlocks)
}
