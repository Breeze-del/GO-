package main

import (
	"io/ioutil"
	"math"
	"os"
)

type Pointd struct {
	x, y float64
}
type Path []Pointd

func main() {
	//defer println("我最后执行")
	//println(sum(1,2,3,4,8,7,10))
	//var filnaem string  ="D:/Pythonfile/venv/baikeSpider/output.txt"
	// bb, err := ReadFile(filnaem)
	// if err != nil {
	//	 println("failed")
	// }
	//println(string(bb))

	//var path Path =Path{
	//	{1,5},{2,4},{3,6},{4,5},
	//}
	//println(path.Distance())

	//var str [][]string
	//var a []string
	//var b []string
	//a=append(a,"aaa")
	//a=append(a,"ccc")
	//str=append(str, a)
	//b=append(b,"qqq")
	//b=append(b,"www")
	//str=append(str,b)
	//for k := range str {
	//	for j := range str[k] {
	//		println(str[k][j])
	//	}
	//}
}

//D:\Pythonfile\venv\baikeSpider\output.txt
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func (p Pointd) Distance(q Pointd) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
