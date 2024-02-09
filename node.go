// Copyright 2024 LangVM Project
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package migo

import scanner "cee-scanner"

type PosRange struct {
	From, To scanner.Position
}

func (p PosRange) GetPosRange() PosRange { return p }

type Token struct {
	PosRange
	Kind, Format int
	Literal      string
}

type Ident struct {
	Token
}
