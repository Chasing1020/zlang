package scanner

import (
	"io"
	"unicode/utf8"
	"zjclang/token"
)

type Scanner struct {
	source
	Tok     token.Token
	Literal string
	kind    token.LitKind
}

func (s *Scanner) Init(buf string, errHandler func(line, col uint, msg string)) {
	s.source.init(buf, errHandler)
	s.ch = s.buf[0]
	s.nextCh()
}

func (s *Scanner) Next() {

	// skip blank
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.nextCh()
	}

	switch s.ch {

	case '=':
		if s.peekChar() == '=' {
			ch := s.ch
			s.Next()
			s.Literal = string(ch) + string(s.ch)
			s.Tok = token.Eql
		} else {
			s.Tok = token.Assign
			s.Literal = string(s.ch)
		}
	case '+':
		s.Tok = token.Plus
		s.Literal = string(s.ch)
	case '-':
		s.Tok = token.Minus
		s.Literal = string(s.ch)
	case '!':
		if s.peekChar() == '=' {
			ch := s.ch
			s.nextCh()
			s.Tok = token.Neq
			s.Literal = string(ch) + string(s.ch)
		} else {
			s.Tok = token.Bang
			s.Literal = string(s.ch)
		}
	case '/':
		s.Tok = token.Slash
		s.Literal = string(s.ch)
	case '*':
		s.Tok = token.Star
		s.Literal = string(s.ch)
	case '<':
		s.Tok = token.Lss
		s.Literal = string(s.ch)
	case '>':
		s.Tok = token.Gtr
		s.Literal = string(s.ch)
	case ';':
		s.Tok = token.Semi
		s.Literal = string(s.ch)
	case ',':
		s.Tok = token.Comma
		s.Literal = string(s.ch)
	case '{':
		s.Tok = token.Lbrace
		s.Literal = string(s.ch)
	case '}':
		s.Tok = token.Rbrace
		s.Literal = string(s.ch)
	case '(':
		s.Tok = token.Lparen
		s.Literal = string(s.ch)
	case ')':
		s.Tok = token.Rparen
		s.Literal = string(s.ch)
	case 0:
		s.Literal = ""
		s.Tok = token.EOF
	default:
		// if current ch can be an identifier
		if isLetter(s.ch) && s.readIdentChar(true) {
			s.Literal = s.ident()
			s.Tok = checkIdent(s.Literal)
			return
		} else if isDigit(s.ch) {
			s.Tok = token.Int
			s.Literal = s.readNumber()
			return
		} else {
			s.errorf("invalid character %#U in identifier", s.ch)
			return
		}
	}

	s.nextCh()
	return
}

func (s *Scanner) peekChar() byte {
	if s.index+1 > len(s.buf) {
		s.err = io.EOF
		return 0
	} else {
		return s.buf[s.index+1]
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (s *Scanner) readIdentChar(isFirst bool) bool {
	switch {
	case isLetter(s.ch) || s.ch == '_':
	case isDigit(s.ch):
		if isFirst {
			s.errorf("identifier cannot begin with digit %#U", s.ch)
		}
	case s.ch >= utf8.RuneSelf:
		s.errorf("invalid character %#U in identifier", s.ch)
	default:
		return false
	}
	return true
}

func (s *Scanner) ident() string {
	index := s.index
	// accelerate common case (7bit ASCII)
	for isLetter(s.ch) || isDigit(s.ch) {
		s.nextCh()
	}
	// general case
	if s.ch >= utf8.RuneSelf {
		for s.readIdentChar(false) {
			s.nextCh()
		}
	}
	return s.buf[index:s.index]
}

func (s *Scanner) readNumber() string {
	index := s.index
	for isDigit(s.ch) {
		s.nextCh()
	}
	return s.buf[index:s.index]
}

// checkIdent verify whether an indent is a keyword
func checkIdent(ident string) token.Token {
	if keyword, ok := token.KeywordMap[ident]; ok {
		return keyword
	}
	return token.Ident
}
