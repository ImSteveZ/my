// SIMPLE IS THE BEST
package mygo

import (
	"context"
	"net/http"
)

// router类型包装了一个原生的http.ServerMux类型的指针
// 并提供自有的HandleFunc方法
type router struct {
	mux *http.ServeMux
}

// Meta : 注入原始http请求引用
type Meta struct {
	Resp http.ResponseWriter
	Req  *http.Request
}

// HandleFunc : 绑定路由中件间与路由处理函数
func (rtr *router) HandleFunc(pattern string, middlewares []func(context.Context, func()), controller func(*context.Context)) {
	rtr.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), "meta", &Meta{Resp: w, Req: r})
		var (
			next func()
			i    int
		)
		next = func() {
			if i > len(middlewares)-1 {
				controller(ctx)
			} else {
				i++
				middlewares[i-1](ctx, next)
			}
		}
		next()
		return
	})
}

// HandleFile方法线定文件服务
func (rtr *router) HandleFile(pattern, dir string) {
	fileServer := http.FileServer(http.Dir(dir))
	rtr.mux.Handle(pattern, http.StripPrefix(pattern, fileServer))
}

// ServeHTTP方法实现http.Handler接口
func (rtr *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rtr.mux.ServeHTTP(w, r)
}

// NewRouter方法返回一个已初始化http.ServeMux的router类型指针
func NewRouter() *router {
	return &router{mux: http.NewServeMux()}
}
