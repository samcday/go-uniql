package uniql

import (
	"bufio"
	"bytes"
	"io"
)

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader

	line    int
	col     int
	prevCh  rune
	prevCol int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (Token, Pos, string) {
	ch, pos := s.read()

	if ch == eof {
		return EOF, pos, ""
	}

	if isWhitespace(ch) {
		s.unread()
		tok, lit := s.scanWhitespace()
		return tok, pos, lit
	} else if isLetter(ch) || ch == '_' {
		s.unread()
		tok, lit := s.scanIdentifier()
		return tok, pos, lit
	} else if isDigit(ch) || ch == '.' {
		s.unread()
		return s.scanNumber()
	} else if ch == '"' {
		s.unread()
		return s.scanString()
	}

	switch ch {
	case '=':
		if ch2, _ := s.read(); ch2 == '=' {
			return EQ, pos, ""
		}
		s.unread()
	case '(':
		return LPAREN, pos, ""
	case ')':
		return RPAREN, pos, ""
	case '!':
		if ch2, _ := s.read(); ch2 == '=' {
			return NEQ, pos, ""
		}
		s.unread()
	case '>':
		if ch2, _ := s.read(); ch2 == '=' {
			return GTE, pos, ""
		}
		s.unread()
		return GT, pos, ""
	case '<':
		if ch2, _ := s.read(); ch2 == '=' {
			return LTE, pos, ""
		}
		s.unread()
		return LT, pos, ""
	case '~':
		if ch2, _ := s.read(); ch2 == '=' {
			return REG_MATCH, pos, ""
		}
		s.unread()
	}

	return ILLEGAL, pos, string(ch)
}

func (s *Scanner) scanWhitespace() (Token, string) {
	var buf bytes.Buffer

	for {
		if ch, _ := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WHITESPACE, buf.String()
}

func (s *Scanner) scanIdentifier() (Token, string) {
	var buf bytes.Buffer

	for {
		if ch, _ := s.read(); ch == eof {
			break
		} else if !isLegalIdentChar(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	ident := buf.String()

	if tok := LookupToken(ident); tok != IDENTIFIER {
		return tok, ""
	}

	return IDENTIFIER, ident
}

func (s *Scanner) scanNumber() (Token, Pos, string) {
	var buf bytes.Buffer
	fp := false

	_, firstPos := s.read()
	s.unread()

	for {
		ch, pos := s.read()
		if ch == eof {
			break
		} else if ch == '.' {
			// Can't have multiple floating points...
			if fp {
				return ILLEGAL, pos, string(ch)
			}

			// Period must be followed by a digit.
			ch2, pos2 := s.read()
			if !isDigit(ch2) {
				return ILLEGAL, pos2, string(ch2)
			}

			fp = true
			s.unread()
		} else if !isDigit(ch) {
			s.unread()
			break
		}

		buf.WriteRune(ch)
	}

	return NUMBER, firstPos, buf.String()
}

func (s *Scanner) scanString() (Token, Pos, string) {
	var buf bytes.Buffer

	_, firstPos := s.read()

	for {
		if ch, pos := s.read(); ch == eof {
			// Un-terminated string literal...
			return ILLEGAL, pos, ""
		} else if ch == '"' {
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return STRING, firstPos, buf.String()
}

func (s *Scanner) read() (rune, Pos) {
	ch, _, err := s.r.ReadRune()
	s.prevCh = ch

	if err != nil {
		return eof, Pos{s.line, s.col}
	}

	line, col := s.line, s.col

	if ch == '\n' {
		s.prevCol = s.col
		s.line++
		s.col = 0
	} else {
		s.col++
	}

	return ch, Pos{line, col}
}

func (s *Scanner) unread() {
	// Rewind position if necessary
	if s.prevCh == '\n' {
		s.line--
		s.col = s.prevCol
	} else {
		s.col--
	}

	_ = s.r.UnreadRune()
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// Returns true if char is legal in an identifier (EXCLUDING the starting char, which must be [A-Za-z_])
func isLegalIdentChar(ch rune) bool {
	return isLetter(ch) || isDigit(ch) || ch == '_' || ch == '.'
}
