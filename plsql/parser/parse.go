// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"regexp"
	"strconv"
	"strings"
)

type fingerprintVisitor struct {
	*BasePlSqlParserListener
	templates    string
	parameters   []string
	columns      []string
	tables       []string
	sqlType      string
	serialType   int
	numberParams map[string]int
}

const (
	DDL    = "DDL"
	TCL    = "TCL"
	DQL    = "DQL"
	DML    = "DML"
	Others = "Others"
)

const (
	SerialSelect = iota + 1
	SerialInsert
	SerialUpdate
	SerialDelete
	SerialOthers
)

func (l *fingerprintVisitor) EnterCreate_table(ctx *Create_tableContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCreate_view(ctx *Create_viewContext) {
	l.sqlType = DDL
}
func (l *fingerprintVisitor) EnterCreate_index(ctx *Create_indexContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCreate_trigger(ctx *Create_triggerContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterDrop_operator(ctx *Drop_operatorContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterAlter_table(ctx *Alter_tableContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCommit_statement(ctx *Commit_statementContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterRollback_statement(ctx *Rollback_statementContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterSavepoint_statement(ctx *Savepoint_statementContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterSelect_statement(ctx *Select_statementContext) {
	l.sqlType = DQL
	l.serialType = SerialSelect
}

func (l *fingerprintVisitor) EnterInsert_statement(ctx *Insert_statementContext) {
	l.sqlType = DML
	l.serialType = SerialInsert
}

func (l *fingerprintVisitor) EnterUpdate_statement(ctx *Update_statementContext) {
	l.sqlType = DML
	l.serialType = SerialUpdate
}

func (l *fingerprintVisitor) EnterDelete_statement(ctx *Delete_statementContext) {
	l.sqlType = DML
	l.serialType = SerialDelete
}

func (l *fingerprintVisitor) EnterConstant(ctx *ConstantContext) {
	l.parameters = append(l.parameters, ctx.GetText())
}

func (l *fingerprintVisitor) ExitConstant(ctx *ConstantContext) {
	originalText := ctx.GetText()
	_, err := strconv.Atoi(originalText)
	if err != nil {
		l.templates = strings.Replace(l.templates, originalText, "?", 1)
	} else {
		l.numberParams[ctx.GetText()] = len(ctx.GetText())
		//re := regexp.MustCompile(`\b` + regexp.QuoteMeta(originalText) + `\b`)
		//l.templates = re.ReplaceAllString(l.templates, "?")
	}
}

type Result struct {
	Template   string
	Parameters []string
	SQLType    string
	SerialType int
	Tables     []string
	Columns    []string
}

func FingerprintAndTemplateExtra(sql string) Result {
	is := antlr.NewInputStream(sql)
	lexer := NewPlSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewPlSqlParser(stream)
	tree := parser.Unit_statement()

	listener := &fingerprintVisitor{
		templates:    sql,
		numberParams: make(map[string]int),
		sqlType:      Others,
		serialType:   SerialOthers,
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	counts := len(listener.numberParams)
	for i := 0; i < counts+1; i++ {
		vv := returnMaxLenElem(listener.numberParams)
		if vv == "" {
			continue
		}
		re := regexp.MustCompile(`\b` + regexp.QuoteMeta(vv) + `\b`)
		listener.templates = re.ReplaceAllString(listener.templates, "?")
	}

	result := Result{
		Template:   strings.Trim(listener.templates, "\n"),
		Parameters: listener.parameters,
		SerialType: listener.serialType,
		SQLType:    listener.sqlType,
		Tables:     listener.tables,
		Columns:    listener.columns,
	}

	return result
}

func returnMaxLenElem(m map[string]int) string {
	mmv := ""
	maxValue := 0
	for k, v := range m {
		if v > maxValue {
			maxValue = v
			mmv = k
		}
	}

	delete(m, mmv)
	return mmv
}
