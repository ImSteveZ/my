package mygo

import (
	"net/http"
	"sync"
)

// Mux类型包装了一个原生的http.ServerMux类型的指针
// 并提供自有的HandleFunc方法
type Mux struct {
	once     sync.Once
	ServeMux *http.ServeMux
}

// HandleFunc方法依次调用传入的函数集, 并注入一个自定的上下文类型指针*Ctx，当其中一个方法返回值为false时，
// 后续方法集将不再执行。未初始化ServeMux成员的Mux在调用HandleFunc方法时会自动进行一次初始化操作，
// sync.Once保证了该初始化操作的并发安全和单次执行, 在使用过程中应避免将ServeMux赋值为nil，因为这会导致空
// 指针操作的panic
func (mux *Mux) HandleFunc(pattern string, funcs ...func(*Ctx, http.ResponseWriter, *http.Request) bool) {
	mux.once.Do(mux.initServeMux) // 检测初始化ServeMux
	mux.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := new(Ctx)
		for _, f := range funcs {
			if !f(ctx, w, r) {
				return
			}
		}
		return
	})
}

// ServeFile方法开启一个文件服务
func (mux *Mux) ServeFile(pattern, dir string) {
	fileHandle := http.FileServer(http.Dir(dir))
	mux.ServeMux.Handle(pattern, http.StripPrefix(pattern, fileHandle))
}

// Mux类型实现http.Handler接口
func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux.ServeMux.ServeHTTP(w, r)
}

// initServeMux初始化Mux.ServeMux成员
func (mux *Mux) initServeMux() {
	if mux.ServeMux == nil {
		mux.ServeMux = http.NewServeMux()
	}
}

// NewMyMux方法返回一个初始化了ServeMux的Mux类型指针
func NewMyMux() *Mux { return &Mux{ServeMux: http.NewServeMux()} }
