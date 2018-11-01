package main

import (
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"myapp1/tempconv"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lisa(w http.ResponseWriter, r *http.Request) {
	Lissajous(w)
}
func Lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "From[%q]= %q\n", k, v)
	}
	mu.Lock()
	count++
	mu.Unlock()
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func surface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	tempconv.Surface(w)
}

func draw(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: black; fill: red; stroke-width: 1.7' width='1000' height='1000'>\n")
	fmt.Fprintf(w, "<circle cx='100' cy='50' r='40' fill='yellow'/>")
	fmt.Fprintf(w, "<rect x='400' y='200' width='300' height='100' style='fill:rgb(0,0,255);stroke-width:1;stroke:rgb(0,0,0)'/>\n")
	fmt.Fprintf(w, "<ellipse cx='300' cy='80' rx='100' ry='50' style='fill:#eeeeee;stroke:purple;stroke-width:2'/>\n")
	fmt.Fprintf(w, "<line x1='300' y1='350' x2='400' y2='400' style='stroke:rgb(255,0,0);stroke-width:2'/>\n")
	fmt.Fprintf(w, "<polyline points='20,20 40,25 60,40 80,120 120,140 200,180' style='fill:none;stroke:black;stroke-width:3' />\n")
	fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g' />\n", 10.0, 210.0, 100.0, 400.0, 20.0, 300.0)
	fmt.Fprintf(w, "<path d='M 100 350 q 150 -300  300 0 q 150 -300' stroke='blue' stroke-width='5' fill='none' />")

	// 绘制控制线
	fmt.Fprintf(w, "<path d='M500 500 L500 600' stroke='#000000' fill='none' style='stroke:red;stroke-width: 2px;'/>\n")
	fmt.Fprintf(w, "<path d='M600 600 L700 500' stroke='#000000' fill='none' style='stroke:red;stroke-width: 2px;'/>\n")
	fmt.Fprintf(w, "<path d='M700 500 L800 400' stroke='#000000' fill='none' style='stroke:blue;stroke-width: 2px;'/>\n")
	fmt.Fprintf(w, "<path d='M900 500 L900 600' stroke='#000000' fill='none' style='stroke:red;stroke-width: 2px;'/>\n")
	// 绘制曲线
	fmt.Fprintf(w, "<path d='M500 500 C500 600 600 600 700 500 S900 500 900 600' stroke='#000000' fill='none' style='stroke-width: 2px;'/>\n")
	fmt.Fprintf(w, "</svg>\n")
}

func temp(w http.ResponseWriter, r *http.Request) {
	// 一定要设置这个 header，不然浏览器不会把输出的 xml 解释成 svg.
	w.Header().Set("Content-Type", "image/svg+xml")
	draw(w)
}
func jspp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, err := template.ParseFiles("F:/mygo/src/myapp1/index/svgTempory.html")
	if err != nil {
		fmt.Fprint(w, "error %s", err)
	}
	t.Execute(w, nil)

}
func main() {
	//如果不自己定义 那么会使用默认的MUX和server
	//自定义Server
	server := &http.Server{
		Addr:         ":8000",
		WriteTimeout: 2 * time.Second, //写超时2秒 函数处理时间不能超过两秒不返回
	}
	//自己写的mux路由
	mux := http.NewServeMux()
	//Handler 就是方法  路由匹配路径然后调用相应方法
	mux.Handle("/complex", &myHandler{})
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/count", counter)
	mux.HandleFunc("/lisa", lisa)
	mux.HandleFunc("/surface", surface)
	mux.HandleFunc("/draw", temp)
	mux.HandleFunc("/json", jspp)
	//文件目录
	mux.Handle("/jspp/", http.StripPrefix("/jspp/", http.FileServer(http.Dir("F:/mygo/src/myapp1/index"))))
	//绑定mux到Server上
	server.Handler = mux
	//log.Fatal(http.ListenAndServe("localhost:8000",mux))
	log.Fatal(server.ListenAndServe())
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tempconv.ComplexImage(w)
}
