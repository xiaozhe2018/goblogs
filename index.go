package main

import (
	"log"
	//"fmt"
	"html/template"
	"net/http"
	//"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.gtpl")
	t.Execute(w, nil)
}

func main() {
	//设置路由
	http.HandleFunc("/", index)
	//设置侦听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
