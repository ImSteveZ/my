# my
my在原net/http包的基础上实现了一些简洁的golang web工具，支持并发安全的上下文(Contex)传递及链式的http请求与响应处理(中间件)
使用demo:
```go
package main

import (
	"log"
	"github.com/imstevez/mygo"
	"net/http"
)

func main() {
	myMux := mygo.NewMyMux()
	myMux.ServeFile("/", "./public")
	myMux.HandleFunc("/index/", auth, index)
	myMux.HandleFunc("/usr/", auth, usr)
	log.Fatal(http.ListenAndServe(":3000", myMux))
}

func auth(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	ctx.Set("auth", "login")
	return true
}

func index(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	authval, _ := ctx.Get("auth")
	w.Write([]byte("index_" + authval.(string)))
	return true
}

func usr(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	authval, _ := ctx.Get("auth")
	w.Write([]byte("usr_" + authval.(string)))
	return true
}
```
