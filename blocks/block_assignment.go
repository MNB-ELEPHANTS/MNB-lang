package blocks

type Assignment struct {
	BlockValue string
}

func NewAssignment(block string) Assignment {
	return Assignment{
		BlockValue: block,
	}
}

func (a Assignment) GetBlockValue() string {
	return a.BlockValue
}

func (a Assignment) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadAssignmentNextBlocks)
}
