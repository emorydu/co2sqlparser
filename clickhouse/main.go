// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/clickhouse/parser"
)

func main() {
	template, parameters := parser.FingerprintAndTemplateExtra(`
INSERT INTO hello values('emorydu', 312432, 3124312)`)
	fmt.Println(template, parameters)
}
