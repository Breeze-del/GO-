package tempconv

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

/*
func main() {
	for _,url:=range os.Args[1:] {
		resp,err:=http.Get(url)
		if(err!=nil) {
			fmt.Printf("fetch: %v\n",err)
			os.Exit(1)
		}
		b,err:=ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if(err!=nil) {
			fmt.Printf("fetch reading %s:%v\n",url,err)
		}
		fmt.Printf("%s",b)
	}
}
*/
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d  %s", secs, nbytes, url)
}
