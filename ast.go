package uniql

type Node interface {
	node()
}

type Expression interface {
	Node
	expr()
}

type StringLiteral struct {
	Value string
}

type NumberLiteral struct {
	Value float64
}

type BooleanLiteral struct {
	Value bool
}

type UndefinedLiteral struct {
}

type NullLiteral struct {
}

type Identifier struct {
	Name string
}

type BinaryExpression struct {
	Op    Token
	Left  Expression
	Right Expression
}

func (*StringLiteral) node()    {}
func (*NumberLiteral) node()    {}
func (*BooleanLiteral) node()   {}
func (*UndefinedLiteral) node() {}
func (*Identifier) node()       {}
func (*BinaryExpression) node() {}

func (*StringLiteral) expr()    {}
func (*NumberLiteral) expr()    {}
func (*BooleanLiteral) expr()   {}
func (*UndefinedLiteral) expr() {}
func (*NullLiteral) expr()      {}
func (*Identifier) expr()       {}
func (*BinaryExpression) expr() {}
