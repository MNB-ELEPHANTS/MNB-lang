package blocks

type Block interface {
	GetBlockValue() string
	CheckNextBlock(Block) bool
}
