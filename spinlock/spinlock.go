//
// 旋转锁
// 不需要进行上下文切换，而mutex要进行上下文切换
// 根据测试，大部分情况下，spinlock的效率要略高于mutex
//

package spinlock

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	flag int32
}

func (s *SpinLock) Lock() {
	for !s.TryLock() {
		runtime.Gosched() // 放弃本时间片(http://golang.org/pkg/runtime/#Gosched)
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.flag, 0)
}

func (s *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapInt32(&s.flag, 0, 1)
}
