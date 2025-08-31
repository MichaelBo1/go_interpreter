package ast

import (
	"testing"

	"github.com/MichaelBo1/go_interpreter/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "testVar"},
					Value: "testVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let testVar = anotherVar;" {
		t.Errorf("programin.String() wrong, got %q", program.String())
	}
}
