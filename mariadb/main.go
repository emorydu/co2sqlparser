// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/mariadb/parser"
)

func main() {
	template, parameters := parser.FingerprintAndTemplateExtra(`
INSERT INTO hello VALUES('emorydu', 1231, 321314132)`)
	fmt.Println(template, parameters)
}
