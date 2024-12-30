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
	*BaseMySQLParserListener
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

func (l *fingerprintVisitor) EnterShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {
	l.sqlType = DDL
}
func (l *fingerprintVisitor) EnterShowCreateViewStatement(ctx *ShowCreateViewStatementContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCreateIndex(ctx *CreateIndexContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCreateTrigger(ctx *CreateTriggerContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterDropStatement(ctx *DropStatementContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterAlterTable(ctx *AlterTableContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterBeginWork(ctx *BeginWorkContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterSavepointStatement(ctx *SavepointStatementContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterLockStatement(ctx *LockStatementContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterSelectStatement(ctx *SelectStatementContext) {
	l.sqlType = DQL
	l.serialType = SerialSelect
}

func (l *fingerprintVisitor) EnterUpdateStatement(ctx *UpdateStatementContext) {
	l.sqlType = DML
	l.serialType = SerialUpdate
}

func (l *fingerprintVisitor) EnterInsertStatement(ctx *InsertStatementContext) {
	l.sqlType = DML
	l.serialType = SerialInsert
}

func (l *fingerprintVisitor) EnterDeleteStatement(ctx *DeleteStatementContext) {
	l.sqlType = DML
	l.serialType = SerialDelete
}

func (l *fingerprintVisitor) EnterLiteralOrNull(ctx *LiteralOrNullContext) {
	l.parameters = append(l.parameters, ctx.GetText())
}

func (l *fingerprintVisitor) ExitLiteralOrNull(ctx *LiteralOrNullContext) {
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
	lexer := NewMySQLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewMySQLParser(stream)
	tree := parser.SimpleStatement()

	listener := &fingerprintVisitor{
		templates:    sql,
		sqlType:      Others,
		numberParams: make(map[string]int),
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
