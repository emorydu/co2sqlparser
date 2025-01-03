// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

// PlSqlLexerBase state
type PlSqlLexerBase struct {
	*antlr.BaseLexer

	lastToken antlr.Token
}

// NextToken from the character stream.
func (l *PlSqlLexerBase) NextToken() antlr.Token {
	next := l.BaseLexer.NextToken() // Get next token
	if next.GetChannel() == antlr.TokenDefaultChannel {
		// Keep track of the last token on default channel
		l.lastToken = next
	}
	return next
}

// IsRegexPossible returns true if the lexer can match a
// regex literal.
func (l *PlSqlLexerBase) IsNewlineAtPos(pos int) bool {
	if l.lastToken == nil {
		return true
	}
	la := l.GetInputStream().LA(pos)
	return la == -1 || la == '\n'
}
