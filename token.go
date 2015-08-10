package uniql

import (
	"strings"
)

type Pos struct {
	Line int
	Col  int
}

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WHITESPACE

	literal_beg
	IDENTIFIER
	NUMBER
	STRING
	TRUE
	FALSE
	NULL
	UNDEFINED
	literal_end

	operator_beg
	LPAREN
	RPAREN
	EQ
	NEQ
	GT
	GTE
	LT
	LTE
	REG_MATCH
	AND
	OR
	NOT
	IN
	operator_end
)

var tokens = [...]string{
	ILLEGAL:    "ILLEGAL",
	EOF:        "EOF",
	WHITESPACE: "WHITESPACE",

	IDENTIFIER: "IDENTIFIER",
	NUMBER:     "NUMBER",
	TRUE:       "TRUE",
	FALSE:      "FALSE",
	NULL:       "NULL",
	UNDEFINED:  "UNDEFINED",
	STRING:     "STRING",

	LPAREN:    "LPAREN",
	RPAREN:    "RPAREN",
	EQ:        "EQ",
	NEQ:       "NEQ",
	GT:        "GT",
	GTE:       "GTE",
	LT:        "LT",
	LTE:       "LTE",
	REG_MATCH: "REG_MATCH",
	AND:       "AND",
	OR:        "OR",
	NOT:       "NOT",
	IN:        "IN",
}

var keywords = map[string]Token{
	"true":      TRUE,
	"false":     FALSE,
	"null":      NULL,
	"undefined": UNDEFINED,
	"and":       AND,
	"or":        OR,
	"not":       NOT,
	"in":        IN,
}

func (tok Token) String() string {
	return tokens[tok]
}

func LookupToken(ident string) Token {
	if tok, exists := keywords[strings.ToLower(ident)]; exists {
		return tok
	}
	return IDENTIFIER
}
