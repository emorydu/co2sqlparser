package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/tsql/parser"
)

func main() {
	result :=  parser.FingerprintAndTemplateExtra(`
UPDATE Customers
SET ContactName='Juan'
WHERE Country='Mexico';`)
	fmt.Printf("%#+v\n", result)
}
