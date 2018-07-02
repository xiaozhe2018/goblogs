package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

//首页
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.gtpl")
	t.Execute(w, nil)
}

//登录
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if len(username) == 0 || len(password) == 0 {
			io.WriteString(w, "账号或密码不能为空")
		} else {
			io.WriteString(w, "处理逻辑")
		}

	}

}
func main() {
	//设置路由
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	//设置侦听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
