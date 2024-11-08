package scanner

import (
	"fmt"
	"glorp/token"
)

type Scanner struct {
	// When does a scanner have an erorr?
	// Do we ever?
	// Unknown token, prob ident?
	Source  string
	Current int
	Start   int
	Line    int
}

func NewScanner() *Scanner {
	return &Scanner{
		Current: 0,
		Start:   0,
		Line:    0,
	}
}

func (s *Scanner) Scan(source string) ([]token.Token, error) {
	s.Source = source
	var tokens []token.Token
	// While we are not at the end of the line, meaning current ptr < len(source)
	for !s.isAtEnd() {
		// Set start to current
		s.Start = s.Current
		tok, err := s.tokenize()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, tok)
	}

	eofTok := token.NewToken(token.EOF, "", s.Line)
	tokens = append(tokens, *eofTok)
	return tokens, nil
}

func (s *Scanner) tokenize() (token.Token, error) {
	// Look at char, match single char chars first
	// Then see if the next chars line up with what would be expected
	c := s.advance()

	switch c {
	case '+':
		fmt.Println("Plus")
		return s.addToken(token.PLUS), nil
	case '=':
		if s.next('=') {
			fmt.Println("Double Equals")
			return s.addToken(token.EQUAL_EQUAL), nil
		} else {
			fmt.Println("Assignment Equal")
			return s.addToken(token.EQUAL), nil
		}
	default:
		fmt.Println("Other char")
		return s.addToken(token.IDENTIFIER), nil
	}
}

func (s *Scanner) addToken(tokType token.TokenType) token.Token {
	lex := s.Source[s.Start:s.Current]
	return *token.NewToken(tokType, lex, s.Line)
}

func (s *Scanner) advance() rune {
	if s.isAtEnd() { return '0' } // Temp idk
	// Update cur ptr
	s.Current++
	// Return old char, because how would we return i=0
	return rune(s.Source[s.Current-1])
}

func (s *Scanner) next(val byte) bool {
	// How do we know that = isnt the last char in a source?
	// We have to check 2 ahead
	if s.Current+1 >= len(s.Source) { return false }
	if s.Source[s.Current+1] == val { return true } 
	return false // With range but val isnt the next char
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

// source
// var x = 10;
// start = current
// while in alphanumeric,
// consume char
// current++
//
