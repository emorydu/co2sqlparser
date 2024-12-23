// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

// PlSqlParserBase implementation.
type PlSqlParserBase struct {
	*antlr.BaseParser
	_isVersion12 bool
	_isVersion10 bool
}

func (p *PlSqlParserBase) isVersion12() bool {
	return p._isVersion12
}

func (p *PlSqlParserBase) setVersion12(value bool) {
	p._isVersion12 = value
}

func (p *PlSqlParserBase) isVersion10() bool {
	return p._isVersion10
}

func (p *PlSqlParserBase) setVersion10(value bool) {
	p._isVersion10 = value
}
