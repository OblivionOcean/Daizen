package utils

import (
	"bytes"
	"sync"
)

// BufferPool是一个字节缓冲区的池
type BufferPool struct {
	pool *sync.Pool
	size int
}

// NewBufferPool创建一个新的字节缓冲区池，给定池中缓冲区的个数
func NewBufferPool(size int) *BufferPool {
	bp := &BufferPool{
		size: size,
		pool: &sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0, 1024))
			},
		},
	}

	// 在启动时预先创建指定个数的缓冲区放入池中
	for i := 0; i < size; i++ {
		bp.pool.Put(bp.pool.New())
	}

	return bp
}

// Get从池中获取一个字节缓冲区
//
//go:inline
func (bp *BufferPool) Get() *bytes.Buffer {
	return bp.pool.Get().(*bytes.Buffer)
}

// Put将字节缓冲区释放回池中
//
//go:inline
func (bp *BufferPool) Put(buf *bytes.Buffer) {
	buf.Reset()
	bp.pool.Put(buf)
}
