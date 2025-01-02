package main

import (
	"fmt"
	"github.com/emorydu/co2sqlparser/tsql/parser"
)

func main() {
	result := parser.FingerprintAndTemplateExtra(`
select numbers from hello as t1 left join world as t2 on t1.column1 = t2.column1 where t2.username = 'emorydu' and t1.password like '%1234%'`)
	fmt.Printf("%#+v\n", result)
}
