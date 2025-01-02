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
	*BaseSQLiteParserListener
	templates    string
	parameters   []string
	columns      []string
	tables       []string
	sqlType      string
	serialType   int
	numberParams map[string]int
	tableMap     map[string]struct{}
	columnMap    map[string]struct{}
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

func (l *fingerprintVisitor) EnterCreate_table_stmt(ctx *Create_table_stmtContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCreate_index_stmt(ctx *Create_index_stmtContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCreate_trigger_stmt(ctx *Create_trigger_stmtContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCreate_view_stmt(ctx *Create_view_stmtContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterDrop_stmt(ctx *Drop_stmtContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterAlter_table_stmt(ctx *Alter_table_stmtContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterBegin_stmt(ctx *Begin_stmtContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterCommit_stmt(ctx *Commit_stmtContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterRollback_stmt(ctx *Rollback_stmtContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterSavepoint_stmt(ctx *Savepoint_stmtContext) {
	l.sqlType = TCL
}

func (l *fingerprintVisitor) EnterSelect_stmt(ctx *Select_stmtContext) {
	l.sqlType = DQL
	l.serialType = SerialSelect
}

func (l *fingerprintVisitor) EnterUpdate_stmt(ctx *Update_stmtContext) {
	l.sqlType = DML
	l.serialType = SerialUpdate
}

func (l *fingerprintVisitor) EnterInsert_stmt(ctx *Insert_stmtContext) {
	l.sqlType = DML
	l.serialType = SerialInsert
}

func (l *fingerprintVisitor) EnterDelete_stmt(ctx *Delete_stmtContext) {
	l.sqlType = DML
	l.serialType = SerialDelete
}

func (l *fingerprintVisitor) EnterLiteral_value(ctx *Literal_valueContext) {
	l.parameters = append(l.parameters, ctx.GetText())
}

func (l *fingerprintVisitor) EnterColumn_name(ctx *Column_nameContext) {
	l.columnMap[ctx.GetText()] = struct{}{}
}

func (l *fingerprintVisitor) ExitLiteral_value(ctx *Literal_valueContext) {
	originalText := ctx.GetText()
	_, err := strconv.Atoi(originalText)
	if err != nil {
		l.templates = strings.Replace(l.templates, originalText, "?", 1)
	} else {
		l.numberParams[ctx.GetText()] = len(ctx.GetText())
	}
}

func (l *fingerprintVisitor) EnterTable_name(ctx *Table_nameContext) {
	l.tableMap[ctx.GetText()] = struct{}{}
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
	lexer := NewSQLiteLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewSQLiteParser(stream)
	tree := parser.Sql_stmt()

	listener := &fingerprintVisitor{
		templates:    sql,
		numberParams: make(map[string]int),
		sqlType:      Others,
		serialType:   SerialOthers,
		tableMap:     make(map[string]struct{}),
		columnMap:    make(map[string]struct{}),
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

	for k := range listener.tableMap {
		listener.tables = append(listener.tables, k)
	}

	for k := range listener.columnMap {
		listener.columns = append(listener.columns, k)
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
