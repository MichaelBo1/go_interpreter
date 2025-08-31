package parser

import (
	"fmt"
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

func TestParsesIdentifierExpression(t *testing.T) {
	input := "foobar;"

	lex := lexer.New(input)
	par := New(lex)
	program := par.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	checkParserErrors(t, par)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. Got %d statements", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ExpressionStatementm got %T", stmt)
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier, got %T", ident)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input            string
		expectedOperator string
		expectedIntValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, test := range prefixTests {
		lex := lexer.New(test.input)
		par := New(lex)
		program := par.ParseProgram()

		if program == nil {
			t.Fatalf("ParseProgram() returned nil")
		}
		checkParserErrors(t, par)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != test.expectedOperator {
			t.Fatalf("exp operator is not '%s', got %s", test.expectedOperator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, test.expectedIntValue) {
			return
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

func testIntegerLiteral(t testing.TB, il ast.Expression, expectedValue int64) bool {
	t.Helper()

	literal, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("exp not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if literal.Value != expectedValue {
		t.Errorf("literal.Value not %d. got=%d", expectedValue, literal.Value)
		return false
	}

	if literal.TokenLiteral() != fmt.Sprintf("%d", expectedValue) {
		t.Errorf("literal.TokenLiteral not %d. got=%s", expectedValue,
			literal.TokenLiteral())
		return false
	}
	return true
}
