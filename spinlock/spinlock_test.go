package spinlock

import (
	"fmt"
	"github.com/gansidui/go-utils/safemap"
	"sync"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	var s SpinLock
	if !s.TryLock() {
		t.Fatal()
	}
}

func TestLock(t *testing.T) {
	s := &SpinLock{}
	s.Lock()

	if s.TryLock() {
		t.Fatal()
	}

	s.Unlock()

	if !s.TryLock() {
		t.Fatal()
	}
}

type SpinLockMap struct {
	SpinLock
	mp map[int]bool
}

func (s *SpinLockMap) Set(i int) {
	s.Lock()
	defer s.Unlock()
	s.mp[i] = true
}

func (s *SpinLockMap) Get(i int) bool {
	s.Lock()
	defer s.Unlock()
	return s.mp[i]
}

type MutextMap struct {
	sync.Mutex
	mp map[int]bool
}

func (m *MutextMap) Set(i int) {
	m.Lock()
	defer m.Unlock()
	m.mp[i] = true
}

func (m *MutextMap) Get(i int) bool {
	m.Lock()
	defer m.Unlock()
	return m.mp[i]
}

type RWMutexMap struct {
	sync.RWMutex
	mp map[int]bool
}

func (rwm *RWMutexMap) Set(i int) {
	rwm.Lock()
	defer rwm.Unlock()
	rwm.mp[i] = true
}

func (rwm *RWMutexMap) Get(i int) bool {
	rwm.RLock()
	defer rwm.RUnlock()
	return rwm.mp[i]
}

const N = 5000
const M = 1000

func TestSpinLock(t *testing.T) {
	waitGroup := &sync.WaitGroup{}
	s := &SpinLockMap{mp: make(map[int]bool)}

	start := time.Now()

	waitGroup.Add(M * 2)

	for j := 0; j < M; j++ {
		go func() {
			for i := 0; i < N; i++ {
				s.Set(i)
			}
			waitGroup.Done()
		}()

		go func() {
			for i := 0; i < N; i++ {
				s.Get(i)
			}
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println("TestSpinLock:", time.Since(start))
}

func TestMutex(t *testing.T) {
	waitGroup := &sync.WaitGroup{}
	s := &MutextMap{mp: make(map[int]bool)}

	start := time.Now()

	waitGroup.Add(M * 2)

	for j := 0; j < M; j++ {
		go func() {
			for i := 0; i < N; i++ {
				s.Set(i)
			}
			waitGroup.Done()
		}()

		go func() {
			for i := 0; i < N; i++ {
				s.Get(i)
			}
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println("TestMutex:", time.Since(start))
}

func TestRWMutext(t *testing.T) {
	waitGroup := &sync.WaitGroup{}
	s := &RWMutexMap{mp: make(map[int]bool)}

	start := time.Now()

	waitGroup.Add(M * 2)

	for j := 0; j < M; j++ {
		go func() {
			for i := 0; i < N; i++ {
				s.Set(i)
			}
			waitGroup.Done()
		}()

		go func() {
			for i := 0; i < N; i++ {
				s.Get(i)
			}
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println("TestRWMutex:", time.Since(start))
}

func TestSafeMap(t *testing.T) {
	waitGroup := &sync.WaitGroup{}
	s := safemap.New()

	start := time.Now()

	waitGroup.Add(M * 2)

	for j := 0; j < M; j++ {
		go func() {
			for i := 0; i < N; i++ {
				s.Set(i, i)
			}
			waitGroup.Done()
		}()

		go func() {
			for i := 0; i < N; i++ {
				s.Get(i)
			}
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println("TestSafeMap:", time.Since(start))
}
