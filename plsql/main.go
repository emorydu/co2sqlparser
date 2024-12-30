// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/plsql/parser"
)

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
	result := parser.FingerprintAndTemplateExtra(sql)
	fmt.Printf("%#+v\n", result)
}
