//
// 线程安全的map
//

package safemap

import (
	"sync"
)

type SafeMap struct {
	lock *sync.RWMutex
	mp   map[interface{}]interface{}
}

func New() *SafeMap {
	return &SafeMap{
		lock: new(sync.RWMutex),
		mp:   make(map[interface{}]interface{}),
	}
}

// 返回 key 对应的 value
func (this *SafeMap) Get(key interface{}) interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if value, ok := this.mp[key]; ok {
		return value
	}
	return nil
}

// 设置 (key, value)
func (this *SafeMap) Set(key, value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.mp[key] = value
}

// 判断 key 是否存在
func (this *SafeMap) IsExist(key interface{}) bool {
	this.lock.RLock()
	defer this.lock.RUnlock()
	_, ok := this.mp[key]
	return ok
}

// 删除 key
func (this *SafeMap) Delete(key interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.mp, key)
}

// 返回 safemap 的大小
func (this *SafeMap) Size() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return len(this.mp)
}

// 返回 safemap 中所有元素
func (this *SafeMap) Items() map[interface{}]interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	m := make(map[interface{}]interface{})
	for k, v := range this.mp {
		m[k] = v
	}
	return m
}
