// SIMPLE IS THE BEST
package mygo

import (
	"net/http"
	"strings"
)

// router类型包装了一个原生的http.ServerMux类型的指针
// 并提供自有的HandleFunc方法
type router struct {
	mux *http.ServeMux
}

// HandleFunc方法依次调用传入的函数集, 并注入一个自定的上下文类型指针*Ctx，
// 当执行到其中一个函数返回值为false时，后续函数集将不再执行。
func (rtr *router) HandleFunc(pattern string, funcs ...func(*Context, http.ResponseWriter, *http.Request) bool) {
	// 为了更加明确URL路径与URL查询参数的动态与静态分工，提供简洁单一的路由绑定方式
	// router.HandleFunc方法屏蔽了原生http.ServeMux对以"/"结尾的路由模式的支持，
	// 只支持根路由"/"及精确匹配的路由(不以"/"结束的路由)
	if pattern != "/" && strings.HasSuffix(pattern, "/") {
		panic("router: unsupported pattern")
	}
	rtr.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := new(Context)
		for _, f := range funcs {
			if !f(c, w, r) {
				return
			}
		}
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
