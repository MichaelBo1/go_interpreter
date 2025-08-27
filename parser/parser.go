package parser

import (
	"github.com/MichaelBo1/go_interpreter/ast"
	"github.com/MichaelBo1/go_interpreter/lexer"
	"github.com/MichaelBo1/go_interpreter/token"
)

type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {

}
