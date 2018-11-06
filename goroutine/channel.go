package main

import "fmt"

// make(chan int) 无缓冲的channel 又叫同步channel
func main() {
	fmt.Println("aaaaa")
	natuals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			natuals <- x
		}
	}()

	// Squarer
	go func() {
		for x := range natuals {
			squares <- x * x
		}
	}()

	// Print
	for x := range squares {
		fmt.Println(x)
	}
}
