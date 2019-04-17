// SIMPLE IS THE BEST
package mygo

import (
	"sync"
)

// Context类型提供并发安全的Context结构
type Context struct {
	sync.Mutex
	data map[string]interface{}
}

// Set方法设置Context
func (c *Context) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()
	// c.data仅在使用时初始化
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	c.data[key] = value
}

// Get方法获取Context
func (c *Context) Get(key string) interface{} {
	c.Lock()
	defer c.Unlock()
	if c.data == nil {
		return nil
	}
	return c.data[key]
}
