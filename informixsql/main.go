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
create table cti_vccinfo(
    vccid CHAR(6) not null,
    vccname VARCHAR(255),
    effective INTEGER default 0 not null,
    agentmax INTEGER default 0 not null,
    ivrmax INTEGER default 0 not null,
    updatekey VARCHAR(30),
    primary key (vccid) constraint PK_CTI_VI
);`)
	// todo
	fmt.Println(template, parameters)
}
