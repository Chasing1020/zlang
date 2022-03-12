/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-11:09 PM
File: expression.go
*/

package parser

import (
	"errors"
	"fmt"
	"strconv"
	"zlang/ast"
	"zlang/ast/expression"
	"zlang/token"
)

// parseExpression
// Expression -> Expression (+ | - | * | /) Expression  // Infix expression
//             | (- | !) Expression                     // Prefix expression
//             | (Expression)                           // Infix expression
//             | Literal                                // Infix expression
//             | Function '{' Statement '}'             // Prefix Expression
func (p *Parser) parseExpression(precedence Precedence) ast.Expr {
	prefixFunc := p.getPrefixParseFunc()
	//prefixFunc := p.prefixParseFns[p.curTok.Type]
	if prefixFunc == nil {
		p.errs = append(p.errs, errors.New(p.curTok.Type.String()))
		return nil
	}
	leftExp := prefixFunc()

	for !p.peekTokenIs(token.Semi) && precedence < p.peekPrecedence() {
		//infixFunc := p.infixParseFns[p.curTok.Type]
		infixFunc := p.getInfixParseFunc()
		if infixFunc == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infixFunc(leftExp)
	}
	return leftExp
}

// parsePrefixExpression
// set priority as PREFIX
func (p *Parser) parsePrefixExpression() ast.Expr {
	e := &expression.Prefix{
		Token:    p.curTok,
		Operator: p.curTok.Literal,
	}

	p.nextToken()
	e.Right = p.parseExpression(PREFIX)

	return e
}

// parseInfixExpression
// precedence depends on the current
func (p *Parser) parseInfixExpression(left ast.Expr) ast.Expr {
	e := &expression.Infix{
		Token:    p.curTok,
		Operator: p.curTok.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	e.Right = p.parseExpression(precedence)

	return e
}

// parseIdentifier
// Literal -> Identifier
func (p *Parser) parseIdentifier() ast.Expr {
	return &expression.Identifier{Token: p.curTok, Value: p.curTok.Literal}
}

// parseInteger
// Literal -> Integer
func (p *Parser) parseInteger() ast.Expr {
	lit := &expression.Integer{Token: p.curTok}

	value, err := strconv.Atoi(p.curTok.Literal)
	if err != nil {
		p.errs = append(p.errs,
			fmt.Errorf("could not parse %s as integer, detail: %s", p.curTok.Literal, err.Error()))
		return nil
	}

	lit.Value = int64(value)
	return lit
}

// parseString
// Literal -> String
func (p *Parser) parseString() ast.Expr {
	return &expression.String{Token: p.curTok, Value: p.curTok.Literal}
}

// parseBoolean
// Literal -> Boolean
func (p *Parser) parseBoolean() ast.Expr {
	return &expression.Boolean{Token: p.curTok, Value: p.curTokenIs(token.True)}
}

// parseGroupedExpression
// Expression -> '(' Expression ')'
func (p *Parser) parseGroupedExpression() ast.Expr {
	p.nextToken()

	e := p.parseExpression(LOWEST)

	if !p.expectPeek(token.Rparen) {
		p.errs = append(p.errs, errors.New("right paren not found"))
		return nil
	}
	return e
}

// parseCallExpression
// Expression -> FunctionName( ExpressionList )
func (p *Parser) parseCallExpression(function ast.Expr) ast.Expr {
	exp := &expression.Call{Token: p.curTok, Function: function}
	exp.Arguments = p.parseExpressionList(token.Rparen)
	return exp
}

// parseExpressionList
// ExpressionList -> Expression* (, Expression)*
func (p *Parser) parseExpressionList(end token.Type) []ast.Expr {
	var list []ast.Expr

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.Comma) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

// parseArray
// Expression -> [(Expression)* (, Expression)*]
func (p *Parser) parseArray() ast.Expr {
	array := &expression.Array{Token: p.curTok}
	array.Elements = p.parseExpressionList(token.Rbrack)
	return array
}

// parseIndexExpression
// Expression -> ArrayName[ Integer ]
func (p *Parser) parseIndexExpression(left ast.Expr) ast.Expr {
	exp := &expression.Index{Token: p.curTok, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.Rbrack) {
		return nil
	}

	return exp
}

// parseMapLiteral
// Expression -> { ("Identifier": Expr)* (, "Identifier": Expr)* }
func (p *Parser) parseMapLiteral() ast.Expr {
	hash := &expression.Map{Token: p.curTok}
	hash.Pairs = make(map[ast.Expr]ast.Expr)

	for !p.peekTokenIs(token.Rbrace) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.expectPeek(token.Colon) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)

		hash.Pairs[key] = value

		if !p.peekTokenIs(token.Rbrace) && !p.expectPeek(token.Comma) {
			return nil
		}
	}

	if !p.expectPeek(token.Rbrace) {
		return nil
	}
	return hash
}
