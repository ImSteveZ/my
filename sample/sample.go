package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nicego"
	"time"
)

type CtxKey string

func main() {
	// 获取route句柄
	r := nicego.NewRoute(context.Background())

	r.From("/").Use(Logger).Static("./public")     // 静态文件服务
	r.From("/news").Use(Logger).Do(NewsCtrl)       // 单中间件路由注册
	r.From("/user").Use(Logger, Auth).Do(UserCtrl) // 多中间件路由注册

	// 开启http服务
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}

// Logger 中间件记录访问日志
func Logger(ctx context.Context, next func(context.Context)) {
	t := time.Now()
	_, r := nicego.GetMeta(ctx)
	log.SetFlags(log.Ltime)

	comFmt := "%-4s | %-20s |\n"
	log.Printf(comFmt, r.Method, r.URL.Path)

	next(ctx)

	outFmt := "%-4s | %-20s | %-10s\n"
	log.Printf(outFmt, r.Method, r.URL.Path, time.Now().Sub(t))
}

// Auth 中间件处理用户权限
func Auth(ctx context.Context, next func(context.Context)) {
	authrizedUsers := map[string]struct{}{
		"Alice": struct{}{},
		"Bob":   struct{}{},
	}
	w, r := nicego.GetMeta(ctx)
	params := r.URL.Query()
	if username := params.Get("name"); username != "" {
		if _, ok := authrizedUsers[username]; ok {
			next(context.WithValue(ctx, CtxKey("username"), username))
		} else {
			fmt.Fprintf(w, "Unauthrized user: %s.\n", username)
		}
	} else {
		fmt.Fprintf(w, "Please provide your name.\n")
	}
}

// NewsCtrl 中间件
func NewsCtrl(ctx context.Context) {
	w, _ := nicego.GetMeta(ctx)
	fmt.Fprintf(w, "Welcome to news category.\n")
}

// UserCtrl 中间件
func UserCtrl(ctx context.Context) {
	w, _ := nicego.GetMeta(ctx)
	user := ctx.Value(CtxKey("username")).(string)
	fmt.Fprintf(w, "Hello, %s.\n", user)
}
