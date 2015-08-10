package uniql

import (
	"fmt"
	"io"
	"strings"
)

type ParseError struct {
	pos Pos
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Parse error at line %v col %v", e.pos.Line, e.pos.Col)
}

type Parser struct {
	s *Scanner
}

func NewParser(r io.Reader) *Parser {
	p := &Parser{s: NewScanner(r)}
	return p
}

func Parse(s string) (*BinaryExpression, error) {
	p := NewParser(strings.NewReader(s))

	return p.parseBinaryExpression()
}

func (p *Parser) parseBinaryExpression() (*BinaryExpression, error) {
	return nil, nil
}

func (p *Parser) parseUnaryExpression() (*Expression, error) {
	tok, pos, lit := p.nextToken()

	switch tok {
	case IDENTIFIER:
		return &Identifier{Name: lit}
	}

}

func (p *Parser) nextToken() (Token, Pos, string) {
	tok, pos, lit := p.s.Scan()

	if tok == WHITESPACE {
		return p.nextToken()
	}

	return tok, pos, lit
}
