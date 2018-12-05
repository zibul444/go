package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", hello)
	http.ListenAndServe(":8800", nil)
}

func hello(res http.ResponseWriter, req *http.Request){
	res.Header().Set(
		"Content-Typ",
		"text/html",
		)
	fmt.Println(req.FormValue("param1"))
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
	)
}