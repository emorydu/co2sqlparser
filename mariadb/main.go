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
