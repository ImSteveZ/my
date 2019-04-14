package mygo

import (
	"sync"
)

// Ctx提供了一个可以安全并发的Context
type Ctx struct {
	sync.Mutex
	data map[string]interface{}
}

// Set方法设置一组键值对到Ctx
func (ctx *Ctx) Set(key string, value interface{}) {
	ctx.Lock()
	defer ctx.Unlock()
	// 为避免不必要的内存开销，ctx.data不预先初始化
	if ctx.data == nil {
		ctx.data = make(map[string]interface{})
	}
	ctx.data[key] = value
}

// Get方法获取键为key的Ctx值
func (ctx *Ctx) Get(key string) (interface{}, bool) {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.data == nil {
		return nil, false
	}
	value, ok := ctx.data[key]
	return value, ok
}
