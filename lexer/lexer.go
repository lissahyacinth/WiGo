package lexer

import "wigo/token"

type NullByte struct {
	Byte  byte
	Valid bool
}

type Lexer struct {
	input    string
	position int // Current Position in input
}

func New(input string) *Lexer {
	l := &Lexer{input: input, position: 0}
	return l
}

func (l *Lexer) peek() NullByte {
	if l.position+1 < len(l.input) {
		return NullByte{l.input[l.position+1], true}
	}
	return NullByte{0, false}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	if l.atEnd() {
		tok.Literal = ""
		tok.Type = token.EOF
		return tok
	}

	switch l.cur() {
	case '=':
		if nextToken := l.peek(); nextToken.Valid && nextToken.Byte == '=' {
			// Consume the next token
			ch := string(l.cur())
			l.advance()
			tok = token.Token{Type: token.EQ, Literal: ch + string(l.cur())}
		} else {
			tok = newToken(token.ASSIGN, l.cur())
		}
	case '+':
		tok = newToken(token.PLUS, l.cur())
	case '/':
		tok = newToken(token.SLASH, l.cur())
	case '*':
		tok = newToken(token.ASTERISK, l.cur())
	case '<':
		tok = newToken(token.LT, l.cur())
	case '>':
		tok = newToken(token.GT, l.cur())
	case '-':
		tok = newToken(token.MINUS, l.cur())
	case '!':
		if nextToken := l.peek(); nextToken.Valid && nextToken.Byte == '=' {
			ch := string(l.cur())
			l.advance()
			tok = token.Token{Type: token.NOT_EQ, Literal: ch + string(l.cur())}
		} else {
			tok = newToken(token.BANG, l.cur())
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.cur())
	case '(':
		tok = newToken(token.LPAREN, l.cur())
	case ')':
		tok = newToken(token.RPAREN, l.cur())
	case ',':
		tok = newToken(token.COMMA, l.cur())
	case '{':
		tok = newToken(token.LBRACE, l.cur())
	case '}':
		tok = newToken(token.RBRACE, l.cur())
	default:
		if isLetter(l.cur()) {
			tok.Literal = l.readAsString(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.cur()) {
			tok.Type = token.INT
			tok.Literal = l.readAsString(isDigit)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.cur())
		}
	}

	l.advance()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for !l.atEnd() && (l.cur() == ' ' || l.cur() == '\t' || l.cur() == '\n' || l.cur() == '\r') {
		l.advance()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readAsString(detector func(byte) bool) string {
	position := l.position
	for !l.atEnd() && detector(l.cur()) {
		l.advance()
	}
	return l.input[position:l.position]
}

func (l *Lexer) atEnd() bool {
	return l.position >= len(l.input)
}

func (l *Lexer) cur() byte {
	return l.input[l.position]
}

func (l *Lexer) advance() {
	l.position++
}
