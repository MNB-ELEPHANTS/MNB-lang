package parser

func InBlocks(v string, l []string) bool {
	for _, b := range l {
		if b == v {
			return true
		}
	}
	return false
}
