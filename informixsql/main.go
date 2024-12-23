// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/informixsql/parser"
)

func main() {
	template, parameters := parser.FingerprintAndTemplateExtra(`
RENAME COLUMN customer.customer_num TO c_num; `)
	// todo
	fmt.Println(template, parameters)
}
