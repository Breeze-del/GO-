package tempconv

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320 //以像素表示画布的大小
	cells=100   //网络单元格数目
	xyrange = 30.0  //坐标轴的范围
	xyscale =width/2/xyrange //坐标轴上每个单位的 长度的像素
	zxcale = height * 0.4 //z 轴上每个单位长度的像素
	angle = math.Pi/6  //x，y轴上的角度
)

var sin30,cos30 = math.Sin(angle), math.Cos(angle)

func Surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: black; fill: red; stroke-width: 1.7' width='600' height='320'>\n")
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := Corner(i+1,j)
			bx, by := Corner(i, j)
			cx, cy := Corner(i, j+1)
			dx, dy := Corner(i+1,j+1)
			fmt.Fprintf(out,"<polygon points=%g,%g,%g,%g,%g,%g,%g,%g>\n",
				ax,ay,bx,by,cx,cy,dx,dy)
		}
	}
	fmt.Fprint(out,"</svg>\n")
}

func Corner(i, j int) (float64, float64) {
	//求出网格单元（i，j）的顶点坐标（x，y）
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//计算曲面的高度
	z := f(x,y)
	//将x,y,z投影到二维SVG的绘图平面上，坐标是(sx,xy)
	sx := width/2+(x-y)*cos30*xyrange
	sy := height/2 + (x+y)*sin30*xyrange - z*zxcale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //点到0，0的距离 依次计算高度
	return math.Sin(r) / r
}