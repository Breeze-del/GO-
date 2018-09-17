package main

import (
	"bytes"
	"fmt"
	"myapp1/tempconv"
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
	//fmt.Println(intsToString([]int{1, 2, 3}))
	//var str []string = []string{"abcdefgddddddddd", "ccccc"}
	//for key, value := range str {
	//	str[key] = "cccccccccc"
	//	fmt.Println(key, value)
	//}
	//fmt.Println(str)
	////stirng 切片还是string  并不是【】strings
	//ss := "aaaaaaaaaa"
	//var sl []string = []string{ss[:4]}
	//var sll []string
	//sll = sl
	//fmt.Println(sll, sl)

	//实验Stack  切片的赋值需要range 赋值
	//fc是实现stack方法的接口  stack只是一个当做任意类型的数据的口接口
	var fc tempconv.Function
	stack := make(tempconv.Stack, 4)
	t := []int{1, 2, 3, 4}
	for i, v := range t {
		stack[i] = v
		fmt.Println(i, v)
	}
	fc = &stack
	tp, err := fc.Top()
	fmt.Println(tp, err)
}
