package uniql

import (
	"strings"
	"testing"
)

func TestScanIdentifier(t *testing.T) {
	tok, _, lit := makeScanner("test").Scan()

	if tok != IDENTIFIER {
		t.Errorf("Unexpected token received: %v", tok)
	}
	if lit != "test" {
		t.Errorf("Unexpected token lit received: '%v'", lit)
	}
}

func TestScanIdentifierPreceedingWhitespace(t *testing.T) {
	tok, _, lit := makeScanner("hello world").Scan()

	if tok != IDENTIFIER {
		t.Errorf("Unexpected token received: %v", tok)
	}
	if lit != "hello" {
		t.Errorf("Unexpected token lit received: '%v'", lit)
	}
}

func TestScanKeywords(t *testing.T) {
	for keyword, token := range keywords {
		tok, _, _ := makeScanner(keyword).Scan()

		if tok != token {
			t.Errorf("Unexpected token received: %v", tok)
		}
	}
}

func TestScanKeywordsPreceedingWhitespace(t *testing.T) {
	for keyword, token := range keywords {
		tok, _, _ := makeScanner(keyword + "  \t  and").Scan()

		if tok != token {
			t.Errorf("Unexpected token received: %v", tok)
		}
	}
}

func TestScanNumber(t *testing.T) {
	tok, _, lit := makeScanner("123.123").Scan()

	if tok != NUMBER {
		t.Errorf("Unexpected token received: %v", tok)
	}
	if lit != "123.123" {
		t.Errorf("Unexpected token lit received: '%v'", lit)
	}
}

func TestScanNumberIllegal(t *testing.T) {
	tok, pos, _ := makeScanner("123.").Scan()

	if tok != ILLEGAL {
		t.Errorf("Unexpected token received: %v", tok)
	}
	if pos.Col != 4 {
		t.Errorf("Unexpected illegal pos: %v", pos)
	}
}

func TestScanNumberIllegal2(t *testing.T) {
	tok, pos, _ := makeScanner("123.123.").Scan()

	if tok != ILLEGAL {
		t.Errorf("Unexpected token received: %v", tok)
	}
	if pos.Col != 7 {
		t.Errorf("Unexpected illegal pos: %v", pos)
	}
}

func TestScanString(t *testing.T) {
	tok, _, lit := makeScanner(`"test"`).Scan()

	if tok != STRING {
		t.Errorf("Unexpected token received: %v", tok)
	}
	if lit != "test" {
		t.Errorf("Unexpected token lit received: '%v'", lit)
	}
}

func TestScanStringUnterminated(t *testing.T) {
	tok, pos, _ := makeScanner(`"test`).Scan()

	if tok != ILLEGAL {
		t.Errorf("Unexpected token received: %v", tok)
	}
	if pos.Col != 5 {
		t.Errorf("Unexpected illegal pos: %v", pos)
	}
}

func makeScanner(s string) *Scanner {
	return NewScanner(strings.NewReader(s))
}
