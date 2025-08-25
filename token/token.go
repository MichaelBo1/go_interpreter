package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func FindIdentifier(identifier string) TokenType {
	if tokType, ok := keywords[identifier]; ok {
		return tokType
	}
	return IDENTIFIER // User-defined identifier.
}

const (
	UNKNOWN TokenType = iota
	EOF

	IDENTIFIER
	INT

	ASSIGN
	PLUS
	MINUS

	EQ
	NOT_EQ

	SLASH
	BANG
	ASTERISK

	LESS_THAN
	LESS_THAN_OR_EQ
	GREATER_THAN
	GREATER_THAN_OR_EQ

	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FUNCTION
	LET
	IF
	ELSE
	RETURN
	TRUE
	FALSE
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
	case MINUS:
		return "MINUS"
	case EQ:
		return "EQ"
	case NOT_EQ:
		return "NOT_EQ"
	case SLASH:
		return "SLASH"
	case BANG:
		return "BANG"
	case ASTERISK:
		return "ASTERISK"
	case LESS_THAN:
		return "LESS_THAN"
	case LESS_THAN_OR_EQ:
		return "LESS_THAN_OR_EQ"
	case GREATER_THAN:
		return "GREATER_THAN"
	case GREATER_THAN_OR_EQ:
		return "GREATER_THAN_OR_EQ"
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
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case RETURN:
		return "RETURN"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	default:
		return "UNKNOWN"
	}
}
