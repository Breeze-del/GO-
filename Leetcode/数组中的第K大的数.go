package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	var res int
	res = findKthLargest(data, 1)
	fmt.Println(res)
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	index := len(nums) - k
	return nums[index]
}
