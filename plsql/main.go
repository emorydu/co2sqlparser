// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/plsql/parser"
)

func main() {
	sql := `select numbers from hello t1 left join world t2 on t1.column1 = t2.column1 where t2.username = 'emorydu' and t1.password like '%1234%'`
	result := parser.FingerprintAndTemplateExtra(sql)
	fmt.Printf("%#+v\n", result)
}
