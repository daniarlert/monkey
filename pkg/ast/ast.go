package ast

import (
	"bytes"
	"monkey/pkg/token"
	"strings"
)

// Node is an interface that every node in the AST has to implement.
type Node interface {
	// TokenLiteral returns the literal valur of the token it's associated with. Is
	// used for debugging and testing.
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (v *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(v.TokenLiteral() + " ")
	out.WriteString(v.Name.String())
	out.WriteString(" = ")

	if v.Value != nil {
		out.WriteString(v.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token // the 'return' keyword
	ReturnValue Expression
}

func (s *ReturnStatement) statementNode() {}

func (s *ReturnStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	if s.ReturnValue != nil {
		out.WriteString(s.ReturnValue.String())
	}

	out.WriteString(s.TokenLiteral() + ";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (s *StringLiteral) expressionNode() {}

func (s *StringLiteral) TokenLiteral() string {
	return s.Token.Literal
}

func (s *StringLiteral) String() string {
	return s.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token // the prefix token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // the operator token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

type IfExpression struct {
	Token       token.Token // the "if" token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (e *IfExpression) expressionNode() {}

func (e *IfExpression) TokenLiteral() string {
	return e.Token.Literal
}

func (e *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(e.Condition.String())
	out.WriteString(" ")
	out.WriteString(e.Consequence.String())

	if e.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(e.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (s *BlockStatement) statementNode() {}

func (s *BlockStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range s.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // the 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) expressionNode() {}

func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token // the '(' token
	Func      Expression
	Arguments []Expression
}

func (c *CallExpression) expressionNode() {}

func (c *CallExpression) TokenLiteral() string {
	return c.Token.Literal
}

func (c *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(c.Func.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type ArrayListeral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayListeral) expressionNode() {}

func (al *ArrayListeral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayListeral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()

}
