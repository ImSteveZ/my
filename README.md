# NICEGO: SIMPLE IS THE BEST
---
### Desc
+ nicego基于标准包的http.ServeMux，在原生路由功能的基础上提供对中间件装饰和上下文传递支持
+ nicego提供方式单一且优雅的外部调用接口：
```go
r := nicego.NewRouter(context.Background())
r.From(路由规则).Use(中间件列表).Do(控制器)
```
### Usage
1. __NewRouter(context.Context) *route__ 函数以传入的context为父节点创建一个实现了http.Handler接口的nicego.route句柄
```go
rt := nicego.NewRouter(context.Background())
```
2. __route.From(pattern string) *router__ 方法包装调用者route指针并以传入的路径pattern创建一个待注入中间件与控制器的router指针(From方法是创建该指针的唯一方式)
```go
rtr := rt.From("/")
```
3. __router.Use(...func(context.Context, func(context.Context))) *router__ 方法将传入的中间件处理函数绑定到调用者router上，并返回自身方便执行链式调用
```go
rtr.Use(Logger, Auth)
```
4. __router.Do(func(context.Context))__ 方法执行实际的路绑定，间接调用封装在router.route类型结构中的http.ServeMux指针，实现路由注册到基于挂载的中间件和控制器构建的handler
```go
rtr.Do(IndexCtrl)
```
5. __router.Static(string)__ 以传入文件夹路径为根进行静态文件服务
```go
rtr.Static("./public")
```
6. __推荐使用链式方式__
```go
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
```
### Todo
+ 基于trie模型重构路由规则
+ RESTFUL风格注册接口支持
### Complete Sample
```go
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
```
