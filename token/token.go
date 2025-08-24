package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	UNKNOWN TokenType = iota
	EOF

	IDENTIFIER
	INT

	ASSIGN
	PLUS

	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FUNCTION
	LET
)

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

// There are easier ways to generate this, but this is just done manually for now.
func (t TokenType) String() string {
	switch t {
	case UNKNOWN:
		return "UNKNOWN"
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case INT:
		return "INT"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	default:
		return "UNKNOWN"
	}
}
