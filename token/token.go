// Copyright 2024 LangVM Project
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package token

const (
	ILLEGAL int = iota

	IDENT // main

	LITERAL_BEGIN

	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"

	LITERAL_END

	KEYWORD_BEGIN

	OPERATOR_BEGIN

	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	MEMBER_SELECT //.

	LAND // &&
	LOR  // ||

	EQL // ==
	NEQ // !=
	LEQ // <=
	GEQ // >=

	LSS    // <
	GTR    // >
	ASSIGN // =

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	NOT // !

	ELLIPSIS // ...

	INC // ++
	DEC // --

	AS // as
	IN // in

	OPERATOR_END

	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	TRAIT
	MAP
	PACKAGE
	RANGE
	RETURN

	SWITCH
	SELECT
	STRUCT
	TYPE
	VAR
	VAL

	KEYWORD_END

	DELIMITER_BEGIN

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	NEWLINE   // \n

	DELIMITER_END

	token_end
)

var KeywordLiterals = [...]string{
	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	MEMBER_SELECT: ".",

	LAND: "&&",
	LOR:  "||",

	INC: "++",
	DEC: "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ:      "!=",
	LEQ:      "<=",
	GEQ:      ">=",
	ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",
	NEWLINE:   "\n",

	BREAK:    "break",
	CASE:     "case",
	CHAN:     "chan",
	CONST:    "const",
	CONTINUE: "continue",

	DEFAULT: "default",
	DEFER:   "defer",
	ELSE:    "else",
	FOR:     "for",

	FUNC:   "fun",
	GO:     "go",
	GOTO:   "goto",
	IF:     "if",
	IMPORT: "import",

	TRAIT:   "interface",
	MAP:     "map",
	PACKAGE: "package",
	RANGE:   "range",
	RETURN:  "return",

	SWITCH: "switch",
	SELECT: "select",
	STRUCT: "struct",
	TYPE:   "type",
	VAR:    "var",
	VAL:    "val",

	AS: "as",
	IN: "in",

	token_end: "",
}

func IsLiteralValue(kind int) bool { return LITERAL_BEGIN < kind && kind < LITERAL_END }

var PrefixUnaryOperators = [...]bool{
	MUL: true,
	AND: true,

	token_end: false,
}

var PostfixUnaryOperators = [...]bool{
	INC: true,
	DEC: true,

	token_end: false,
}

var BinaryOperators = [...]int{
	MUL: 1,
	QUO: 1,
	REM: 1,

	ADD: 1,
	SUB: 1,
	SHL: 1,
	SHR: 1,

	token_end: 0,
}

func IsOperator(kind int) bool { return OPERATOR_BEGIN < kind && kind < OPERATOR_END }

var KeywordEnums = map[string]int{}

func IsKeyword(term int) bool { return KEYWORD_BEGIN <= term && term <= KEYWORD_END }

var Whitespaces = map[rune]int{
	' ':  1,
	'\t': 1,
	'\r': 1,
}

var Delimiters = map[rune]int{
	'{': LBRACE,
	'}': RBRACE,
	'[': LBRACK,
	']': RBRACK,
	'(': LPAREN,
	')': RPAREN,

	',': COMMA,
	';': SEMICOLON,

	'"':  1,
	'\'': 1,

	'\n': NEWLINE, // Newline, might be a statement terminator.
}

func init() {
	for i := 0; i < token_end; i++ {
		KeywordEnums[KeywordLiterals[i]] = i
	}
}
