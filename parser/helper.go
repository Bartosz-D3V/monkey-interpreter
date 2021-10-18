package parser

import (
	"fmt"
	"monkey_interpreter/ast"
	"monkey_interpreter/token"
	"strconv"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//TODO For now we're skipping tokens until semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{
		Token: p.curToken,
	}

	//TODO For now we're skipping tokens until semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.curToken,
	}
	stmt.Expression = p.parseExpression(LOWEST)

	//TODO For now we're skipping tokens until semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{
		Token: p.curToken,
	}
	val, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Could not parse %s into int", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = val
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	exp.Right = p.parseExpression(PREFIX)

	return exp
}

func (p *Parser) parseBoolean() ast.Expression {
	parseBool, err := strconv.ParseBool(p.curToken.Literal)
	if err != nil {
		msg := fmt.Sprintf("Invalid boolean value %s", p.curToken.Literal)
		p.errors = append(p.errors, msg)
	}

	exp := &ast.Boolean{
		Token: p.curToken,
		Value: parseBool,
	}

	p.nextToken()

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:     p.curToken,
		LeftValue: left,
		Operator:  p.curToken.Literal,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	exp.RightValue = p.parseExpression(precedence)

	return exp
}

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("Missing prefixParseFn for token %s", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.Type) string {
	msg := fmt.Sprintf("Expected next token to be '%s' - got '%s' instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
	return msg
}

func (p *Parser) registerPrefixParseFn(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfixParseFn(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}