// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/mysql/parser"
)

func main() {
	//	template, parameters := parser.FingerprintAndTemplateExtra(`
	//insert into hello (name, age, zipc25ode) values ('emorydu', -25, 0)`)

	result := parser.FingerprintAndTemplateExtra(`
select column1, column2 from helloworld where column1 = 'emorydu'`)

	fmt.Printf("%#+v\n", result.Columns)
	//
	//template, parameters := parser.FingerprintAndTemplateExtra(`
	//CREATE TABLE table_name (
	// column1 varchar(50),
	// column2 int,
	// column3 timestamp,
	//	CONSTRAINT check_salary CHECK (salary >= 0)
	//);`)

}
