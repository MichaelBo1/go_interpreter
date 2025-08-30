package parser

import (
	"fmt"

	"github.com/MichaelBo1/go_interpreter/ast"
	"github.com/MichaelBo1/go_interpreter/lexer"
	"github.com/MichaelBo1/go_interpreter/token"
)

type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{lex: lex}
	// We read two tokens ahead so curToken and peekToken are set by
	// lexing two tokens. If the input to the lexer is empty, we will
	// check and see that the curToken is a token.EOF and don't worry about the peekToken in that case.
	parser.NextToken()
	parser.NextToken()

	return parser
}

func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		program.Statements = append(program.Statements, stmt)
		p.NextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping expression parsing for now;
	for p.currentToken.Type != token.SEMICOLON {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.NextToken()

	// TODO: skipping expression parsing for now;
	for p.currentToken.Type != token.SEMICOLON {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) peekError(expectedType token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s.", expectedType, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(expectedType token.TokenType) bool {
	if p.peekToken.Type == expectedType {
		p.NextToken()
		return true
	}
	p.peekError(expectedType)
	return false
}
