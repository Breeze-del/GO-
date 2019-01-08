package main

import (
	"fmt"
	"sort"
)

type ints []int

func main() {
	var list = ints{4, 3, 25, 1, 6}
	sort.Sort(list)
	fmt.Println(sort.IntsAreSorted(list))
	// 倒序输出 改变序列号 实现倒序输出
	// sort.Sort(sort.Reverse(list))
	sort.Sort(list)
	fmt.Println(list)
}

// 自定义类型实现sort排序 需要实现下面三个方法 Len Less Swap
// 注意大写字母，包级别名，可以导出
func (p ints) Len() int {
	return len(p)
}
func (p ints) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p ints) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
