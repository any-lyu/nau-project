package pool

import (
	"runtime/debug"
	"testing"
)

func TestBytesBufferPool(t *testing.T) {
	debug.SetGCPercent(-1)
	defer debug.SetGCPercent(100)

	// new
	func(t *testing.T) {
		buffer := BytesBufferPool16KB.GetBytesBufferFromPool()
		defer BytesBufferPool16KB.PutBytesBufferToPool(buffer)

		have := buffer.String()
		want := ""
		if have != want {
			t.Errorf("have:%s, want:%s", have, want)
			return
		}
	}(t)

	// clean up
	{
		for i := 0; i < 10; i++ {
			BytesBufferPool16KB.GetBytesBufferFromPool()
		}
	}

	// new twice
	func(t *testing.T) {
		buffer := BytesBufferPool16KB.GetBytesBufferFromPool()
		defer BytesBufferPool16KB.PutBytesBufferToPool(buffer)

		have := buffer.String()
		want := ""
		if have != want {
			t.Errorf("have:%s, want:%s", have, want)
			return
		}
		buffer.WriteString("buffer")

		buffer2 := BytesBufferPool16KB.GetBytesBufferFromPool()
		defer BytesBufferPool16KB.PutBytesBufferToPool(buffer2)

		have = buffer2.String()
		want = ""
		if have != want {
			t.Errorf("have:%s, want:%s", have, want)
			return
		}
	}(t)

	// clean up
	{
		for i := 0; i < 10; i++ {
			BytesBufferPool16KB.GetBytesBufferFromPool()
		}
	}

	// reuse
	func(t *testing.T) {
		func(t *testing.T) {
			buffer := BytesBufferPool16KB.GetBytesBufferFromPool()
			defer BytesBufferPool16KB.PutBytesBufferToPool(buffer)

			have := buffer.String()
			want := ""
			if have != want {
				t.Errorf("have:%s, want:%s", have, want)
				return
			}
			buffer.WriteString("buffer")
		}(t)

		func(t *testing.T) {
			buffer := BytesBufferPool16KB.GetBytesBufferFromPool()
			defer BytesBufferPool16KB.PutBytesBufferToPool(buffer)

			have := buffer.String()
			want := "buffer"
			if have != want {
				t.Errorf("have:%s, want:%s", have, want)
				return
			}
		}(t)
	}(t)

	// clean up
	{
		for i := 0; i < 10; i++ {
			BytesBufferPool16KB.GetBytesBufferFromPool()
		}
	}

	// reuse and put nil
	func(t *testing.T) {
		func(t *testing.T) {
			buffer := BytesBufferPool16KB.GetBytesBufferFromPool()
			defer BytesBufferPool16KB.PutBytesBufferToPool(buffer)

			have := buffer.String()
			want := ""
			if have != want {
				t.Errorf("have:%s, want:%s", have, want)
				return
			}
			buffer.WriteString("buffer")
		}(t)

		for i := 0; i < 10; i++ {
			BytesBufferPool16KB.PutBytesBufferToPool(nil)
		}

		func(t *testing.T) {
			buffer := BytesBufferPool16KB.GetBytesBufferFromPool()
			defer BytesBufferPool16KB.PutBytesBufferToPool(buffer)

			have := buffer.String()
			want := "buffer"
			if have != want {
				t.Errorf("have:%s, want:%s", have, want)
				return
			}
		}(t)
	}(t)
}
