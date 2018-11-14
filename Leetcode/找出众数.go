package main

import "fmt"

func majorityElement(nums []int) int {
	datas := make(map[int]int)
	for _, v := range nums {
		datas[v] += 1
	}
	max := 0
	tp := 0
	for i := range datas {
		if datas[i] > max {
			max = datas[i]
			tp = i
		}
	}
	return tp
}
func main() {
	nums := []int{5, 5, 3, 3, 2, 2, 5}
	fmt.Print(majorityElement(nums))
}
