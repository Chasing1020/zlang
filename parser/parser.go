/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-10:34 PM
File: parser.go
*/

package parser

import (
	"log"
	"zlang/ast"
	"zlang/scanner"
	"zlang/token"
)

type Parser struct {
	scanner.Scanner
	err []error

	curTok  token.Token
	peekTok token.Token

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) init(buf string) {
	p.Scanner.Init(buf, func(line, col uint, msg string) {
		log.Println("compiler error:", "msg:", msg, "line:", line, "col:", col)
	})
}
