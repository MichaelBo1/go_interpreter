package lexer

import "github.com/MichaelBo1/go_interpreter/token"

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
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPos]
	}

	l.currentPos = l.nextPos
	l.nextPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, string(l.ch))
	case '+':
		tok = token.NewToken(token.PLUS, string(l.ch))
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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.FindIdentifier(tok.Literal)
			return tok // Early exit as `readIdentifier` calls readChar() and eats the input.
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
