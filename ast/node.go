// Copyright 2024 LangVM Project
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package ast

import scanner "github.com/langvm/cee-scanner"

type PosRange struct {
	From, To scanner.Position
}

func (p PosRange) GetPosRange() PosRange { return p }

type Node interface {
	GetPosRange() PosRange
}

type Token struct {
	PosRange
	Kind    int
	Literal string
}

type (
	Type interface {
		Node
	}

	StructType struct {
		PosRange
		Fields []GenDecl
	}

	TraitType struct {
		PosRange
		// TODO
	}

	TypeAlias struct {
		Ident
	}

	FuncType struct {
		PosRange
		Params  []GenDecl
		Results []Type
	}
)

type (
	Expr interface {
		Node
	}

	BadExpr struct {
		PosRange
	}

	LiteralValue struct {
		Token
	}

	Ident struct {
		Token
	}

	UnaryExpr struct {
		PosRange
		Operator Token
		Expr     Expr
	}

	BinaryExpr struct {
		PosRange
		Operator     Token
		ExprL, ExprR Expr
	}

	CallExpr struct {
		PosRange
		Callee Expr
		Params []Expr
	}

	IndexExpr struct {
		PosRange
		Expr  Expr
		Index Expr
	}

	CastExpr struct {
		PosRange
	}

	BranchExpr struct {
		PosRange
		Cond       Expr
		Branch     StmtBlockExpr
		ElseBranch StmtBlockExpr
	}

	MatchExpr struct {
		PosRange
		Subject  Expr
		Patterns []StmtBlockExpr
	}

	UnwrapExpr struct {
		PosRange
		Expr Expr
	}

	StmtBlockExpr struct {
		PosRange
		Type  Type // nil for void
		Stmts []Stmt
		Value Expr // optional
	}

	MemberSelectExpr struct {
		PosRange
		Member Ident
		Expr   Expr
	}
)

type (
	Stmt interface {
		Node
	}

	ImportDecl struct {
		PosRange
		CanonicalName LiteralValue
		Alias         *Ident
	}

	ValDecl struct {
		PosRange
		Name  Ident
		Value Expr
	}

	GenDecl struct {
		PosRange
		Idents []Ident
		Type   Type
	}

	FuncDecl struct {
		PosRange
		Type  FuncType
		Ident *Ident
		Stmt  *StmtBlockExpr
	}

	ReturnStmt struct {
		PosRange
		Exprs []Expr
	}

	BreakStmt struct {
		PosRange
	}

	ContinueStmt struct {
		PosRange
	}

	LoopStmt struct {
		PosRange
		Cond Expr
		Stmt StmtBlockExpr
	}

	ForeachStmt struct {
		PosRange
		IdentList []Ident
		Expr      StmtBlockExpr
	}
)
