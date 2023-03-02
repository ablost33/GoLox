package scanner

import (
	"strconv"

	"../error"
	"../token"
)

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	token   []token.Token
}

func (scanner *Scanner) scanTokens() []token.Token {
	for !scanner.isAtEnd() {
		scanner.start = scanner.current
		scanner.scanToken()
	}
	scanner.token = append(scanner.token, token.Token{
		TokenType: token.EOF,
		Line:      scanner.line,
	},
	)
	return scanner.token
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (sc *Scanner) scanToken() {
	c := sc.advance()
	switch c {
	case '(':
		sc.addToken(token.LEFT_PAREN)
	case ')':
		sc.addToken(token.RIGHT_PAREN)
	case '{':
		sc.addToken(token.LEFT_BRACE)
	case '}':
		sc.addToken(token.RIGHT_BRACE)
	case ',':
		sc.addToken(token.COMMA)
	case '.':
		sc.addToken(token.DOT)
	case '-':
		sc.addToken(token.MINUS)
	case '+':
		sc.addToken(token.PLUS)
	case ';':
		sc.addToken(token.SEMICOLON)
	case '*':
		sc.addToken(token.STAR)
	case '!':
		if sc.match('=') {
			sc.addToken(token.BANG_EQUAL)
		} else {
			sc.addToken(token.BANG)
		}
	case '=':
		if sc.match('=') {
			sc.addToken(token.EQUAL_EQUAL)
		} else {
			sc.addToken(token.EQUAL)
		}
	case '<':
		if sc.match('=') {
			sc.addToken(token.LESS_EQUAL)
		} else {
			sc.addToken(token.LESS)
		}
	case '>':
		if sc.match('=') {
			sc.addToken(token.GREATER_EQUAL)
		} else {
			sc.addToken(token.GREATER)
		}
	case '/':
		if sc.match('/') {
			for sc.peek() != '\n' && !sc.isAtEnd() {
				sc.advance()
			}
		} else {
			sc.addToken(token.SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		sc.line++
	case '"':
		sc.string()
	default:
		if isDigit(c) {
			sc.number()
		} else {
			error.ReportError(sc.line, "Unexpected character.")
		}
	}
}

func (sc *Scanner) match(expected byte) bool {
	if sc.isAtEnd() {
		return false
	}
	if sc.source[sc.current] != expected {
		return false
	}
	sc.current++
	return true
}

func (sc *Scanner) advance() byte {
	sc.current++
	return sc.source[sc.current-1]
}

func (sc *Scanner) addToken(tokenType token.TokenType) {
	sc.addTokenType(tokenType, nil)
}

func (sc *Scanner) addTokenType(tokenType token.TokenType, literal interface{}) {
	text := sc.source[sc.start:sc.current]
	token := token.Token{
		TokenType: tokenType,
		Lexeme:    text,
		Literal:   literal,
		Line:      sc.line,
	}
	sc.token = append(sc.token, token)
}

func (sc *Scanner) peek() byte {
	if sc.isAtEnd() {
		return 0
	}
	return sc.source[sc.current]
}

func (sc *Scanner) string() {
	for sc.peek() != '"' && !sc.isAtEnd() {
		if sc.peek() == '\n' {
			sc.line++
		}
		sc.advance()
	}
	if sc.isAtEnd() {
		error.ReportError(sc.line, "Unterminated string.")
		return
	}
	sc.advance()
	value := sc.source[sc.start+1 : sc.current-1]
	sc.addTokenType(token.STRING, value)
}

func isDigit(c byte) bool {
	return c >= 0 && c <= 9
}

func (sc *Scanner) number() {
	for isDigit(sc.peek()) {
		sc.advance()
	}
	if sc.peek() == '.' && isDigit(sc.peekNext()) {
		sc.advance()
		for isDigit(sc.peek()) {
			sc.advance()
		}
	}
	number, err := strconv.ParseFloat(sc.source[sc.start:sc.current], 64)
	if err != nil {
		error.ReportError(sc.line, "Invalid number format")
		return
	}
	sc.addTokenType(token.NUMBER, number)
}
