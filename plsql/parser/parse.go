// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"regexp"
	"strings"
)

type fingerprintVisitor struct {
	*BasePlSqlParserListener
	templates  string
	parameters []string
}

func (l *fingerprintVisitor) EnterConstant(ctx *ConstantContext) {
	l.parameters = append(l.parameters, ctx.GetText())
}

func (l *fingerprintVisitor) ExitConstant(ctx *ConstantContext) {
	originalText := ctx.GetText()
	re := regexp.MustCompile(`\b` + regexp.QuoteMeta(originalText) + `\b`)
	l.templates = re.ReplaceAllString(l.templates, "?")
	//l.templates = strings.Replace(l.templates, originalText, "?", 1)
}

func FingerprintAndTemplateExtra(sql string) (template string, parameters []string) {
	is := antlr.NewInputStream(sql)
	lexer := NewPlSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewPlSqlParser(stream)
	tree := parser.Unit_statement()

	listener := &fingerprintVisitor{
		templates: sql,
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	template = strings.Trim(listener.templates, "\n")
	parameters = listener.parameters
	return
}
