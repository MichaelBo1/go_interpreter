package lexer

import (
	"github.com/MichaelBo1/go_interpreter/token"
)

type Lexer struct {
	input      string
	currentPos int
	nextPos    int
	ch         byte
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}
	lexer.readChar()
	return lexer
}

// TODO: this doesn't support Unicode (& UTF-8), which would need to use runes and would also
// need to work for multi-byte-length encodings.
func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0 // Signifier for EOF
	} else {
		l.ch = l.input[l.nextPos]
	}

	l.currentPos = l.nextPos
	l.nextPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	// TODO: extract the two-char token logic.
	switch l.ch {
	case '=':
		if l.peek() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = token.NewToken(token.ASSIGN, string(l.ch))
		}
	case '+':
		tok = token.NewToken(token.PLUS, string(l.ch))
	case '-':
		tok = token.NewToken(token.MINUS, string(l.ch))
	case '/':
		tok = token.NewToken(token.SLASH, string(l.ch))
	case '!':
		if l.peek() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewToken(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = token.NewToken(token.BANG, string(l.ch))
		}
	case '*':
		tok = token.NewToken(token.ASTERISK, string(l.ch))
	case '<':
		if l.peek() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewToken(token.LESS_THAN_OR_EQ, string(ch)+string(l.ch))
		} else {
			tok = token.NewToken(token.LESS_THAN, string(l.ch))
		}
	case '>':
		if l.peek() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewToken(token.GREATER_THAN_OR_EQ, string(ch)+string(l.ch))
		} else {
			tok = token.NewToken(token.GREATER_THAN, string(l.ch))
		}
	case ',':
		tok = token.NewToken(token.COMMA, string(l.ch))
	case ';':
		tok = token.NewToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = token.NewToken(token.LPAREN, string(l.ch))
	case ')':
		tok = token.NewToken(token.RPAREN, string(l.ch))
	case '{':
		tok = token.NewToken(token.LBRACE, string(l.ch))
	case '}':
		tok = token.NewToken(token.RBRACE, string(l.ch))
	case 0:
		tok = token.NewToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.FindIdentifier(tok.Literal)
			return tok // Early exit as `readIdentifier` calls readChar() and eats the input.
		}
		if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readInt()
			return tok
		}
		tok = token.NewToken(token.UNKNOWN, string(l.ch))
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.currentPos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.currentPos]
}

func (l *Lexer) readInt() string {
	pos := l.currentPos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.currentPos]
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Simplified version here since currently we are only dealing with ASCII chars.
// See (TODO) above for supporting Unicode.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peek() byte {
	if l.nextPos >= len(l.input) {
		return 0
	}
	return l.input[l.nextPos]
}
