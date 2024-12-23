// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	plsql "github.com/emorydu/co2sqlparser/plsql/parser"
	"regexp"
	"strconv"
	"strings"
)

type fingerprintVisitor struct {
	*plsql.BasePlSqlParserListener
	templates  string
	parameters []string
}

func (l *fingerprintVisitor) EnterConstant(ctx *plsql.ConstantContext) {
	l.parameters = append(l.parameters, ctx.GetText())
}

func (l *fingerprintVisitor) ExitConstant(ctx *plsql.ConstantContext) {
	originalText := ctx.GetText()
	_, err := strconv.Atoi(originalText)
	if err != nil {
		l.templates = strings.Replace(l.templates, originalText, "?", 1)
	} else {
		re := regexp.MustCompile(`\b` + regexp.QuoteMeta(originalText) + `\b`)
		l.templates = re.ReplaceAllString(l.templates, "?")
	}
}

func main() {

	sql := `CREATE TABLE table_name (
	column1 varchar(50),
	column2 int,
	column3 timestamp,
		CONSTRAINT check_salary CHECK (salary >= 0)
	);`

	is := antlr.NewInputStream(sql)
	lexer := plsql.NewPlSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := plsql.NewPlSqlParser(stream)

	tree := parser.Unit_statement()

	listener := &fingerprintVisitor{
		templates: sql,
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	fmt.Println("templates:", listener.templates)
	fmt.Println("parameters:", listener.parameters)

	var output []string
	fmt.Println(tree.ToStringTree(output, parser))

}
