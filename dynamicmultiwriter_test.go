package dynamicmultiwriter

import (
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var hello = []byte("hello")

func Test(t *testing.T) {
	r1, w1 := io.Pipe()
	r2, w2 := io.Pipe()

	dw := New()
	dw.Add(w1, w2)
	dw.Write(hello)

	p1 := make([]byte, len(hello))
	p2 := make([]byte, len(hello))
	r1.Read(p1)
	r2.Read(p2)

	assert.EqualValues(t, string(p1), hello)
	assert.EqualValues(t, string(p2), hello)

	dw.Remove(w2)
	dw.Write(hello)
	pp1 := make([]byte, len(hello))
	pp2 := make([]byte, len(hello))
	r1.Read(pp1)

	// blocks
	go func() {
		r2.Read(pp2)
	}()
	time.Sleep(time.Second * 1)

	assert.EqualValues(t, string(pp1), hello)
	assert.EqualValues(t, pp2, make([]byte, len(hello)))
}
