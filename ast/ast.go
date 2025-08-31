package ast

import (
	"bytes"

	"github.com/MichaelBo1/go_interpreter/token"
)

type Node interface {
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

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var result bytes.Buffer

	for _, s := range p.Statements {
		result.WriteString(s.String())
	}
	return result.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var result bytes.Buffer

	result.WriteString(ls.TokenLiteral() + " ")
	result.WriteString(ls.Name.String())
	result.WriteString(" = ")

	if ls.Value != nil {
		result.WriteString(ls.Value.String())
	}
	result.WriteString(";")

	return result.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var result bytes.Buffer

	result.WriteString(rs.TokenLiteral() + " ")
	if rs.Value != nil {
		result.WriteString(rs.Value.String())
	}
	result.WriteString(";")

	return result.String()
}

// Scripting languages often differentiate an expression statement as for example here `x + 10;` is valid code; we can have a line consisting solely of an expression;
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
