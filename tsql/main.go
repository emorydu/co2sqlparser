package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/tsql/parser"
)

func main() {
	template, parameters := parser.FingerprintAndTemplateExtra(`
UPDATE Customers
SET ContactName='Juan'
WHERE Country='Mexico';`)
	fmt.Println("templates:", template)
	fmt.Println("parameters:", parameters)
}
