/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-2:36 PM
File: source.go
*/

package scanner

import (
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

const (
	baseLine   = 1
	baseColumn = 1
	sentinel   = utf8.RuneSelf
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
	s.err = errors.New(msg)
	s.errHandler(line, col, msg)
}

func (s *source) errorf(format string, args ...interface{}) {
	s.error(fmt.Sprintf(format, args...))
}


func (s *source) nextCh() {
	if s.next >= len(s.buf) {
		s.err = io.EOF
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

