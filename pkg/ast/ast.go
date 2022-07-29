package ast

import "monkey/pkg/token"

// Node is an interface that every node in the AST has to implement.
type Node interface {
	// TokenLiteral returns the literal valur of the token it's associated with. Its
	// used for debugging and testing.
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST the parser produces.
type Program struct {
	// Statements contain the series of statements that represent a valid Monkey
	// program.
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type VarStatement struct {
	Token token.Token // the token.VAR token
	Name  *Identifier // holds the identifier of the binding
	Value Expression  // holds the expression that produces the value.
}

func (v *VarStatement) statementNode() {}

func (v *VarStatement) TokenLiteral() string {
	return v.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token // the 'return' keyword
	ReturnValue Expression
}

func (s *ReturnStatement) statementNode() {}

func (s *ReturnStatement) TokenLiteral() string {
	return s.Token.Literal
}
