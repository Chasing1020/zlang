/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-11:01 PM
File: statement.go
*/

package parser

import (
	"errors"
	"zlang/ast"
	"zlang/ast/expression"
	"zlang/ast/statement"
	"zlang/token"
)

// parseLetStatement()
// Statement -> LET identifier = Expression
func (p *Parser) parseLetStatement() *statement.Let {
	stmt := &statement.Let{Token: p.curTok}
	// if the next statement is not an identifier
	if !p.expectPeek(token.Ident) {
		p.errs = append(p.errs, errors.New("let must continue with an identifier"))
		return nil
	}

	stmt.Name = &expression.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	if !p.expectPeek(token.Assign) {
		return nil
	}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.Semi) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement
// Statement -> RETURN Expression
func (p *Parser) parseReturnStatement() *statement.Return {
	stmt := &statement.Return{Token: p.curTok}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.Semi) {
		p.nextToken()
	}

	return stmt
}

// parseExpressionStatement
// Statement -> Expression
func (p *Parser) parseExpressionStatement() *statement.Expression {
	stmt := &statement.Expression{Token: p.curTok}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.Semi) {
		p.nextToken()
	}
	return stmt
}

// parseIfExpression
// Statement -> If '(' Expression ')' '{' Statement '}' (Else '{' Statement '}')?
func (p *Parser) parseIfExpression() ast.Expr {
	s := &statement.If{Token: p.curTok}
	if !p.expectPeek(token.Lparen) {
		return nil
	}
	p.nextToken()
	s.Condition = p.parseExpression(LOWEST)

	// TODO: determine if this a boolean expression
	// switch e.Condition.(type) {
	// }

	if !p.expectPeek(token.Rparen) || !p.expectPeek(token.Lbrace) {
		return nil
	}

	s.Consequence = p.parseBlockStatement()
	if p.peekTokenIs(token.Else) {
		p.nextToken()
		if !p.expectPeek(token.Lbrace) {
			return nil
		}
		s.Alternative = p.parseBlockStatement()
	}
	return s
}

// parseBlockStatement
// Block -> Statement*
func (p *Parser) parseBlockStatement() *statement.Block {
	b := &statement.Block{Token: p.curTok}
	p.nextToken()
	for !p.curTokenIs(token.Rbrace) && !p.curTokenIs(token.EOF) {
		stat := p.parseStatement()
		if stat != nil {
			b.Statements = append(b.Statements, stat)
		}
		p.nextToken()
	}
	return b
}

// parseFunction
// Statement -> Function( (Expression)* (, Expression)* ) Block
func (p *Parser) parseFunction() ast.Expr {
	f := &statement.Function{Token: p.curTok}
	if !p.expectPeek(token.Lparen) {
		return nil
	}
	f.Parameters = p.parseFunctionParameters()
	if !p.expectPeek(token.Lbrace) {
		return nil
	}
	f.Body = p.parseBlockStatement()
	return f
}

// parseFunctionParameters
// Parameters -> (Expression)* (, Expression)*
func (p *Parser) parseFunctionParameters() []*expression.Identifier {
	var identifiers []*expression.Identifier
	// function has no parameters
	if p.peekTokenIs(token.Rparen) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()
	ident := &expression.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.Comma) {
		p.nextToken()
		p.nextToken()
		identifiers = append(identifiers, &expression.Identifier{Token: p.curTok, Value: p.curTok.Literal})
	}

	if !p.expectPeek(token.Rparen) {
		return nil
	}
	return identifiers
}

// parseForStatement
// For -> for '(' InitStat; Condition; UpdateStat ')' { Body }
func (p *Parser) parseForStatement() *statement.For {
	f := &statement.For{Token: p.curTok}
	p.nextToken()
	p.nextToken()
	if p.curTok.Type == token.Semi {
		f.InitStat = nil
	} else {
		f.InitStat = p.parseLetStatement()
	}
	//p.nextToken()
	if p.curTokenIs(token.Semi) {
		p.nextToken()
	}
	f.Condition = p.parseExpression(LOWEST)
	p.nextToken()
	if p.curTokenIs(token.Semi) {
		p.nextToken()
	}
	if p.curTok.Type == token.Rparen {
		f.UpdateStat = nil
	} else {
		f.UpdateStat = p.parseStatement()
	}
	p.nextToken()
	f.Body = p.parseBlockStatement()
	return f
}

// s := &statement.If{Token: p.curTok}
//	if !p.expectPeek(token.Lparen) {
//		return nil
//	}
//	p.nextToken()
//	s.Condition = p.parseExpression(LOWEST)
//
//	// TODO: determine if this a boolean expression
//	// switch e.Condition.(type) {
//	// }
//
//	if !p.expectPeek(token.Rparen) || !p.expectPeek(token.Lbrace) {
//		return nil
//	}
//
//	s.Consequence = p.parseBlockStatement()
//	if p.peekTokenIs(token.Else) {
//		p.nextToken()
//		if !p.expectPeek(token.Lbrace) {
//			return nil
//		}
//		s.Alternative = p.parseBlockStatement()
//	}
//	return s
