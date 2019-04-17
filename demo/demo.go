package main

import (
	"fmt"
	"log"
	"mygo"
	"net/http"
)

func main() {
	// 创建mygo.router指针
	r := mygo.NewRouter()

	// 绑定文件服务
	r.HandleFile("/", "./public")
	// 绑定单个处理函数
	r.HandleFunc("/news", News)
	// 线定多个有序http处理函数
	r.HandleFunc("/usr", Auth, Usr)

	// 启动服务
	log.Fatal(http.ListenAndServe(":3000", r))
}

// Auth http处理函数
func Auth(c *mygo.Context, w http.ResponseWriter, r *http.Request) (next bool) {
	// 设置Context
	c.Set("username", "Alice")
	// 向下执行
	next = true
	return
}

// News http处理函数
func News(c *mygo.Context, w http.ResponseWriter, r *http.Request) (next bool) {
	fmt.Fprintf(w, "Welcome to news category\n")
	return
}

// Usr http处理函数
func Usr(c *mygo.Context, w http.ResponseWriter, r *http.Request) (next bool) {
	// 获取Context
	user := c.Get("username")
	if user == "" {
		user = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s\n", user.(string))
	return
}
