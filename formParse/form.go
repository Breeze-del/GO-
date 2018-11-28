package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hey)
	http.ListenAndServe(":8080", nil)
}

const tpl = `
<html>
	<head>
		<title>Hey</title>
	</head>
	<body>
		<form method="post" action="/">
			Username: <input type="text" name="usname">
			Password: <input type="password" name="pwd">
			<button >确定</button>
		</form>
	</body>
</html>
`

func Hey(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回页面
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(w, nil)
	} else {
		// 解析表单返回的值
		fmt.Println(r.FormValue("usname"))
	}
}
