// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/sqlite/parser"
)

func main() {
	s := `select column1, column2 as ttt from helloworld where column1 in (select numbers from hello)`
	fmt.Println(s)
	result := parser.FingerprintAndTemplateExtra(s)
	fmt.Printf("%#+v\n", result)
}
