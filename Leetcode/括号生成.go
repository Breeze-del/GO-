package main

import (
	"fmt"
)

// LeetCode 括号生成题
func main() {
	var res []string
	// n为括号的数目
	res = generateParenthesis(2)
	for _, v := range res {
		fmt.Println(v)
	}
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	generate(n, n, "", &res)
	return res
}

func generate(left, right int, result string, answer *[]string) {
	// left表示左边剩下的括号数
	// right表示右边剩下的括号数
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		*answer = append(*answer, result)
	} else {
		if left > 0 {
			generate(left-1, right, result+"(", answer)
		}
		if right > 0 {
			generate(left, right-1, result+")", answer)
		}
	}
}
