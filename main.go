// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	plsql "github.com/emorydu/co2sqlparser/plsql/parser"
	"strings"
)

type fingerprintVisitor struct {
	*plsql.BasePlSqlParserListener
	templates  string
	parameters []string
}

func (l *fingerprintVisitor) EnterConstant(ctx *plsql.ConstantContext) {
	fmt.Println("++++++++:", ctx.GetText())
	l.parameters = append(l.parameters, ctx.GetText())
}

func (l *fingerprintVisitor) ExitConstant(ctx *plsql.ConstantContext) {
	originalText := ctx.GetText()
	l.templates = strings.Replace(l.templates, originalText, "?", 1)
}

func main() {
	sql := `
	CREATE TABLE employees (
	employee_id NUMBER PRIMARY KEY,
	first_name VARCHAR2(50) NOT NULL,
	last_name VARCHAR2(50) NOT NULL,
	hire_date DATE,
	salary NUMBER(10, 2),
	CONSTRAINT unique_first_name UNIQUE (first_name),
	CONSTRAINT check_salary CHECK (salary >= 0)
	);`
	//sql := `select * from employee
	//START WITH mangerid = 'king'
	// connect by prior empid = mangerid`
	//sql := `select * from hello where name = 'emorydu' and age >= 20 and price <= 20`
	//sql := `-- INSERT INTO helloworld VALUES (10001, 'Emory.Du', 23, 233423);`
	//	sql := `SELECT LEVEL, employee_id, first_name
	//FROM employees
	//CONNECT BY PRIOR employee_id = manager_id
	//START WITH manager_id IS NULL;`
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

	//select * from hello where name = 'emorydu' and age >= 20

}
