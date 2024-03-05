package blocks

import "reflect"

func InBlocks(block Block, blocks []Block) bool {
	for _, b := range blocks {
		if reflect.TypeOf(b) == reflect.TypeOf(block) {
			return true
		}
	}
	return false
}
