# MYGO: SIMPLE IS THE BEST
+ mygo在原生http.ServeMux的基础上，提供一个简洁，支持上下文传递与链式http处理的go web路由绑定工具
### 使用示例
```go
package main

import (
	"fmt"
	"log"
	"github.com/iamstevez/mygo"
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
```
