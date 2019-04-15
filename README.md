# mygo
---
+ mygo对原有http.ServeMux进行封装，提供支持上下文传递，单路径参数，链式中间件注入功能的go web路由绑定工具.
### 使用demo
```go
package main

import (
	"fmt"
	"log"
	"github.com/imstevez/mygo"
	"net/http"
)

func main() {
	myx := mygo.NewMyx()
	myx.ServeFile("/", "./public")
	myx.HandleFunc("/news", News)
	myx.HandleFunc("/usr/:username", Auth, Usr)
	log.Fatal(http.ListenAndServe(":3000", myx))
}

var users = map[string]struct{}{
	"alice": struct{}{},
	"bob":   struct{}{},
}

func Auth(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	if user, ok := ctx.Get("username"); ok {
		if _, ok := users[user.(string)]; ok {
			ctx.Set("login", true)
			return true
		}
	}
	http.NotFound(w, r)
	return false
}

func News(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	w.Write([]byte("welcome to news category\n"))
	return true
}

func Usr(ctx *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	user, _ := ctx.Get("username")
	fmt.Fprintf(w, "Hello, %s\n", user.(string))
	return true
}
```
### 测试
```sh
$ curl localhost:3000
<!doctype html>
<html>Hello, world</html>

$ curl localhost:3000/news
welcome to news category

$ curl localhost:3000/usr/alice
Hello, alice

$ curl localhost:3000/usr/steve
404 page not found
```

