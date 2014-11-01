package mempool

import (
	"math/rand"
	"testing"
)

func TestMemPool(t *testing.T) {
	pool := New(6)

	if pool.chunkSize != 6 {
		t.Fatal()
	}

	buf := pool.Alloc(5)
	if len(buf) != 5 || len(pool.buf) != 1 {
		t.Fatal()
	}

	buf = pool.Alloc(10)
	if len(buf) != 10 || len(pool.buf) != 2 {
		t.Fatal()
	}

	buf = pool.Alloc(2)
	if len(buf) != 2 || len(pool.buf) != 0 {
		t.Fatal()
	}
}

const N = 1000000
const M = 512

func BenchmarkMemPool(b *testing.B) {
	pool := New(4 * 1024)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < N; i++ {
		_ = pool.Alloc(M)
	}
}

func BenchmarkCommonAlloc(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < N; i++ {
		_ = make([]byte, M)
	}
}

func BenchmarkMemPool_rand(b *testing.B) {
	pool := New(4 * 1024)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < N; i++ {
		_ = pool.Alloc(rand.Intn(500) + 250)
	}
}

func BenchmarkCommonAlloc_rand(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < N; i++ {
		_ = make([]byte, rand.Intn(500)+250)
	}
}

const N2 = 1000000
const M2 = 50

func BenchmarkMemPool_2(b *testing.B) {
	pool := New(4 * 1024)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < N2; i++ {
		_ = pool.Alloc(M2)
	}
}

func BenchmarkCommonAlloc_2(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < N2; i++ {
		_ = make([]byte, M2)
	}
}

func BenchmarkMemPool_rand_2(b *testing.B) {
	pool := New(4 * 1024)

	b.StopTimer()
	b.StartTimer()
	for i := 0; i < N2; i++ {
		_ = pool.Alloc(rand.Intn(50) + 25)
	}
}

func BenchmarkCommonAlloc_rand_2(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < N2; i++ {
		_ = make([]byte, rand.Intn(50)+25)
	}
}
