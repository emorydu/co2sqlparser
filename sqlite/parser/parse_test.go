package parser

import (
	"fmt"
	"testing"
)

func Test_returnMaxLenElem(t *testing.T) {
	var m = map[string]int{
		"a": 8,
		"b": 8,
		"c": 2,
	}
	count := len(m)
	for i := 0; i < count+1; i++ {
		fmt.Println(returnMaxLenElem(m))
	}

}
