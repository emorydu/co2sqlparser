// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package parser

// SqlMode represents SQL modes that control parsing behavior.

type SqlMode int

// Enum values for SqlMode
const (
	NoMode SqlMode = iota
	AnsiQuotes
	HighNotPrecedence
	PipesAsConcat
	IgnoreSpace
	NoBackslashEscapes
)

// String provides string representation for SqlMode
func (sm SqlMode) String() string {
	return [...]string{
		"NoMode",
		"AnsiQuotes",
		"HighNotPrecedence",
		"PipesAsConcat",
		"IgnoreSpace",
		"NoBackslashEscapes",
	}[sm]
}
