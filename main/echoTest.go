package main
//
//import (
//	"flag"
//	"fmt"
//)
//
//var n = flag.Bool("n", false, "omit trailing newline")
//var sep = flag.String("s", "", "separator")
//func main() {
//	/*
//	flag.Parse()
//	fmt.Print(strings.Join(flag.Args(),*sep))
//	if !*n {
//		fmt.Println()
//	}
//	*/
//	var s *[]int
//	s =fib(50)
//	fmt.Println(*s)
//}
////最大公约数求解   元组赋值可以完成两个数组的交换 x,y=y,x
//func gcd(x, y int) int {
//	for y != 0 {
//		x, y = y, x%y
//	}
//	return y;
//}
//
////Fibonacci数列
//func fib(n int) *[]int {
//	data := make([]int, 100)
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		data[i]=x
//		data[i+1]=y
//		x, y = y, x+y
//	}
//	return &data
//}