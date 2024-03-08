// This file contains tokens for lexer
// Every token struct has a value
package lexer

// Assignment token
type Assignment struct{ Value string }

func (a Assignment) GetValue() string { return a.Value }

// Equal token
type Equal struct{ Value string }

func (e Equal) GetValue() string { return e.Value }

// "If" expression token
type If struct{ Value string }

func (i If) GetValue() string { return i.Value }

// "Else" expression token
type Else struct{ Value string }

func (e Else) GetValue() string { return e.Value }

// "If equal" expression token
type IsEqual struct{ Value string }

func (ie IsEqual) GetValue() string { return ie.Value }

// "If not equal" expression token
type IsNotEqual struct{ Value string }

func (ie IsNotEqual) GetValue() string { return ie.Value }

// Bigger sign token
type Bigger struct{ Value string }

func (b Bigger) GetValue() string { return b.Value }

// Less sign token
type Less struct{ Value string }

func (l Less) GetValue() string { return l.Value }

// Open bracket token
type OpenBracket struct{ Value string }

func (ob OpenBracket) GetValue() string { return ob.Value }

// Close bracket token
type CloseBracket struct{ Value string }

func (cb CloseBracket) GetValue() string { return cb.Value }

// Open figure bracket token
type OpenFigureBracket struct{ Value string }

func (of OpenFigureBracket) GetValue() string { return of.Value }

// Close figure bracket token
type CloseFigureBracket struct{ Value string }

func (cf CloseFigureBracket) GetValue() string { return cf.Value }

// Plus sign token
type Plus struct{ Value string }

func (p Plus) GetValue() string { return p.Value }

// Minus sign token
type Minus struct{ Value string }

func (m Minus) GetValue() string { return m.Value }

// Multiplication sign token
type Multiplication struct{ Value string }

func (m Multiplication) GetValue() string { return m.Value }

// Division sign token
type Division struct{ Value string }

func (d Division) GetValue() string { return d.Value }

// Put function token
type Put struct{ Value string }

func (p Put) GetValue() string { return p.Value }

// Separator sign token
type Separator struct{ Value string }

func (s Separator) GetValue() string { return s.Value }

// Value token
type Value struct{ Value string }

func (v Value) GetValue() string { return v.Value }

// Variable token
type Variable struct{ Value string }

func (v Variable) GetValue() string { return v.Value }
