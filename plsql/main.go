// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/plsql/parser"
)

func main() {
	sql := `select c1, c2, c3 from hello where username = 'emorydu'`
	result := parser.FingerprintAndTemplateExtra(sql)
	fmt.Printf("%#+v\n", result)
}
