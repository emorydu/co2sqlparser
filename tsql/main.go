package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/tsql/parser"
)

func main() {
	result := parser.FingerprintAndTemplateExtra(`
select c1, c2, c3 from helloworld where username = 'emorydu'`)
	fmt.Printf("%#+v\n", result)
}
