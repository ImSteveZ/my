package main

import (
	"context"
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
	r.HandleFunc("/news", []func(context.Context, func()){}, News)
	// 线定多个有序http处理函数
	r.HandleFunc("/usr", []func(context.Context, func()){Auth}, Usr)

	// 启动服务
	log.Fatal(http.ListenAndServe(":3000", r))
}

// Auth http处理函数
func Auth(ctx context.Context, next func()) {
	ctx.WithValue("username", "Bob")
	next()
}

// News http处理函数
func News(ctx context.Context) {
	meta := ctx.Value("meta").(*mygo.Meta)
	fmt.Fprintf(meta.W, "Welcome to news category\n")
}

// Usr http处理函数
func Usr(ctx context.Context) {
	// 获取Context
	user := ctx.Value("username").(string)
	meta := ctx.Value("meta").(*mygo.Meta)
	fmt.Fprintf(meta.W, "Hello, %s\n", user)
}
