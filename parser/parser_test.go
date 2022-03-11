/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-12:22 AM
File: parser_test.go
*/

package parser

import (
	"fmt"
	"testing"
	"zlang/scanner"
	"zlang/token"
)

func TestParser(t *testing.T) {
	p := Parser{
		Scanner:        scanner.Scanner{},
		errs:           nil,
		curTok:         token.Token{},
		peekTok:        token.Token{},
		prefixParseFns: nil,
		infixParseFns:  nil,
	}
	p.init("function(a, b) {return a + b;}")
	file := p.ParseFile()
	fmt.Println(file.Stats)
}
