// Copyright 2024 LangVM Project
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package parser

import (
	scanner "github.com/langvm/cee-scanner"
	. "migo/ast"
	. "migo/token"
)

type Parser struct {
	scanner.Scanner

	ReachedEOF bool

	Token Token
}

func NewParser() Parser {
	p := Parser{
		Scanner: scanner.Scanner{
			Whitespaces: map[rune]int{
				' ':  1,
				'\t': 1,
				'\r': 1,
				'\n': 1,
			},
			Delimiters: map[rune]int{
				'(': 1,
				')': 1,
			},
		},
	}
	return p
}

func (p *Parser) Scan() error {
	begin, kind, format, litRunes, err := p.Scanner.ScanToken()
	switch err := err.(type) {
	case nil:
	case scanner.EOFError:
		if p.ReachedEOF {
			return err
		} else {
			p.ReachedEOF = true
			return nil
		}
	default:
		return err
	}

	lit := string(litRunes)

	switch kind {
	case scanner.IDENT:
		kind = IDENT
		if k := KeywordEnums[lit]; k != 0 {
			kind = k
		}
	case scanner.MARK:
		if k := KeywordEnums[lit]; k != 0 {
			kind = k
		} else {
			return UnknownOperatorError{p.Token}
		}
	case scanner.CHAR:
		kind = CHAR
	case scanner.STRING:
		kind = STRING
	case scanner.INT:
		kind = INT
	case scanner.FLOAT:
		kind = FLOAT
	case scanner.COMMENT:
		return p.Scan()
	default:
		panic("impossible")
	}

	p.Token = Token{
		PosRange: PosRange{From: begin, To: p.Position},
		Kind:     kind,
		Format:   format,
		Literal:  lit,
	}

	return nil
}

func (p *Parser) MatchTerm(term int) error {
	if p.Token.Kind != term {
		return UnexpectedToken{p.Token}
	}

	err := p.Scan()
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) ExpectLiteralValue() (LiteralValue, error) {
	t := p.Token

	err := p.Scan()
	if err != nil {
		return LiteralValue{}, err
	}

	return LiteralValue{t}, nil
}

func (p *Parser) ExpectIdent() (Ident, error) {
	if p.Token.Kind != IDENT {
		return Ident{}, UnexpectedToken{p.Token}
	}

	t := p.Token

	err := p.Scan()
	if err != nil {
		return Ident{}, err
	}

	return Ident{t}, nil
}
