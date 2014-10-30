//
// 自旋锁
// 调用spinslock的线程如果无法获取锁，则一直等待，并不会睡眠  -- 用户态
// 调用mutex lock的线程如果无法获取锁，则会进入睡眠状态，让出cpu时间片  -- 内核态
// 根据测试：
// 在开启runtime.GOMAXPROCS(runtime.NumCPU())的情况下：
// 4核：spinlock的性能大概是mutex 的2倍
//
// 没开启runtime.GOMAXPROCS(runtime.NumCPU())，也就是单核的情况下，mutex的性能略高，
// 并且在该测试程序中 两者的性能 都优于 上面的多核(4核)
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
		// 放弃本时间片(http://golang.org/pkg/runtime/#Gosched)
		// Gosched yields the processor, allowing other goroutines to run.
		// It does not suspend the current goroutine, so execution resumes automatically.
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.flag, 0)
}

func (s *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapInt32(&s.flag, 0, 1)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
