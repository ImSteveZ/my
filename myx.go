package mygo

import (
	"net/http"
	"sync"
)

// Myx类型包装了一个原生的http.ServerMux类型的指针
// 并提供自有的HandleFunc方法
type Myx struct {
	once     sync.Once
	ServeMux *http.ServeMux
}

// HandleFunc方法依次调用传入的函数集, 并注入一个自定的上下文类型指针*Ctx，当其中一个方法返回值为false时，
// 后续方法集将不再执行。未初始化ServeMux成员的Mux在调用HandleFunc方法时会自动进行一次初始化操作，
// sync.Once保证了该初始化操作的并发安全和单次执行, 在使用过程中应避免将ServeMux赋值为nil，因为这会导致空
// 指针操作的panic
func (myx *Myx) HandleFunc(pattern string, funcs ...func(*Ctx, http.ResponseWriter, *http.Request) bool) {
	myx.once.Do(myx.initServeMux) // 检测初始化ServeMux
	myx.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
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
func (myx *Myx) ServeFile(pattern, dir string) {
	fileHandle := http.FileServer(http.Dir(dir))
	myx.ServeMux.Handle(pattern, http.StripPrefix(pattern, fileHandle))
}

// ServeHTTP方法实现http.Handler接口
func (myx *Myx) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	myx.ServeMux.ServeHTTP(w, r)
}

// initServeMux初始化Mux.ServeMux成员
func (myx *Myx) initServeMux() {
	if myx.ServeMux == nil {
		myx.ServeMux = http.NewServeMux()
	}
}

// NewMyx方法返回一个初始化了ServeMux的Myx类型指针
func NewMyx() *Myx { return &Myx{ServeMux: http.NewServeMux()} }
