package dynamicmultiwriter

import (
	"io"
	"sync"
)

var mux sync.Mutex

// DynamicMultiWriter ...
type DynamicMultiWriter struct {
	writers map[io.Writer]io.Writer
}

// New ...
func New() *DynamicMultiWriter {
	return &DynamicMultiWriter{
		writers: make(map[io.Writer]io.Writer),
	}
}

// Write ...
func (d *DynamicMultiWriter) Write(p []byte) (n int, err error) {
	for w := range d.writers {
		_, err := w.Write(p)
		if err == io.ErrClosedPipe {
			// meanwhile it's closed, I hope somebody .Removes the destination someday
		} else if err != nil {
			panic(err)
		}
	}
	return len(p), err
}

// Add ...
func (d *DynamicMultiWriter) Add(ws ...io.Writer) {
	mux.Lock()
	for _, w := range ws {
		d.writers[w] = w
	}
	mux.Unlock()
}

// Remove ...
func (d *DynamicMultiWriter) Remove(ws ...io.Writer) {
	mux.Lock()
	for _, w := range ws {
		delete(d.writers, w)
	}
	mux.Unlock()
}
