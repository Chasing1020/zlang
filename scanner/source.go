/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-2:36 PM
File: source.go
*/

package scanner

import (
	"fmt"
	"unicode/utf8"
)

const (
	baseLine   = 1
	baseColumn = 1
	sentinel = utf8.RuneSelf
)

// source represents a source of an input buffer
type source struct {
	buf        string
	ch         byte
	line       uint
	col        uint
	index      int // current position in input
	next       int // index + 1 represents the next ch to read
	errHandler func(line, col uint, msg string)
	err        error
}

func (s *source) init(buf string, errHandler func(line, col uint, msg string)) {
	s.buf = buf
	s.ch = ' '
	s.errHandler = errHandler
}

func (s *source) pos() (line, col uint) {
	return s.line + baseLine, s.col + baseColumn
}

// error reports the error msg at source position s.pos().
func (s *source) error(msg string) {
	line, col := s.pos()
	s.errHandler(line, col, msg)
}

func (s *source) errorf(format string, args ...interface{}) {
	s.error(fmt.Sprintf(format, args...))
}

//func (s *source) nextCh() {
//
//
//	if s.index+1 > len(s.buf) {
//		s.ch = 0
//		s.err = io.EOF
//	} else {
//		s.ch = s.buf[s.index]
//	}
//	s.index++
//}


func (s *source) nextCh() {
	if s.next >= len(s.buf) {
		s.ch = 0
	} else {
		s.ch = s.buf[s.next]
	}

	s.col += 1
	if s.ch == '\n' {
		s.line++
		s.col = 0
	}

	s.index = s.next
	s.next++
}
//func error (line, col uint, msg string) {
//	if msg[0] != '/' {
//		// error
//		if msg != "comment not terminated" {
//			t.Errorf("%q: %s", test.src, msg)
//		}
//		return
//	}
//	got = comment{line - linebase, col - colbase, msg} // keep last one
//}
