package pool

import (
	"bytes"
	"sync"
)

// BytesBufferPool 是 *bytes.Buffer 的 pool 接口.
type BytesBufferPool interface {
	// GetBytesBufferFromPool 从 pool 里获取一个 *bytes.Buffer.
	GetBytesBufferFromPool() *bytes.Buffer
	// PutBytesBufferToPool 将 *bytes.Buffer 放回 pool, 忽略 nil 的 *bytes.Buffer.
	PutBytesBufferToPool(buf *bytes.Buffer)
}

// BytesBufferPool16KB 是 16KB 大小的 BytesBufferPool
var BytesBufferPool16KB BytesBufferPool = NewBytesBufferPool(16 << 10)

// NewBytesBufferPool 创建一个 bytes.Buffer 大小为 size 的 NewBytesBufferPool
func NewBytesBufferPool(size int) BytesBufferPool {
	if size <= 0 {
		size = 16 << 10
	}
	return &bytesBufferPoolImpl{
		pool: sync.Pool{
			New: func(size int) func() interface{} {
				return func() interface{} {
					return bytes.NewBuffer(make([]byte, 0, size))
				}
			}(size),
		},
	}
}

var _ BytesBufferPool = (*bytesBufferPoolImpl)(nil)

// bytesBufferPoolImpl 是 BytesBufferPool 的一个实现.
type bytesBufferPoolImpl struct {
	pool sync.Pool
}

func (impl *bytesBufferPoolImpl) GetBytesBufferFromPool() *bytes.Buffer {
	return impl.pool.Get().(*bytes.Buffer)
}

func (impl *bytesBufferPoolImpl) PutBytesBufferToPool(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	impl.pool.Put(buf)
}
