package object

import (
	"bytes"
	"fmt"
	"monkey/pkg/ast"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ    = "INTEGER"
	BOOLEAN_OBJ    = "BOOLEAN"
	NULL_OBJ       = "NULL"
	FUNC_OBJ       = "FUNC"
	RETURN_VAL_OBJ = "RETURN_VALUE"
	ERROR_OBJ      = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return "null"
}

type Func struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Env
}

func (f *Func) Type() ObjectType {
	return FUNC_OBJ
}

func (f *Func) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString((") {\n"))
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type ReturnVal struct {
	Value Object
}

func (r *ReturnVal) Type() ObjectType {
	return RETURN_VAL_OBJ
}

func (r *ReturnVal) Inspect() string {
	return r.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "error: " + e.Message
}
