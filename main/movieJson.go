package main

import (
	"fmt"
	"log"
	"myapp1/json"
)

func main() {
	//Decode 解析方式
	result, err := json.SearchMovie()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s %s", result.Month, result.Year, result.SafeTitle)
	fmt.Println()
	//unmarshal 解析方式
	stu, err := json.Unmarshal()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range *stu {
		fmt.Println(k, v)
	}
	//结构体转josn
	mm, err := json.Stoj()
	fmt.Printf("%s", *mm)
}
