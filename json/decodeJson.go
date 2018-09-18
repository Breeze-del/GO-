package json

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//访问 https：//omdbapi.com 解析json数据然后下载图片 根据电影名字下载图片

//网址
const MovieURL = "https://xkcd.com/"

//定义数据结构
type Movie struct {
	Month     string
	Year      string
	SafeTitle string `json:"safe_title"`
}

type Student struct {
	Name string
	Age  int
	High bool
}

type Mm struct {
	Title  string
	Year   int
	Color  bool
	Actors []string
}

//Decode 方式
func SearchMovie() (*Movie, error) {
	resp, err := http.Get(MovieURL + "571/info.0.json")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed : &s", resp.Status)

	}
	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//unmarshal 方式解析
func Unmarshal() (*[]Mm, error) {
	//resp, err := http.Get(MovieURL+"571/info.0.json")
	var student []Mm
	data, err := Stoj()
	if err != nil {
		return nil, nil
	}
	if err := json.Unmarshal(*data, &student); err != nil {
		log.Fatal("json unmarshaling failed : %s", err)
	}
	return &student, nil
}

//将结构体转换为json数据
var mm = []Mm{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func Stoj() (*[]byte, error) {
	data, err := json.Marshal(mm)
	if err != nil {
		log.Fatal("JSON marshaling failed : %s", err)
	}
	//fmt.Printf("%s\n", data)
	return &data, nil
}
