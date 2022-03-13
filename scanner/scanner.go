/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-8:36 PM
File: scanner.go
*/

package scanner

import (
	"io"
	"unicode/utf8"
	"zlang/token"
)

type Scanner struct {
	source
	//Tok     token.Type
	//Literal string
	token.Token
	kind token.LitKind
}

func (s *Scanner) Init(buf string, errHandler func(line, col uint, msg string)) {
	s.source.init(buf, errHandler)
	s.ch = s.buf[0]
	s.nextCh()
}

// NextTok Get the NextTok token
func (s *Scanner) NextTok() {
	if s.err != nil && s.err != io.EOF {
		panic(s.err)
	}

	// skip blank
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.nextCh()
	}

	switch s.ch {
	case '=':
		if s.peekChar() == '=' {
			ch := s.ch
			s.nextCh()
			s.Token = token.Token{Type: token.Eql, Literal: string(ch) + string(s.ch)}
		} else {
			s.Token = token.Token{Type: token.Assign, Literal: string(s.ch)}
		}
	case '+':
		// TODO: add incr
		s.Token = token.Token{Type: token.Plus, Literal: string(s.ch)}
	case '-':
		// TODO: add decr
		s.Token = token.Token{Type: token.Minus, Literal: string(s.ch)}
	case '!':
		if s.peekChar() == '=' {
			ch := s.ch
			s.nextCh()
			s.Token = token.Token{Type: token.Neq, Literal: string(ch)}
		} else {
			s.Token = token.Token{Type: token.Bang, Literal: string(s.ch)}
		}
	case '/':
		s.Token = token.Token{Type: token.Slash, Literal: string(s.ch)}
	case '*':
		s.Token = token.Token{Type: token.Star, Literal: string(s.ch)}
	case '<':
		if s.peekChar() == '=' {
			ch := s.ch
			s.nextCh()
			s.Token = token.Token{Type: token.Leq, Literal: string(ch) + string(s.ch)}
		} else {
			s.Token = token.Token{Type: token.Lss, Literal: string(s.ch)}
		}
	case '>':
		if s.peekChar() == '=' {
			ch := s.ch
			s.nextCh()
			s.Token = token.Token{Type: token.Geq, Literal: string(ch) + string(s.ch)}
		} else {
			s.Token = token.Token{Type: token.Gtr, Literal: string(s.ch)}
		}
	case ';':
		s.Token = token.Token{Type: token.Semi, Literal: string(s.ch)}
	case ',':
		s.Token = token.Token{Type: token.Comma, Literal: string(s.ch)}
	case '{':
		s.Token = token.Token{Type: token.Lbrace, Literal: string(s.ch)}
	case '}':
		s.Token = token.Token{Type: token.Rbrace, Literal: string(s.ch)}
	case '(':
		s.Token = token.Token{Type: token.Lparen, Literal: string(s.ch)}
	case ')':
		s.Token = token.Token{Type: token.Rparen, Literal: string(s.ch)}
	case '[':
		s.Token = token.Token{Type: token.Lbrack, Literal: string(s.ch)}
	case ']':
		s.Token = token.Token{Type: token.Rbrack, Literal: string(s.ch)}
	case ':':
		s.Token = token.Token{Type: token.Colon, Literal: string(s.ch)}
	case '"':
		s.Token = token.Token{Type: token.String, Literal: s.readString()}
	case 0:
		s.Token = token.Token{Type: token.EOF, Literal: string(s.ch)}
	default:
		// if current ch can be an identifier
		if isLetter(s.ch) && s.readIdentChar(true) {
			literal := s.ident()
			s.Token = token.Token{Type: checkIdent(literal), Literal: literal}
			return
		} else if isDigit(s.ch) {
			s.Token = token.Token{Type: token.Int, Literal: s.readNumber()}
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
	if s.index+1 >= len(s.buf) {
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

func (s *Scanner) readString() (res string) {
	s.nextCh()
	index := s.index
	for s.ch != '"' && s.err != io.EOF {
		s.nextCh()
	}
	res = s.buf[index:s.index]
	//s.nextCh()
	return
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
func checkIdent(ident string) token.Type {
	if keyword, ok := token.KeywordMap[ident]; ok {
		return keyword
	}
	return token.Ident
}
