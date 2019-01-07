package main

import (
	"fmt"
)

type windwo struct {
	left     int // 窗口的左边界
	max      int // 最长子串的长度
	activite int // 当前子串长度
}

// 窗口法 通过滑动窗口 找出最长的子串
func main() {
	var res int
	res = lengthOfLongestSubstring("tmmzuxt")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	list := make(map[int32]int)
	info := windwo{}
	info.left = 0
	info.max = 0
	info.activite = 0
	for k, v := range s {
		if locate, ok := list[v]; ok { //已经有这个元素
			if locate < info.left { //判断是否是在边界内
				list[v] = k
				info.activite += 1
				if info.activite > info.max {
					info.max = info.activite
				}
				continue
			}
			info.left = locate + 1 // 更新边界
			list[v] = k            // 更新元素的位置
			info.activite = k - info.left + 1
			// 换一次边界就要重新复制当前字符串长度，相当于一次新的开始
		} else {
			list[v] = k
			info.activite += 1
			if info.activite > info.max {
				info.max = info.activite
			}
		}
	}
	return info.max
}
