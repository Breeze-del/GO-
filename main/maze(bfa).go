package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := read("1.txt")
	ss := walk(arr, point{0, 0}, point{len(arr) - 1, len(arr[0]) - 1})
	for _, v := range ss {
		for _, vv := range v {
			fmt.Print(vv, "  ")
		}
		fmt.Println()
	}
}

var dirs = [4]point{ //左下右上
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

//定义一个结构体数组
type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}
func walk(maze [][]int, start, end point) [][]int {
	//创建一个离原点多远可以到达的二维数组
	steps := make([][]int, len(maze))
	for k, _ := range steps {
		steps[k] = make([]int, len(maze[k]))
	}

	//将头放进队列
	Q := []point{start}
	for len(Q) > 0 {
		//取出第一个point结构体 然后将 第一个弹出队列
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			//到达目的地
			break
		}
		//发现四个点
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			//val ==1 说明撞墙 ok为false说明超出边界
			if !ok || val == 1 {
				continue
			}
			//走另一张图
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			//是否回到原点，因为maze和steps图中原点的值都是0
			if next == start {
				continue
			}
			//可以走了
			i, _ := cur.at(steps)
			steps[next.i][next.j] = i + 1
			//将该点加入队列继续找
			Q = append(Q, next)
		}
	}
	return steps
}

//边界判断
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	//获取迷宫的值 如果返回为1 说明撞墙了
	return grid[p.i][p.j], true
}
func read(filename string) [][]int {
	f, e := os.Open(filename)
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	buf := make([]byte, 2048)
	n, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	str := string(buf[:n])
	//以换行符分割所哟字符串
	split := strings.Split(str, "\n")

	//获取列和行
	var row, col int
	//去掉字符串前后端所有的空白
	ss := strings.TrimSpace(split[0])
	row, err = strconv.Atoi(strings.Split(ss, " ")[0])
	if err != nil {
		//抛出异常 正常程序终止，然后执行defer语句 再报告异常信息 最后退出
		//recover 异常捕获函数  可以捕获异常 使程序继续正常执行下去，只能再defer中返回异常
		panic(err)
	}
	col, err = strconv.Atoi(strings.Split(ss, " ")[1])
	fmt.Println(row, col)
	if err != nil {
		panic(err)
	}
	//创建行数 然后再创建列数
	arr := make([][]int, row)
	for k, _ := range arr {
		sArr := strings.Split(strings.TrimSpace(split[k+1]), " ")
		arr[k] = make([]int, col)
		for kk, _ := range arr[k] {
			i, err := strconv.Atoi(sArr[kk])
			if err != nil {
				panic(err)
			}
			arr[k][kk] = i
		}
	}
	return arr
}
