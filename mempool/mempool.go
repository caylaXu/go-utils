//
// simple memory pool, not thread safe.
// not recommend to use the memmory pool
//

package mempool

type Pool struct {
	buf       []byte
	chunkSize int
}

func New(chunkSize int) *Pool {
	return &Pool{
		buf:       make([]byte, chunkSize),
		chunkSize: chunkSize,
	}
}

func (p *Pool) Alloc(size int) []byte {
	if len(p.buf) < size {
		n := p.chunkSize
		for n < size {
			n += n
		}
		p.buf = make([]byte, n)
	}

	buf := p.buf[:size]
	p.buf = p.buf[size:]
	return buf
}
