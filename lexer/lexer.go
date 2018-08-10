// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/teslamotors/jsonql/token"
)

const (
	NoState    = -1
	NumStates  = 89
	NumSymbols = 78
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '_'
1: '.'
2: '.'
3: '"'
4: '\'
5: '"'
6: '"'
7: '''
8: '\'
9: '''
10: '''
11: '~'
12: '='
13: '!'
14: '~'
15: '='
16: '+'
17: '-'
18: '*'
19: '/'
20: '%'
21: '^'
22: '!'
23: 'n'
24: 'u'
25: 'l'
26: 'l'
27: 't'
28: 'r'
29: 'u'
30: 'e'
31: 'f'
32: 'a'
33: 'l'
34: 's'
35: 'e'
36: '.'
37: '['
38: ']'
39: '_'
40: 'e'
41: 'E'
42: '+'
43: '-'
44: '0'
45: '0'
46: 'x'
47: 'X'
48: '\'
49: 'x'
50: '\'
51: 'u'
52: 'b'
53: 'f'
54: 'n'
55: 'r'
56: 't'
57: 'v'
58: '\'
59: '\'
60: ' '
61: '\t'
62: '\f'
63: '\v'
64: \u00a0
65: \u202f
66: \u205f
67: \u3000
68: \ufeff
69: 'A'-'Z'
70: 'a'-'z'
71: '0'-'9'
72: '1'-'9'
73: '0'-'7'
74: 'a'-'f'
75: 'A'-'F'
76: \u2000-\u200a
77: .
*/
