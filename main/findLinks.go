package main

import (
	"fmt"
	"strings"
)

func add1(r rune) rune {
	return r + 1
}
func main() {
	//每一个字符都调用了add1函数
	fmt.Println(strings.Map(add1, "HAL-9000"))
}
