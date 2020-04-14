package serializer

import "sync"

const methodBufferSize = 16

// Buffers provides the scratch space used by Write_ and Read_ functions.
var Buffers = &bufPool{
	pool: sync.Pool{New: func() interface{} { return make([]byte, methodBufferSize) }},
}

type bufPool struct {
	pool sync.Pool
}

func (p *bufPool) Get() []byte {
	return p.pool.Get().([]byte)
}

func (p *bufPool) Put(b []byte) {
	p.pool.Put(b)
}
