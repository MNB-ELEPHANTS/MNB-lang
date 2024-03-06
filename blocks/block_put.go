package blocks

type Put struct {
	BlockValue string
}

func NewPut(block string) Put {
	return Put{
		BlockValue: block,
	}
}

func (p Put) GetBlockValue() string {
	return p.BlockValue
}

func (p Put) CheckNextBlock(next Block) bool {
	return !InBlocks(next, BadPutNextBlocks)
}
