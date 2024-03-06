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

var BadComparisonNextBlocks = []Block{
	Equal{},
	Bigger{},
	Less{},
}

var BadSeparatorNextBlocks = []Block{
	Equal{},
	Bigger{},
	Less{},
	Value{},
	Plus{},
	Minus{},
	Multiplication{},
	Division{},
}

var BadIfNextBlocks = []Block{
	Equal{},
	Bigger{},
	Less{},
	Plus{},
	Minus{},
	Multiplication{},
	Division{},
}
