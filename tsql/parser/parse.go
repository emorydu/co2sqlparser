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
	*BaseTSqlParserListener
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

func (l *fingerprintVisitor) EnterTable_source_item(ctx *Table_source_itemContext) {
	l.tableMap[ctx.GetText()] = struct{}{}
}

//func (l *fingerprintVisitor) EnterTable_alias(ctx *Table_aliasContext) {
//	l.tableMap[ctx.GetText()] = struct{}{}
//}

func (l *fingerprintVisitor) EnterFull_column_name(ctx *Full_column_nameContext) {
	l.columnMap[ctx.GetText()] = struct{}{}
}

func (l *fingerprintVisitor) EnterCreate_table(ctx *Create_tableContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterCreate_view(ctx *Create_viewContext) {
	l.sqlType = DDL
}

func (l *fingerprintVisitor) EnterSelect_statement(ctx *Select_statementContext) {
	l.sqlType = DQL
	l.serialType = SerialSelect
}

func (l *fingerprintVisitor) EnterUpdate_statement(ctx *Update_statementContext) {
	l.sqlType = DML
	l.serialType = SerialUpdate
}
func (l *fingerprintVisitor) EnterInsert_statement(ctx *Insert_statementContext) {
	l.sqlType = DML
	l.serialType = SerialInsert
}
func (l *fingerprintVisitor) EnterDelete_statement(ctx *Delete_statementContext) {
	l.sqlType = DML
	l.serialType = SerialDelete
}

func (l *fingerprintVisitor) EnterPrimitive_constant(ctx *Primitive_constantContext) {
	l.parameters = append(l.parameters, ctx.GetText())
}

func (l *fingerprintVisitor) ExitPrimitive_constant(ctx *Primitive_constantContext) {
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
	lexer := NewTSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewTSqlParser(stream)
	tree := parser.Tsql_file()

	listener := &fingerprintVisitor{
		templates:    sql,
		numberParams: make(map[string]int),
		sqlType:      Others,
		serialType:   SerialOthers,
		columnMap:    make(map[string]struct{}),
		tableMap:     make(map[string]struct{}),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	for k := range listener.tableMap {
		listener.tables = append(listener.tables, k)
	}

	for k := range listener.columnMap {
		listener.columns = append(listener.columns, k)
	}

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
	//var output []string
	//fmt.Println(tree.ToStringTree(output, parser))

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
