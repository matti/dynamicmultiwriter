package dynamicmultiwriter

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hello = []byte("hello")

func Test(t *testing.T) {
	r1, w1 := io.Pipe()
	r2, w2 := io.Pipe()

	didread1 := make(chan bool)
	didread2 := make(chan bool)

	p1 := make([]byte, len(hello))
	p2 := make([]byte, len(hello))

	go func() {
		for {
			r1.Read(p1)
			didread1 <- true
		}
	}()
	go func() {
		for {
			r2.Read(p2)
			didread2 <- true
		}
	}()

	dw := New()
	dw.Add(w1, w2)
	dw.Write(hello)

	<-didread1
	assert.Equal(t, string(hello), string(p1))
	<-didread2
	assert.Equal(t, string(hello), string(p2))

	dw.Remove(w1)
	p2 = []byte("elloh")
	dw.Write(hello)
	<-didread2
	go func() {
		<-didread1
		panic("r1 got another message")
	}()

	assert.Equal(t, string(hello), string(p2))
}
