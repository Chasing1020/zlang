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
	_ Precedence = iota
	LOWEST
	EQUALS       // ==
	LESS_GREATER // > or <

	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL    // myFunction(X)
	INDEX   // array[index]
)

type (
	prefixParseFn func() ast.Expr
	infixParseFn  func(ast.Expr) ast.Expr
)

var PrecedenceMap = map[token.Type]Precedence{
	token.Eql:    EQUALS,
	token.Neq:    EQUALS,
	token.Lss:    LESS_GREATER,
	token.Gtr:    LESS_GREATER,
	token.Plus:   SUM,
	token.Minus:  SUM,
	token.Slash:  PRODUCT,
	token.Star:   PRODUCT,
	token.Lparen: CALL,
	token.Rparen: INDEX,
}

type Parser struct {
	scanner.Scanner
	errs []error

	curTok  token.Token
	peekTok token.Token

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) init(buf string) {
	p.Scanner.Init(buf, func(line, col uint, msg string) {
		log.Println("compiler error:", "msg:", msg, "line:", line, "col:", col)
	})
	p.prefixParseFns = make(map[token.Type]prefixParseFn)
	//registerPrefix
	p.registerPrefix(token.Ident, p.parseIdentifier)
	p.registerPrefix(token.Int, p.parseInteger)
	p.registerPrefix(token.String, p.parseString)
	p.registerPrefix(token.Bang, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)
	p.registerPrefix(token.True, p.parseBoolean)
	p.registerPrefix(token.False, p.parseBoolean)
	p.registerPrefix(token.Lparen, p.parseGroupedExpression)
	p.registerPrefix(token.If, p.parseIfExpression)
	p.registerPrefix(token.Function, p.parseFunction)
	p.registerPrefix(token.Lbrack, p.parseArray)
	p.registerPrefix(token.Lbrace, p.parseMapLiteral)

	p.infixParseFns = make(map[token.Type]infixParseFn)
	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Minus, p.parseInfixExpression)
	p.registerInfix(token.Slash, p.parseInfixExpression)
	p.registerInfix(token.Star, p.parseInfixExpression)
	p.registerInfix(token.Eql, p.parseInfixExpression)
	p.registerInfix(token.Neq, p.parseInfixExpression)
	p.registerInfix(token.Lss, p.parseInfixExpression)
	p.registerInfix(token.Gtr, p.parseInfixExpression)

	p.registerInfix(token.Lparen, p.parseCallExpression)
	p.registerInfix(token.Lbrack, p.parseIndexExpression)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
}

func (p *Parser) ParseFile() *ast.File {
	file := &ast.File{}
	file.Stats = []ast.Stat{}

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
			token.TokenMap[t], p.peekTok.Type.String()))
		return false
	}
}
