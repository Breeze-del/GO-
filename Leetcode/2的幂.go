package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	var chushu, shang = n, 0
	for chushu != 1 {
		shang = chushu % 2
		chushu = chushu >> 2
		if shang != 0 {
			return false
		}
	}
	return true
}
func main() {
	n := 16
	fmt.Println(isPowerOfTwo(n))
}
