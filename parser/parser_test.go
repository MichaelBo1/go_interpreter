package parser

import (
	"testing"

	"github.com/MichaelBo1/go_interpreter/ast"
	"github.com/MichaelBo1/go_interpreter/lexer"
)

func TestParsesLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 1337;
	`

	lex := lexer.New(input)
	par := New(lex)

	program := par.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	checkParserErrors(t, par)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got %d statements", len(program.Statements))
	}

	tests := []struct {
		expectedIndentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, test.expectedIndentifier)
	}
}

func TestParsesReturnStatments(t *testing.T) {
	input := `
	return 5;
	return 1337;`

	lex := lexer.New(input)
	par := New(lex)

	program := par.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	checkParserErrors(t, par)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. Got %d statements", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt is not an *ast.ReturnStatement")
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', instead got %q", returnStmt.TokenLiteral())
		}
	}
}

func testLetStatement(t testing.TB, parsedStmt ast.Statement, expectedName string) {
	t.Helper()

	if parsedStmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not `let`. Got %q", parsedStmt.TokenLiteral())
	}

	letStmt, ok := parsedStmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("parsed statement not a *ast.LetStatement. Got %T", parsedStmt)
	}

	if letStmt.Name.Value != expectedName {
		t.Errorf("letStmt.Name.Value not `%s`. Got `%s`", expectedName, letStmt.Name.Value)
	}

	if letStmt.Name.TokenLiteral() != expectedName {
		t.Errorf("s.Name not `%s`. Got=%s", expectedName, letStmt.Name)
	}
}

func checkParserErrors(t testing.TB, p *Parser) {
	t.Helper()

	errors := p.Errors()
	if len(errors) > 0 {
		t.Errorf("parser had %d errors", len(errors))
		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}
}
