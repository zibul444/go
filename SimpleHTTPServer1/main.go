package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// localhost:8800/
//localhost:8800/?param1=val1&param2%20=val2
//echo -n "test out the server" | nc localhost 8800
func main() {
	http.HandleFunc("/", HomeRouterHandler)  // установим роутер
	err := http.ListenAndServe(":8800", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //анализ аргументов,
	fmt.Println(r.Form) // ввод информации о форме на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	printUrlValues(r.Form)

	fmt.Fprintf(w, "Hello World! 1") // отправляем данные на клиентскую сторону
}

func printUrlValues(f url.Values) {
	for k, v := range f {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}
