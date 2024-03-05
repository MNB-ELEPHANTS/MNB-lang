package blocks

var BadMathOpNextBlocks = []Block{
	Plus{},
	Minus{},
	Multiplication{},
	Division{},
}

var BadVariableNextBlocks = []Block{
	Variable{},
	Value{},
}

var BadValueNextBlocks = []Block{
	Variable{},
	Value{},
}
