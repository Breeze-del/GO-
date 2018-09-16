package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
	var str []string = []string{"abcdefgddddddddd", "ccccc"}
	for key, value := range str {
		str[key] = "cccccccccc"
		fmt.Println(key, value)
	}
	fmt.Println(str)
	//stirng 切片还是string  并不是【】strings
	ss := "aaaaaaaaaa"
	var sl []string = []string{ss[:4]}
	var sll []string
	sll = sl
	fmt.Println(sll, sl)
}
