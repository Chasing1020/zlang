/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-10:34 PM
File: parser.go
*/

package parser

import (
	"fmt"
	"log"
	"zlang/ast"
	"zlang/scanner"
	"zlang/token"
)

type Precedence uint

const (
	_      Precedence = iota
	LOWEST       // lowest
	ASSIGN       // assignment
	EQUALS       // ==
	LESS_GREATER // > or <
	SUM          // +
	PRODUCT      // *
	PREFIX       // -X or !X
	CALL         // myFunction(X)
	INDEX        // array[index]
)

var PrecedenceMap = map[token.Type]Precedence{
	token.Assign: ASSIGN,
	token.Eql:    EQUALS,
	token.Neq:    EQUALS,
	token.Lss:    LESS_GREATER,
	token.Gtr:    LESS_GREATER,
	token.Plus:   SUM,
	token.Minus:  SUM,
	token.Slash:  PRODUCT,
	token.Star:   PRODUCT,
	token.Lparen: CALL,
	token.Lbrack: INDEX,
}

type Parser struct {
	scanner.Scanner
	errs []error

	curTok  token.Token
	peekTok token.Token
} // Read two tokens, so curToken and peekToken are both set

func (p *Parser) Init(buf string) {
	p.Scanner.Init(buf, func(line, col uint, msg string) {
		log.Println("compiler error:", msg, "line:", line, "col:", col)
	})
	p.nextToken()
	p.nextToken()
}

func (p *Parser) getPrefixParseFunc() func() ast.Expr {
	switch p.curTok.Type {
	case token.Ident:
		return p.parseIdentifier
	case token.Int:
		return p.parseInteger
	case token.String:
		return p.parseString
	case token.Bang, token.Minus:
		return p.parsePrefixExpression
	case token.True, token.False:
		return p.parseBoolean
	case token.Lparen:
		return p.parseGroupedExpression
	case token.If:
		return p.parseIfExpression
	case token.Function:
		return p.parseFunction
	case token.Lbrack:
		return p.parseArray
	case token.Lbrace:
		return p.parseMapLiteral
	}
	return nil
}

func (p *Parser) getInfixParseFunc() func(expr ast.Expr) ast.Expr {
	switch p.peekTok.Type {
	case token.Plus, token.Minus, token.Star, token.Slash, // +, -, *, /
		token.Eql, token.Neq, token.Lss, token.Gtr: // ==, != , <, >
		return p.parseInfixExpression
	case token.Lparen:
		return p.parseCallExpression
	case token.Lbrack:
		return p.parseIndexExpression

	}
	return nil
}

func (p *Parser) ParseFile() *ast.File {
	file := &ast.File{}

	for !p.curTokenIs(token.EOF) {
		stat := p.parseStatement()
		if stat != nil {
			file.Stats = append(file.Stats, stat)
		}
		p.nextToken()
	}
	return file
}

func (p *Parser) parseStatement() ast.Stat {
	switch p.curTok.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	case token.For:
		return p.parseForStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) peekPrecedence() Precedence {
	if precedence, ok := PrecedenceMap[p.peekTok.Type]; ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) curPrecedence() Precedence {
	if precedence, ok := PrecedenceMap[p.curTok.Type]; ok {
		return precedence
	}
	return LOWEST
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.NextTok()
	p.peekTok = p.Scanner.Token
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curTok.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekTok.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.errs = append(p.errs, fmt.Errorf("expected next token to be %s, got %s instead",
			token.Map[t], p.peekTok.Type.String()))
		return false
	}
}
