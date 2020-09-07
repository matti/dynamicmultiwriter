package dynamicmultiwriter

import (
	"io"
)

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

func (d *DynamicMultiWriter) Write(p []byte) (n int, err error) {
	for w := range d.writers {
		go func(w io.Writer) {
			w.Write(p)
		}(w)
	}
	return len(p), err
}

// Add ...
func (d *DynamicMultiWriter) Add(ws ...io.Writer) {
	for _, w := range ws {
		d.writers[w] = w
	}
}

// Remove ...
func (d *DynamicMultiWriter) Remove(ws ...io.Writer) {
	for _, w := range ws {
		delete(d.writers, w)
	}
}
