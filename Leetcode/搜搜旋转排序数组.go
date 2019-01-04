package main

func main() {
	print(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

}

func search(nums []int, target int) int {
	data := make(map[int]int)
	for k, v := range nums {
		data[v] = k
	}
	if id, ok := data[target]; ok {
		return id
	} else {
		return -1
	}
}
