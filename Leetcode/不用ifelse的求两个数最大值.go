package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d $d", &a, &b)
	fmt.Println(Max(5, 6))
}

func Max(a, b int) int {
	return (a + b + abs(a-b)) / 2
}

func abs(res int) int {
	if res > 0 {
		return res
	} else {
		return -res
	}
}
